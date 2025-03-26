package brandss

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/db_types"
)

// MockDBProvider mocks DBProvider
type MockDBProvider struct {
	DB *gorm.DB
}

func (p *MockDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	return p.DB, nil
}

// MockRepository mocks Repository
type MockRepository struct {
	mock.Mock
}

func (r *MockRepository) GetList(ctx context.Context, filters []db_types.FilterFunc, sort db_types.SortData, pagination db_types.PaginationData) ([]*DbModel, int64, error) {
	args := r.Called(ctx, filters, sort, pagination)
	return args.Get(0).([]*DbModel), args.Get(1).(int64), args.Error(2)
}

func (r *MockRepository) Get(ctx context.Context, id uint64) (*DbModel, error) {
	args := r.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*DbModel), args.Error(1)
}

func (r *MockRepository) Create(ctx context.Context, item *DbModel) (*DbModel, error) {
	args := r.Called(ctx, item)
	return args.Get(0).(*DbModel), args.Error(1)
}

func (r *MockRepository) Update(ctx context.Context, id uint64, item *DbModel) (*DbModel, error) {
	args := r.Called(ctx, id, item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*DbModel), args.Error(1)
}

func (r *MockRepository) Delete(ctx context.Context, id uint64) error {
	args := r.Called(ctx, id)
	return args.Error(0)
}
