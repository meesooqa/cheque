package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
)

type CategoryRepository struct {
	BaseRepository[models.Category]
}

func NewCategoryRepository(dbProvider db_types.DBProvider) *CategoryRepository {
	return &CategoryRepository{BaseRepository[models.Category]{
		DBProvider: dbProvider,
		Preloads: []string{
			"Parent",
		},
	}}
}
