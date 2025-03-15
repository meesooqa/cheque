package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type CategoryRepository struct {
	common_db.BaseRepository[models.Category]
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{common_db.BaseRepository[models.Category]{
		Preloads: []string{
			"Parent",
			"Children",
		},
	}}
}
