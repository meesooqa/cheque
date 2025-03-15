package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type BrandRepository struct {
	common_db.BaseRepository[models.Brand]
}

func NewBrandRepository() *BrandRepository {
	return &BrandRepository{common_db.BaseRepository[models.Brand]{}}
}
