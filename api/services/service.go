package services

import "github.com/meesooqa/cheque/db/db_types"

type Converter[DbModel, PbModel any] interface {
	DataDbToPb(dbItem *DbModel) *PbModel
	DataPbToDb(pbItem *PbModel) *DbModel
}

type FilterProvider[PbGetListRequest any] interface {
	GetFilters(r *PbGetListRequest) []db_types.FilterFunc
}

type BaseService[DbModel, PbModel, PbGetListRequest any] struct {
	Repo           db_types.Repository[DbModel]
	Converter      Converter[DbModel, PbModel]
	FilterProvider FilterProvider[PbGetListRequest]
}

func NewBaseService[DbModel, PbModel, PbGetListRequest any](
	repo db_types.Repository[DbModel],
	converter Converter[DbModel, PbModel],
	filterProvider FilterProvider[PbGetListRequest],
) *BaseService[DbModel, PbModel, PbGetListRequest] {
	return &BaseService[DbModel, PbModel, PbGetListRequest]{
		Repo:           repo,
		Converter:      converter,
		FilterProvider: filterProvider,
	}
}
