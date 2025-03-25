package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
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
