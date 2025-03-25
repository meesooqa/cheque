package repositories

import (
	"context"
	"fmt"
	"reflect"

	"gorm.io/gorm"

	"receipt-002/db/db_types"
)

type BaseRepository[DbModel any] struct {
	DBProvider db_types.DBProvider
	Self       db_types.HasAssociations[DbModel]
	Preloads   []string
}

func (o *BaseRepository[DbModel]) GetList(ctx context.Context, filters []db_types.FilterFunc, sort db_types.SortData, pagination db_types.PaginationData) ([]*DbModel, int64, error) {
	db, err := o.DBProvider.GetDB(ctx)
	if err != nil {
		return nil, 0, err
	}
	db = o.preload(db)
	query := db.Model(new(DbModel))
	if len(filters) > 0 {
		for _, filter := range filters {
			query = filter(query)
		}
	}
	var total int64
	// before setting the limit
	if err = query.Count(&total).Error; err != nil {
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
	db, err := o.DBProvider.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	db = o.preload(db)
	var item DbModel
	if err = db.First(&item, id).Error; err != nil {
		return nil, fmt.Errorf("item with ID %d not found: %w", id, err)
	}
	return &item, nil
}

func (o *BaseRepository[DbModel]) Create(ctx context.Context, newItem *DbModel) (*DbModel, error) {
	db, err := o.DBProvider.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	// Start a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	// Create the item within the transaction
	if err = tx.Create(newItem).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	// Save associations using our common method
	if err = o.saveAssociations(tx, newItem, newItem); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	if field := reflect.ValueOf(newItem).Elem().FieldByName("ID"); field.IsValid() && field.CanUint() {
		// Get fresh item with all associations loaded
		return o.Get(ctx, field.Uint())
	} else {
		return newItem, nil
	}
}

func (o *BaseRepository[DbModel]) Update(ctx context.Context, id uint64, updatedItem *DbModel) (*DbModel, error) {
	dbItem, err := o.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	db, err := o.DBProvider.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	if err = tx.Model(dbItem).Updates(updatedItem).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	if err = o.saveAssociations(tx, dbItem, updatedItem); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return o.Get(ctx, id)
}

func (o *BaseRepository[DbModel]) Delete(ctx context.Context, id uint64) error {
	db, err := o.DBProvider.GetDB(ctx)
	if err != nil {
		return err
	}
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

func (o *BaseRepository[DbModel]) addSort(query *gorm.DB, sort db_types.SortData) {
	order := "asc"
	if sort.SortOrder == "desc" {
		order = "desc"
	}
	if sort.SortField != "" {
		query = query.Order(sort.SortField + " " + order)
	}
}

func (o *BaseRepository[DbModel]) addPagination(query *gorm.DB, pagination db_types.PaginationData) {
	if pagination.PageSize > 0 {
		query = query.Limit(pagination.PageSize)
	}
	if pagination.Page > 0 {
		offset := pagination.PageSize * (pagination.Page - 1)
		query = query.Offset(offset)
	}
}

// saveAssociations handles saving associations within a transaction
func (o *BaseRepository[DbModel]) saveAssociations(tx *gorm.DB, item *DbModel, updatedItem *DbModel) error {
	if hasAssoc, ok := any(o.Self).(db_types.HasAssociations[DbModel]); ok {
		err := hasAssoc.UpdateAssociations(tx, item, updatedItem)
		if err != nil {
			return fmt.Errorf("failed to update associations: %w", err)
		}
	}
	return nil
}

func (o *BaseRepository[DbModel]) preload(db *gorm.DB) *gorm.DB {
	if len(o.Preloads) > 0 {
		for _, preload := range o.Preloads {
			db = db.Preload(preload)
		}
	}
	return db
}
