package productss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/productpb"
	"github.com/meesooqa/cheque/db/repositories"
)

func GetFilters(req *pb.GetListRequest) []repositories.FilterFunc {
	return []repositories.FilterFunc{
		BrandIDFilter(req.BrandId),
		NameFilter(req.Name),
	}
}

func BrandIDFilter(value uint64) repositories.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("brand_id = ?", value)
		}
		return db
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
