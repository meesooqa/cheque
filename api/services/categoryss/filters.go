package categoryss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/api/services"
)

func GetFilters(req *pb.GetListRequest) []services.FilterFunc {
	return []services.FilterFunc{
		NameFilter(req.Name),
		ParentIDFilter(req.ParentId),
	}
}

func NameFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("name ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func ParentIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("parent_id = ?", value)
		}
		return db
	}
}
