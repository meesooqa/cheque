package repositories

import (
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

type ProductCategoryRepository struct {
	BaseRepository[models.ProductCategory]
}

func NewProductCategoryRepository(dbProvider db_types.DBProvider) *ProductCategoryRepository {
	return &ProductCategoryRepository{BaseRepository[models.ProductCategory]{
		DBProvider: dbProvider,
	}}
}
