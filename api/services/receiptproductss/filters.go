package receiptproductss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/services"
)

func ProductIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("product_id = ?", value)
		}
		return db
	}
}

func ReceiptIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("receipt_id = ?", value)
		}
		return db
	}
}

func PriceFilter(valueGt, valueLt int32) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if valueGt > 0 && valueLt > 0 {
			db = db.Where("price BETWEEN ? AND ?", valueGt, valueLt)
		} else if valueGt > 0 {
			db = db.Where("price >= ?", valueGt)
		} else if valueLt > 0 {
			db = db.Where("price <= ?", valueLt)
		}
		return db
	}
}

func SumFilter(valueGt, valueLt int32) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if valueGt > 0 && valueLt > 0 {
			db = db.Where("sum BETWEEN ? AND ?", valueGt, valueLt)
		} else if valueGt > 0 {
			db = db.Where("sum >= ?", valueGt)
		} else if valueLt > 0 {
			db = db.Where("sum <= ?", valueLt)
		}
		return db
	}
}

func QuantityFilter(valueGt, valueLt float64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if valueGt > 0 && valueLt > 0 {
			db = db.Where("quantity BETWEEN ? AND ?", valueGt, valueLt)
		} else if valueGt > 0 {
			db = db.Where("quantity >= ?", valueGt)
		} else if valueLt > 0 {
			db = db.Where("quantity <= ?", valueLt)
		}
		return db
	}
}
