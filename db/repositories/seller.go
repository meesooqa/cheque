package repositories

import (
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

type SellerRepository struct {
	BaseRepository[models.Seller]
}

func NewSellerRepository(dbProvider db_types.DBProvider) *SellerRepository {
	return &SellerRepository{BaseRepository[models.Seller]{
		DBProvider: dbProvider,
	}}
}
