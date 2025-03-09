package productss

import (
	"gorm.io/gorm"

	"cheque-04/api/services"
)

func ExampleFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
