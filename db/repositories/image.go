package repositories

import (
	"receipt-002/db/db_types"
	"receipt-002/db/models"
)

type ImageRepository struct {
	BaseRepository[models.Image]
}

func NewImageRepository(dbProvider db_types.DBProvider) *ImageRepository {
	return &ImageRepository{BaseRepository[models.Image]{
		DBProvider: dbProvider,
	}}
}
