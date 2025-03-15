package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type ImageRepository struct {
	common_db.BaseRepository[models.Image]
}

func NewImageRepository() *ImageRepository {
	return &ImageRepository{common_db.BaseRepository[models.Image]{}}
}
