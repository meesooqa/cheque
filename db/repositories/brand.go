package repositories

import (
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

type BrandRepository struct {
	BaseRepository[models.Brand]
}

func NewBrandRepository(dbProvider db_types.DBProvider) *BrandRepository {
	return &BrandRepository{BaseRepository[models.Brand]{
		DBProvider: dbProvider,
	}}
}
