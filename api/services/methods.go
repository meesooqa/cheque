package services

import (
	"fmt"

	"gorm.io/gorm"
)

func (o *BaseService[DbModel, PbModel]) GetItem(id uint64) (*PbModel, error) {
	var item DbModel
	if err := o.db.First(&item, id).Error; err != nil {
		return nil, fmt.Errorf("item with ID %d not found: %w", id, err)
	}
	return o.converter.DataDbToPb(&item), nil
}

func (o *BaseService[DbModel, PbModel]) CreateItem(item *PbModel) (*PbModel, error) {
	newItem := o.converter.DataPbToDb(item)
	if err := o.db.Create(&newItem).Error; err != nil {
		return nil, err
	}
	return o.converter.DataDbToPb(newItem), nil
}

func (o *BaseService[DbModel, PbModel]) UpdateItem(id uint64, item *PbModel) (*PbModel, error) {
	var dbItem DbModel
	if err := o.db.First(&dbItem, id).Error; err != nil {
		return nil, err
	}

	updatedItem := o.converter.DataPbToDb(item)
	if err := o.db.Model(&dbItem).Updates(updatedItem).Error; err != nil {
		return nil, err
	}

	if err := o.db.First(&dbItem, id).Error; err != nil {
		return nil, err
	}
	return o.converter.DataDbToPb(&dbItem), nil
}

func (o *BaseService[DbModel, PbModel]) DeleteItem(id uint64) error {
	var dbItem DbModel
	result := o.db.Delete(&dbItem, id)
	if result.Error != nil {
		return fmt.Errorf("item deleting: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("item with ID %d not found", id)
	}
	return nil
}

func (o *BaseService[DbModel, PbModel]) GetList(filters []FilterFunc, sortBy, sortOrder string, pageSize, page int) ([]*PbModel, int64, error) {
	query := o.db.Model(new(DbModel))

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
	o.addSort(query, sortBy, sortOrder)
	o.addPagination(query, pageSize, page)

	var dbItems []DbModel
	if err := query.Find(&dbItems).Error; err != nil {
		return nil, 0, err
	}
	var items []*PbModel
	for _, dbItem := range dbItems {
		items = append(items, o.converter.DataDbToPb(&dbItem))
	}
	return items, total, nil
}

func (o *BaseService[DbModel, PbModel]) addSort(query *gorm.DB, sortBy, sortOrder string) {
	order := "asc"
	if sortOrder == "desc" {
		order = "desc"
	}
	if sortBy != "" {
		query = query.Order(sortBy + " " + order)
	}
}

func (o *BaseService[DbModel, PbModel]) addPagination(query *gorm.DB, pageSize, page int) {
	if pageSize > 0 {
		query = query.Limit(pageSize)
	}
	if page > 0 {
		offset := pageSize * (page - 1)
		query = query.Offset(offset)
	}
}
