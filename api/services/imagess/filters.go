package imagess

import (
	"gorm.io/gorm"

	pb "receipt-002/api/gen/pb/imagepb/v1"
	"receipt-002/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	return []db_types.FilterFunc{
		ProductIDFilter(req.ProductId),
		NameFilter(req.Name),
		UrlFilter(req.Url),
		IsMainFilter(req.IsMain),
	}
}

func ProductIDFilter(value uint64) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("product_id = ?", value)
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

func UrlFilter(value string) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("url ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func IsMainFilter(value bool) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		// TODO value == false
		if value != false {
			return db.Where("is_main = ?", value)
		}
		return db
	}
}

// TODO order
