package imagess

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	ProductIDFilter := db_types.ModelExactFieldFilter[DbModel]("product_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	UrlFilter := db_types.ModelFieldFilter[DbModel]("url")
	IsMainFilter := db_types.ModelExactFieldFilter[DbModel]("is_main")
	// TODO order
	return []db_types.FilterFunc{
		ProductIDFilter(r.ProductId),
		NameFilter(r.Name),
		UrlFilter(r.Url),
		IsMainFilter(r.IsMain),
	}
}
