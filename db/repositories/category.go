package repositories

import (
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
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
