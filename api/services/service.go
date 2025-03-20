package services

import "github.com/meesooqa/cheque/db/repositories"

type Converter[DbModel any, PbModel any] interface {
	DataDbToPb(dbItem *DbModel) *PbModel
	DataPbToDb(pbItem *PbModel) *DbModel
}

type BaseService[DbModel any, PbModel any] struct {
	Repo      repositories.Repository[DbModel]
	converter Converter[DbModel, PbModel]
}

func NewBaseService[DbModel any, PbModel any](repo repositories.Repository[DbModel], converter Converter[DbModel, PbModel]) *BaseService[DbModel, PbModel] {
	return &BaseService[DbModel, PbModel]{
		Repo:      repo,
		converter: converter,
	}
}
