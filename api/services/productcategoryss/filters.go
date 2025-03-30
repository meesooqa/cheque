package productcategoryss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/productcategorypb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	ProductIDFilter := db_types.ModelExactFieldFilter[DbModel]("product_id")
	CategoryIDFilter := db_types.ModelExactFieldFilter[DbModel]("category_id")
	return []db_types.FilterFunc{
		ProductIDFilter(r.ProductId),
		CategoryIDFilter(r.CategoryId),
	}
}
