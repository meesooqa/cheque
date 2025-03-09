package reg

import "cheque-04/gorm-gen-proto/gen"

var gormDataRegistry []*gen.GormForTmpl
var ssDataRegistry []*gen.SsTmplData

func RegisterGormData(data []*gen.GormForTmpl) {
	gormDataRegistry = append(gormDataRegistry, data...)
}

func GetGormDataRegistry() []*gen.GormForTmpl {
	return gormDataRegistry
}

func RegisterSsData(data []*gen.SsTmplData) {
	ssDataRegistry = append(ssDataRegistry, data...)
}

func GetSsDataRegistry() []*gen.SsTmplData {
	return ssDataRegistry
}
