package imagess

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	ProductIDFilter := db_types.ModelExactFieldFilter[DbModel]("product_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	UrlFilter := db_types.ModelFieldFilter[DbModel]("url")
	IsMainFilter := db_types.ModelExactFieldFilter[DbModel]("is_main")
	// TODO order
	return []db_types.FilterFunc{
		ProductIDFilter(req.ProductId),
		NameFilter(req.Name),
		UrlFilter(req.Url),
		IsMainFilter(req.IsMain),
	}
}
