package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_types"
)

type ImageRepository struct {
	BaseRepository[models.Image]
}

func NewImageRepository(dbProvider db_types.DBProvider) *ImageRepository {
	return &ImageRepository{BaseRepository[models.Image]{
		DBProvider: dbProvider,
	}}
}
