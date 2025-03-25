package repositories

import (
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

type SellerPlaceRepository struct {
	BaseRepository[models.SellerPlace]
}

func NewSellerPlaceRepository(dbProvider db_types.DBProvider) *SellerPlaceRepository {
	return &SellerPlaceRepository{BaseRepository[models.SellerPlace]{
		DBProvider: dbProvider,
	}}
}
