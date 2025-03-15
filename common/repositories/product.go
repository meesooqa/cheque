package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type ProductRepository struct {
	common_db.BaseRepository[models.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{common_db.BaseRepository[models.Product]{
		Preloads: []string{
			"Brand",
			"Categories",
			"Images",
		},
	}}
}
