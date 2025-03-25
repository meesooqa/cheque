package brandss

import (
	"gorm.io/gorm"

	pb "receipt-002/api/gen/pb/brandpb/v1"
	"receipt-002/db/db_types"
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
