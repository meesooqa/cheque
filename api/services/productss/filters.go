package productss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/productpb"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	return []db_types.FilterFunc{
		BrandIDFilter(req.BrandId),
		NameFilter(req.Name),
	}
}

func BrandIDFilter(value uint64) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("brand_id = ?", value)
		}
		return db
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
