package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/jessevdk/go-flags"

	_ "cheque-04/common"
	"cheque-04/common/common_log"
	"cheque-04/common/config"
	"cheque-04/gorm-gen-proto/gen"
	"cheque-04/gorm-gen-proto/reg"
)

var allowedGenerators = []string{"proto", "services"}

type Generator string

func (o *Generator) UnmarshalFlag(value string) error {
	for _, v := range allowedGenerators {
		if value == v {
			*o = Generator(value)
			return nil
		}
	}
	return fmt.Errorf("unallowed value '%s', allowed: %s", value, strings.Join(allowedGenerators, ", "))
}

type options struct {
	Generator Generator `short:"g" long:"gen" description:"Generator (proto, services)" required:"true" default:"proto"`
}

var conf *config.Conf
var templates *template.Template

func init() {
	c, err := config.Load("var/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	conf = c

	funcMap := template.FuncMap{
		"notEmpty": func(arr interface{}) bool {
			v := reflect.ValueOf(arr)
			return v.Kind() == reflect.Slice && v.Len() > 0
		},
	}
	templates = template.Must(
		template.New("").
			Funcs(funcMap).
			ParseGlob(fmt.Sprintf("%s/*.tmpl", conf.GormGenProto.PathTmpl)),
	)
}

// `go run ./main.go` OR `go run ./main.go --gen=proto` OR `go run ./main.go -g proto`
// `go run ./main.go --gen=services` OR `go run ./main.go -g services`
func main() {
	logger := common_log.InitConsoleLogger(slog.LevelDebug)
	logger.Info("begin")

	var err error
	var opts options
	if _, err = flags.Parse(&opts); err != nil {
		fmt.Println("options parsing", err)
		os.Exit(1)
	}
	logger.Debug("options", slog.String("generator", string(opts.Generator)))

	switch opts.Generator {
	case "proto":
		err = genProto(logger)
	case "services":
		err = genServices()
	default:
	}
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("end")
}

// genProto generates proto files and runs `protoc`
func genProto(logger *slog.Logger) error {
	gg := reg.GetGormDataRegistry()
	if len(gg) == 0 {
		return fmt.Errorf("no gormData available")
	}
	// generate proto files
	pg := gen.NewProto3Generator(conf.GormGenProto, templates)
	for _, g := range gg {
		err := pg.Run(logger, g)
		if err != nil {
			return err
		}
	}
	// generate go files by `protoc`
	pe := gen.NewProtocExecutor()
	for _, g := range gg {
		protoFilePath := g.GetProtoFilePath(conf.GormGenProto)
		err := pe.Run(conf.GormGenProto.ProtoRoot, protoFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}

// genServices generates service servers files
func genServices() error {
	ss := reg.GetSsDataRegistry()
	ssg := gen.NewServiceServerGenerator(conf.GormGenProto, templates)
	if len(ss) == 0 {
		return fmt.Errorf("no ssData available")
	}
	for _, data := range ss {
		err := ssg.Run(data)
		if err != nil {
			return err
		}
	}
	return nil
}
