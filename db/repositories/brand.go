package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

type BrandRepository struct {
	BaseRepository[models.Brand]
}

func NewBrandRepository(dbProvider db_provider.DBProvider) *BrandRepository {
	return &BrandRepository{BaseRepository[models.Brand]{
		DBProvider: dbProvider,
	}}
}
