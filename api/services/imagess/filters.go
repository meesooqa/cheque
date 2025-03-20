package imagess

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/imagepb"
	"github.com/meesooqa/cheque/db/repositories"
)

func GetFilters(req *pb.GetListRequest) []repositories.FilterFunc {
	return []repositories.FilterFunc{
		ProductIDFilter(req.ProductId),
		NameFilter(req.Name),
		UrlFilter(req.Url),
		IsMainFilter(req.IsMain),
	}
}

func ProductIDFilter(value uint64) repositories.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("product_id = ?", value)
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

func UrlFilter(value string) repositories.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("url ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func IsMainFilter(value bool) repositories.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		// TODO value == false
		if value != false {
			return db.Where("is_main = ?", value)
		}
		return db
	}
}
