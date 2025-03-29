package categoryss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/categorypb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	ParentIDFilter := db_types.ModelExactFieldFilter[DbModel]("parent_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	NameEnFilter := db_types.ModelFieldFilter[DbModel]("name_en")
	return []db_types.FilterFunc{
		NameFilter(r.Name),
		NameEnFilter(r.NameEn),
		ParentIDFilter(r.ParentId),
	}
}

func (o *FilterProvider) GetChildrenFilters(r *pb.GetChildrenRequest) []db_types.FilterFunc {
	ParentIDFilter := db_types.ModelExactFieldFilter[DbModel]("parent_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	NameEnFilter := db_types.ModelFieldFilter[DbModel]("name_en")
	return []db_types.FilterFunc{
		NameFilter(r.Name),
		NameEnFilter(r.NameEn),
		ParentIDFilter(r.Id),
	}
}
