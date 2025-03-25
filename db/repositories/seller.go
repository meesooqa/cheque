package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
)

type SellerRepository struct {
	BaseRepository[models.Seller]
}

func NewSellerRepository(dbProvider db_types.DBProvider) *SellerRepository {
	return &SellerRepository{BaseRepository[models.Seller]{
		DBProvider: dbProvider,
		Preloads: []string{
			"SellerPlaces",
		},
	}}
}
