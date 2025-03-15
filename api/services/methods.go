package services

import (
	"context"

	"github.com/meesooqa/cheque/common/common_db"
)

func (o *BaseService[DbModel, PbModel]) GetItem(ctx context.Context, id uint64) (*PbModel, error) {
	item, err := o.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return o.converter.DataDbToPb(item), nil
}

func (o *BaseService[DbModel, PbModel]) CreateItem(ctx context.Context, item *PbModel) (*PbModel, error) {
	newDbItem := o.converter.DataPbToDb(item)
	newItem, err := o.Repo.Create(ctx, newDbItem)
	if err != nil {
		return nil, err
	}
	return o.converter.DataDbToPb(newItem), nil
}

func (o *BaseService[DbModel, PbModel]) UpdateItem(ctx context.Context, id uint64, item *PbModel) (*PbModel, error) {
	updatedDbItem := o.converter.DataPbToDb(item)
	updatedItem, err := o.Repo.Update(ctx, id, updatedDbItem)
	if err != nil {
		return nil, err
	}
	return o.converter.DataDbToPb(updatedItem), nil
}

func (o *BaseService[DbModel, PbModel]) DeleteItem(ctx context.Context, id uint64) error {
	return o.Repo.Delete(ctx, id)
}

func (o *BaseService[DbModel, PbModel]) GetList(ctx context.Context, filters []common_db.FilterFunc, sort common_db.SortData, pagination common_db.PaginationData) ([]*PbModel, int64, error) {
	dbItems, total, err := o.Repo.GetList(ctx, filters, sort, pagination)
	if err != nil {
		return nil, 0, err
	}
	var items []*PbModel
	for _, dbItem := range dbItems {
		items = append(items, o.converter.DataDbToPb(dbItem))
	}
	return items, total, nil
}
