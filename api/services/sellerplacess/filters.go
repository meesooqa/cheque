package sellerplacess

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/sellerplacepb"
	"github.com/meesooqa/cheque/common/common_db"
)

func GetFilters(req *pb.GetListRequest) []common_db.FilterFunc {
	return []common_db.FilterFunc{
		SellerIDFilter(req.SellerId),
		NameFilter(req.Name),
		AddressFilter(req.Address),
		EmailFilter(req.Email),
	}
}

func SellerIDFilter(value uint64) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("seller_id = ?", value)
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

func AddressFilter(value string) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func EmailFilter(value string) common_db.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
