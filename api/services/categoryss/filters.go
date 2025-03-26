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
	return []db_types.FilterFunc{
		NameFilter(r.Name),
		ParentIDFilter(r.ParentId),
	}
}
