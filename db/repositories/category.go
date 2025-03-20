package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

type CategoryRepository struct {
	BaseRepository[models.Category]
}

func NewCategoryRepository(dbProvider db_provider.DBProvider) *CategoryRepository {
	return &CategoryRepository{BaseRepository[models.Category]{
		DBProvider: dbProvider,
		Preloads: []string{
			"Parent",
		},
	}}
}
