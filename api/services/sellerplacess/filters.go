package sellerplacess

import (
	"gorm.io/gorm"

	"cheque-04/api/services"
)

func SellerIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("seller_id=?", value)
		}
		return db
	}
}

func NameFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("value ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func AddressFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func EmailFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("inn ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
