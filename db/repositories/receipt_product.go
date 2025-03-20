package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_types"
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
