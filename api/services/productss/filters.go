package productss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/productpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	BrandIDFilter := db_types.ModelExactFieldFilter[DbModel]("brand_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	return []db_types.FilterFunc{
		BrandIDFilter(req.BrandId),
		NameFilter(req.Name),
	}
}
