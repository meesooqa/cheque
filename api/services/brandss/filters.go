package brandss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/services"
)

func NameFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("name ILIKE ?", "%"+value+"%")
		}
		return db
	}
}
