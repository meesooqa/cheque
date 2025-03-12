package services

import (
	"log/slog"

	"gorm.io/gorm"
)

type Converter[DbModel any, PbModel any] interface {
	DataDbToPb(dbItem *DbModel) *PbModel
	DataPbToDb(pbItem *PbModel) *DbModel
}

type FilterFunc func(db *gorm.DB) *gorm.DB

type BaseService[DbModel any, PbModel any] struct {
	Logger    *slog.Logger
	db        *gorm.DB
	converter Converter[DbModel, PbModel]
}

func NewBaseService[T any, U any](log *slog.Logger, db *gorm.DB, converter Converter[T, U]) *BaseService[T, U] {
	return &BaseService[T, U]{
		Logger:    log,
		db:        db,
		converter: converter,
	}
}
