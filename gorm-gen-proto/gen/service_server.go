package gen

import (
	"os"
	"path/filepath"
	"text/template"

	"cheque-04/common/config"
)

type ServiceServerGenerator struct {
	conf      *config.GormGenProtoConfig
	templates *template.Template
}

type SsTmplData struct {
	Package string
	DbModel string
	//Imports []string
	ImportPb       string
	ImportServices string
	ImportModels   string
}

func NewServiceServerGenerator(conf *config.GormGenProtoConfig, templates *template.Template) *ServiceServerGenerator {
	return &ServiceServerGenerator{
		conf:      conf,
		templates: templates,
	}
}

func (o *ServiceServerGenerator) Run(data *SsTmplData) error {
	//files := []string{"data", "filters", "methods", "service"}
	files := []string{"service"}
	for _, fileCode := range files {
		err := o.createFile(fileCode, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *ServiceServerGenerator) createFile(fileCode string, data *SsTmplData) error {
	// create file
	protoFilePath := o.conf.ServicesRoot + "/" + data.Package + "/" + fileCode + ".go"
	dir := filepath.Dir(protoFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	file, err := os.Create(protoFilePath)
	if err != nil {
		return err
	}
	// write
	err = o.templates.ExecuteTemplate(file, fileCode+".go.tmpl", data)
	if err != nil {
		return err
	}
	return nil
}
