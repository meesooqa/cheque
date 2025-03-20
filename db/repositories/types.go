package repositories

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/db_provider"
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

type HasAssociations[DbModel any] interface {
	UpdateAssociations(db *gorm.DB, item *DbModel, updatedData *DbModel) error
}

type BaseRepository[DbModel any] struct {
	DBProvider db_provider.DBProvider
	Self       HasAssociations[DbModel]
	Preloads   []string
}

func (o *BaseRepository[DbModel]) GetList(ctx context.Context, filters []FilterFunc, sort SortData, pagination PaginationData) ([]*DbModel, int64, error) {
	db := o.DBProvider.GetDB(ctx)
	db = o.preload(db)
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
	db := o.DBProvider.GetDB(ctx)
	db = o.preload(db)
	var item DbModel
	if err := db.First(&item, id).Error; err != nil {
		return nil, fmt.Errorf("item with ID %d not found: %w", id, err)
	}
	return &item, nil
}

func (o *BaseRepository[DbModel]) Create(ctx context.Context, newItem *DbModel) (*DbModel, error) {
	db := o.DBProvider.GetDB(ctx)
	if err := db.Create(&newItem).Error; err != nil {
		return nil, err
	}
	return newItem, nil
}

func (o *BaseRepository[DbModel]) Update(ctx context.Context, id uint64, updatedItem *DbModel) (*DbModel, error) {
	dbItem, err := o.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	db := o.DBProvider.GetDB(ctx)
	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	if err = tx.Model(dbItem).Updates(updatedItem).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update item: %w", err)
	}
	if hasAssoc, ok := any(o.Self).(HasAssociations[DbModel]); ok {
		err = hasAssoc.UpdateAssociations(tx, dbItem, updatedItem)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update associations: %w", err)
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return o.Get(ctx, id)
}

func (o *BaseRepository[DbModel]) Delete(ctx context.Context, id uint64) error {
	db := o.DBProvider.GetDB(ctx)
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

func (o *BaseRepository[DbModel]) preload(db *gorm.DB) *gorm.DB {
	if len(o.Preloads) > 0 {
		for _, preload := range o.Preloads {
			db = db.Preload(preload)
		}
	}
	return db
}
