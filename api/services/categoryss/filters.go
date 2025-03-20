package categoryss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/db/repositories"
)

func GetFilters(req *pb.GetListRequest) []repositories.FilterFunc {
	return []repositories.FilterFunc{
		NameFilter(req.Name),
		ParentIDFilter(req.ParentId),
	}
}

func NameFilter(value string) repositories.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("name ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func ParentIDFilter(value uint64) repositories.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("parent_id = ?", value)
		}
		return db
	}
}
