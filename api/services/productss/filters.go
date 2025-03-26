package productss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/productpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	BrandIDFilter := db_types.ModelExactFieldFilter[DbModel]("brand_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	return []db_types.FilterFunc{
		BrandIDFilter(r.BrandId),
		NameFilter(r.Name),
	}
}
