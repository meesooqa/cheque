package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type SellerPlaceRepository struct {
	common_db.BaseRepository[models.SellerPlace]
}

func NewSellerPlaceRepository() *SellerPlaceRepository {
	return &SellerPlaceRepository{common_db.BaseRepository[models.SellerPlace]{}}
}
