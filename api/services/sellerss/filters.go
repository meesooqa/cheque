package sellerss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/sellerpb"
	"github.com/meesooqa/cheque/common/common_db"
)

func GetFilters(req *pb.GetListRequest) []common_db.FilterFunc {
	return []common_db.FilterFunc{
		NameFilter(req.Name),
		InnFilter(req.Inn),
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

func InnFilter(value string) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
