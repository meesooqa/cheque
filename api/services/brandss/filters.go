package brandss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/brandpb"
	"github.com/meesooqa/cheque/api/services"
)

func GetFilters(req *pb.GetListRequest) []services.FilterFunc {
	return []services.FilterFunc{
		NameFilter(req.Name),
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
