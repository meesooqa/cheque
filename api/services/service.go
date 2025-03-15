package services

import (
	"github.com/meesooqa/cheque/common/common_db"
)

type Converter[DbModel any, PbModel any] interface {
	DataDbToPb(dbItem *DbModel) *PbModel
	DataPbToDb(pbItem *PbModel) *DbModel
}

type BaseService[DbModel any, PbModel any] struct {
	Repo      common_db.Repository[DbModel]
	converter Converter[DbModel, PbModel]
}

func NewBaseService[T any, U any](repo common_db.Repository[T], converter Converter[T, U]) *BaseService[T, U] {
	return &BaseService[T, U]{
		Repo:      repo,
		converter: converter,
	}
}
