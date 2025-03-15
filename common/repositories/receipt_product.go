package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type ReceiptProductRepository struct {
	common_db.BaseRepository[models.ReceiptProduct]
}

func NewReceiptProductRepository() *ReceiptProductRepository {
	return &ReceiptProductRepository{common_db.BaseRepository[models.ReceiptProduct]{
		Preloads: []string{
			"Receipt",
			"Product",
		},
	}}
}
