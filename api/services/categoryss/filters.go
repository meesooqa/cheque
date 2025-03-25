package categoryss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/categorypb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	ParentIDFilter := db_types.ModelExactFieldFilter[DbModel]("parent_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	return []db_types.FilterFunc{
		NameFilter(req.Name),
		ParentIDFilter(req.ParentId),
	}
}
