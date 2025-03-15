package sellerss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/sellerpb"
	"github.com/meesooqa/cheque/api/services"
)

func GetFilters(req *pb.GetListRequest) []services.FilterFunc {
	return []services.FilterFunc{
		NameFilter(req.Name),
		InnFilter(req.Inn),
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

func InnFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
