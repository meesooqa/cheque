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

	_ "github.com/meesooqa/cheque/common"
	"github.com/meesooqa/cheque/common/common_log"
	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/gorm-gen-proto/gen"
	"github.com/meesooqa/cheque/gorm-gen-proto/reg"
)

var allowedGenerators = []string{"proto", "protoc", "services"}

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
	Generator Generator `short:"g" long:"gen" default:"proto" required:"true" description:"generator (proto, protoc, services)"`
	Conf      string    `short:"f" long:"conf" env:"CHEQUE_CONF" default:"../etc/config.yml" description:"config file (yml)"`
}

var conf *config.Conf
var templates *template.Template

// `go run ./main.go` OR `go run ./main.go --gen=proto` OR `go run ./main.go -g proto`
// `go run ./main.go --gen=protoc` OR `go run ./main.go -g protoc`
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
	logger.Debug("options", slog.Any("opts", opts))

	c, err := config.Load(opts.Conf)
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

	switch opts.Generator {
	case "proto":
		err = genProto(logger)
	case "protoc":
		err = genProtoc(logger)
	case "services":
		err = genServices()
	default:
	}
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("end")
}

// genProto generates proto files
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
	return nil
}

// genProtoc runs `protoc`
func genProtoc(logger *slog.Logger) error {
	gg := reg.GetGormDataRegistry()
	if len(gg) == 0 {
		return fmt.Errorf("no gormData available")
	}
	// generate go files by `protoc`
	pe := gen.NewProtocExecutor()
	for _, g := range gg {
		err := pe.Run(conf.GormGenProto.ProtoRoot, conf.GormGenProto.ProtocRoot, g.Package, g.GetProtoFileBaseName(conf.GormGenProto))
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
