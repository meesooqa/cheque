package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
)

type SellerPlaceRepository struct {
	BaseRepository[models.SellerPlace]
}

func NewSellerPlaceRepository(dbProvider db_types.DBProvider) *SellerPlaceRepository {
	return &SellerPlaceRepository{BaseRepository[models.SellerPlace]{
		DBProvider: dbProvider,
	}}
}
