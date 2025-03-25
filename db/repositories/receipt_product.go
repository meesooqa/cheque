package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
)

type ReceiptProductRepository struct {
	BaseRepository[models.ReceiptProduct]
}

func NewReceiptProductRepository(dbProvider db_types.DBProvider) *ReceiptProductRepository {
	return &ReceiptProductRepository{BaseRepository[models.ReceiptProduct]{
		DBProvider: dbProvider,
		Preloads: []string{
			"Product",
		},
	}}
}
