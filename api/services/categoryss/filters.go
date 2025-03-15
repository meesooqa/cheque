package categoryss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/common/common_db"
)

func GetFilters(req *pb.GetListRequest) []common_db.FilterFunc {
	return []common_db.FilterFunc{
		NameFilter(req.Name),
		ParentIDFilter(req.ParentId),
	}
}

func NameFilter(value string) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("name ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func ParentIDFilter(value uint64) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("parent_id = ?", value)
		}
		return db
	}
}
