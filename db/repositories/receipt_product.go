package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

type ReceiptProductRepository struct {
	BaseRepository[models.ReceiptProduct]
}

func NewReceiptProductRepository(dbProvider db_provider.DBProvider) *ReceiptProductRepository {
	return &ReceiptProductRepository{BaseRepository[models.ReceiptProduct]{
		DBProvider: dbProvider,
		Preloads: []string{
			"Product",
		},
	}}
}
