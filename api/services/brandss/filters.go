package brandss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/brandpb"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	return []db_types.FilterFunc{
		NameFilter(req.Name),
	}
}

func NameFilter(value string) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("name ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
