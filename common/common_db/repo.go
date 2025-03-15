package common_db

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type FilterFunc func(db *gorm.DB) *gorm.DB

type SortData struct {
	SortField string
	SortOrder string
}

type PaginationData struct {
	Page     int
	PageSize int
}

type Repository[DbModel any] interface {
	GetList(ctx context.Context, filters []FilterFunc, sort SortData, pagination PaginationData) ([]*DbModel, int64, error)
	Get(ctx context.Context, id uint64) (*DbModel, error)
	Create(ctx context.Context, newItem *DbModel) (*DbModel, error)
	Update(ctx context.Context, id uint64, updatedItem *DbModel) (*DbModel, error)
	Delete(ctx context.Context, id uint64) error
}

type BaseRepository[DbModel any] struct {
	logger *slog.Logger
}

func (o *BaseRepository[DbModel]) GetList(ctx context.Context, filters []FilterFunc, sort SortData, pagination PaginationData) ([]*DbModel, int64, error) {
	// TODO Preload
	db := GetDB(ctx)
	query := db.Model(new(DbModel))
	if len(filters) > 0 {
		for _, filter := range filters {
			query = filter(query)
		}
	}
	var total int64
	// before setting the limit
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	o.addSort(query, sort)
	o.addPagination(query, pagination)

	var dbItems []*DbModel
	if err := query.Find(&dbItems).Error; err != nil {
		return nil, 0, err
	}
	return dbItems, total, nil
}

func (o *BaseRepository[DbModel]) Get(ctx context.Context, id uint64) (*DbModel, error) {
	// TODO Preload
	db := GetDB(ctx)
	var item DbModel
	if err := db.First(&item, id).Error; err != nil {
		return nil, fmt.Errorf("item with ID %d not found: %w", id, err)
	}
	return &item, nil
}

func (o *BaseRepository[DbModel]) Create(ctx context.Context, newItem *DbModel) (*DbModel, error) {
	db := GetDB(ctx)
	if err := db.Create(&newItem).Error; err != nil {
		return nil, err
	}
	return newItem, nil
}

func (o *BaseRepository[DbModel]) Update(ctx context.Context, id uint64, updatedItem *DbModel) (*DbModel, error) {
	db := GetDB(ctx)
	var dbItem DbModel
	if err := db.First(&dbItem, id).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&dbItem).Updates(updatedItem).Error; err != nil {
		return nil, err
	}
	return o.Get(ctx, id)
}

func (o *BaseRepository[DbModel]) Delete(ctx context.Context, id uint64) error {
	db := GetDB(ctx)
	var dbItem DbModel
	result := db.Delete(&dbItem, id)
	if result.Error != nil {
		return fmt.Errorf("item deleting: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("item with ID %d not found", id)
	}
	return nil
}

func (o *BaseRepository[DbModel]) addSort(query *gorm.DB, sort SortData) {
	order := "asc"
	if sort.SortOrder == "desc" {
		order = "desc"
	}
	if sort.SortField != "" {
		query = query.Order(sort.SortField + " " + order)
	}
}

func (o *BaseRepository[DbModel]) addPagination(query *gorm.DB, pagination PaginationData) {
	if pagination.PageSize > 0 {
		query = query.Limit(pagination.PageSize)
	}
	if pagination.Page > 0 {
		offset := pagination.PageSize * (pagination.Page - 1)
		query = query.Offset(offset)
	}
}
