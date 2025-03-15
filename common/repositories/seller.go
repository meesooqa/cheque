package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type SellerRepository struct {
	common_db.BaseRepository[models.Seller]
}

func NewSellerRepository() *SellerRepository {
	return &SellerRepository{common_db.BaseRepository[models.Seller]{
		Preloads: []string{
			"SellerPlaces",
		},
	}}
}
