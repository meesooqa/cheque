package imagess

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/imagepb"
	"github.com/meesooqa/cheque/common/common_db"
)

func GetFilters(req *pb.GetListRequest) []common_db.FilterFunc {
	return []common_db.FilterFunc{
		ProductIDFilter(req.ProductId),
		NameFilter(req.Name),
		UrlFilter(req.Url),
		IsMainFilter(req.IsMain),
	}
}

func ProductIDFilter(value uint64) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("product_id = ?", value)
		}
		return db
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

func UrlFilter(value string) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("url ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func IsMainFilter(value bool) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		// TODO value == false
		if value != false {
			return db.Where("is_main = ?", value)
		}
		return db
	}
}
