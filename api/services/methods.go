package services

import (
	"context"

	"github.com/meesooqa/cheque/db/db_types"
)

func (o *BaseService[DbModel, PbModel]) GetList(ctx context.Context, filters []db_types.FilterFunc, sortBy, sortOrder string, pageSize, page int) ([]*PbModel, int64, error) {
	dbItems, total, err := o.Repo.GetList(ctx, filters,
		db_types.SortData{SortField: sortBy, SortOrder: sortOrder},
		db_types.PaginationData{Page: page, PageSize: pageSize})
	if err != nil {
		return nil, 0, err
	}
	var items []*PbModel
	for _, dbItem := range dbItems {
		items = append(items, o.Converter.DataDbToPb(dbItem))
	}
	return items, total, nil
}

func (o *BaseService[DbModel, PbModel]) GetItem(ctx context.Context, id uint64) (*PbModel, error) {
	item, err := o.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return o.Converter.DataDbToPb(item), nil
}

func (o *BaseService[DbModel, PbModel]) CreateItem(ctx context.Context, item *PbModel) (*PbModel, error) {
	newDbItem := o.Converter.DataPbToDb(item)
	newItem, err := o.Repo.Create(ctx, newDbItem)
	if err != nil {
		return nil, err
	}
	return o.Converter.DataDbToPb(newItem), nil
}

func (o *BaseService[DbModel, PbModel]) UpdateItem(ctx context.Context, id uint64, item *PbModel) (*PbModel, error) {
	updatedDbItem := o.Converter.DataPbToDb(item)
	updatedItem, err := o.Repo.Update(ctx, id, updatedDbItem)
	if err != nil {
		return nil, err
	}
	return o.Converter.DataDbToPb(updatedItem), nil
}

func (o *BaseService[DbModel, PbModel]) DeleteItem(ctx context.Context, id uint64) error {
	return o.Repo.Delete(ctx, id)
}
