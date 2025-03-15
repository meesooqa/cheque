package brandss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/brandpb"
	"github.com/meesooqa/cheque/common/common_db"
)

func GetFilters(req *pb.GetListRequest) []common_db.FilterFunc {
	return []common_db.FilterFunc{
		NameFilter(req.Name),
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
