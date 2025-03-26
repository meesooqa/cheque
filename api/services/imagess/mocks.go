// Code generated by template. DO NOT EDIT.
package imagess

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
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

func (m *MockRepository) GetList(ctx context.Context, filters []db_types.FilterFunc, sort db_types.SortData, pagination db_types.PaginationData) ([]*DbModel, int64, error) {
	args := m.Called(ctx, filters, sort, pagination)
	var result []*DbModel
	if args.Get(0) != nil {
		result = args.Get(0).([]*DbModel)
	}
	return result, args.Get(1).(int64), args.Error(2)
}

func (m *MockRepository) Get(ctx context.Context, id uint64) (*DbModel, error) {
	args := m.Called(ctx, id)
	var result *DbModel
	if args.Get(0) != nil {
		result = args.Get(0).(*DbModel)
	}
	return result, args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, newItem *DbModel) (*DbModel, error) {
	args := m.Called(ctx, newItem)
	var result *DbModel
	if args.Get(0) != nil {
		result = args.Get(0).(*DbModel)
	}
	return result, args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, id uint64, updatedItem *DbModel) (*DbModel, error) {
	args := m.Called(ctx, id, updatedItem)
	var result *DbModel
	if args.Get(0) != nil {
		result = args.Get(0).(*DbModel)
	}
	return result, args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockConverter is a mock for Converter[DbModel, pb.Model]
type MockConverter struct {
	mock.Mock
}

func (m *MockConverter) DataDbToPb(dbItem *DbModel) *pb.Model {
	args := m.Called(dbItem)
	var result *pb.Model
	if args.Get(0) != nil {
		result = args.Get(0).(*pb.Model)
	}
	return result
}

func (m *MockConverter) DataPbToDb(pbItem *pb.Model) *DbModel {
	args := m.Called(pbItem)
	var result *DbModel
	if args.Get(0) != nil {
		result = args.Get(0).(*DbModel)
	}
	return result
}

// MockFilterProvider is a mock for FilterProvider[pb.GetListRequest]
type MockFilterProvider struct {
	mock.Mock
}

func (m *MockFilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	args := m.Called(r)
	var result []db_types.FilterFunc
	if args.Get(0) != nil {
		result = args.Get(0).([]db_types.FilterFunc)
	}
	return result
}
