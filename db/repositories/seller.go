package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

type SellerRepository struct {
	BaseRepository[models.Seller]
}

func NewSellerRepository(dbProvider db_provider.DBProvider) *SellerRepository {
	return &SellerRepository{BaseRepository[models.Seller]{
		DBProvider: dbProvider,
		Preloads: []string{
			"SellerPlaces",
		},
	}}
}
