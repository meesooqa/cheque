package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type ReceiptRepository struct {
	common_db.BaseRepository[models.Receipt]
}

func NewReceiptRepository() *ReceiptRepository {
	return &ReceiptRepository{common_db.BaseRepository[models.Receipt]{
		Preloads: []string{
			"Operator",
			"SellerPlace",
			"ReceiptProducts",
		},
	}}
}
