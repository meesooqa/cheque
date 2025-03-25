package brandss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/brandpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	return []db_types.FilterFunc{
		NameFilter(req.Name),
	}
}
