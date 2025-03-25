package db_types

import "gorm.io/gorm"

// FieldFilter creates a filter that applies a SQL LIKE condition on a specific field
func FieldFilter(fieldName, value string) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where(fieldName+" ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

// ExactFieldFilter creates a filter that applies an exact match condition on a field
func ExactFieldFilter(fieldName string, value interface{}) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != nil && value != "" && value != 0 {
			return db.Where(fieldName+" = ?", value)
		}
		return db
	}
}

// ModelFieldFilter creates a type-safe field filter for a specific model
func ModelFieldFilter[DbModel any](field string) func(value string) FilterFunc {
	return func(value string) FilterFunc {
		return FieldFilter(field, value)
	}
}
