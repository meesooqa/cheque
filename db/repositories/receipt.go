package repositories

import (
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

type ReceiptRepository struct {
	BaseRepository[models.Receipt]
}

func NewReceiptRepository(dbProvider db_types.DBProvider) *ReceiptRepository {
	return &ReceiptRepository{BaseRepository[models.Receipt]{
		DBProvider: dbProvider,
		Preloads: []string{
			"SellerPlace",
			"ReceiptProducts",
		},
	}}
}
