package brandss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/brandpb"
	"github.com/meesooqa/cheque/db/repositories"
)

func GetFilters(req *pb.GetListRequest) []repositories.FilterFunc {
	return []repositories.FilterFunc{
		NameFilter(req.Name),
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
