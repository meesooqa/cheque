package repositories

import (
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

type ImageRepository struct {
	BaseRepository[models.Image]
}

func NewImageRepository(dbProvider db_provider.DBProvider) *ImageRepository {
	return &ImageRepository{BaseRepository[models.Image]{
		DBProvider: dbProvider,
	}}
}
