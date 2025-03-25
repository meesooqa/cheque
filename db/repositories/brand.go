package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
)

type BrandRepository struct {
	BaseRepository[models.Brand]
}

func NewBrandRepository(dbProvider db_types.DBProvider) *BrandRepository {
	return &BrandRepository{BaseRepository[models.Brand]{
		DBProvider: dbProvider,
	}}
}
