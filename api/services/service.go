package services

import "receipt-002/db/db_types"

type Converter[DbModel any, PbModel any] interface {
	DataDbToPb(dbItem *DbModel) *PbModel
	DataPbToDb(pbItem *PbModel) *DbModel
}

type BaseService[DbModel any, PbModel any] struct {
	Repo      db_types.Repository[DbModel]
	converter Converter[DbModel, PbModel]
}

func NewBaseService[DbModel any, PbModel any](repo db_types.Repository[DbModel], converter Converter[DbModel, PbModel]) *BaseService[DbModel, PbModel] {
	return &BaseService[DbModel, PbModel]{
		Repo:      repo,
		converter: converter,
	}
}
