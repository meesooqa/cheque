package categoryss

import (
	"gorm.io/gorm"

	pb "receipt-002/api/gen/pb/categorypb/v1"
	"receipt-002/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	return []db_types.FilterFunc{
		NameFilter(req.Name),
		ParentIDFilter(req.ParentId),
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

func ParentIDFilter(value uint64) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("parent_id = ?", value)
		}
		return db
	}
}
