package sellerplacess

import (
	"gorm.io/gorm"

	pb "receipt-002/api/gen/pb/sellerplacepb/v1"
	"receipt-002/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	return []db_types.FilterFunc{
		SellerIDFilter(req.SellerId),
		NameFilter(req.Name),
		AddressFilter(req.Address),
		EmailFilter(req.Email),
	}
}

func SellerIDFilter(value uint64) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("seller_id = ?", value)
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

func AddressFilter(value string) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func EmailFilter(value string) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
