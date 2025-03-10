package productcategoryss

import (
	"gorm.io/gorm"

	"cheque-04/api/services"
)

func ProductIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("product_id = ?", value)
		}
		return db
	}
}

func CategoryIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("category_id = ?", value)
		}
		return db
	}
}
