package productcategoryss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/services"
)

func ProductIdFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("product_id = ?", value)
		}
		return db
	}
}

func CategoryIdFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("category_id = ?", value)
		}
		return db
	}
}
