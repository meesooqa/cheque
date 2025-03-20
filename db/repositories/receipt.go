package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

type ReceiptRepository struct {
	BaseRepository[models.Receipt]
}

func NewReceiptRepository(dbProvider db_provider.DBProvider) *ReceiptRepository {
	return &ReceiptRepository{BaseRepository[models.Receipt]{
		DBProvider: dbProvider,
		Preloads: []string{
			"SellerPlace",
			"ReceiptProducts",
		},
	}}
}
