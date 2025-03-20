package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_types"
)

type SellerPlaceRepository struct {
	BaseRepository[models.SellerPlace]
}

func NewSellerPlaceRepository(dbProvider db_types.DBProvider) *SellerPlaceRepository {
	return &SellerPlaceRepository{BaseRepository[models.SellerPlace]{
		DBProvider: dbProvider,
	}}
}
