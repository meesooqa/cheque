package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/meesooqa/cheque/db/db_types"
)

// MockRepository is a mock for db_types.Repository
type MockRepository[T any] struct {
	mock.Mock
}

func (m *MockRepository[T]) GetList(ctx context.Context, filters []db_types.FilterFunc, sort db_types.SortData, pagination db_types.PaginationData) ([]*T, int64, error) {
	args := m.Called(ctx, filters, sort, pagination)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*T), args.Get(1).(int64), args.Error(2)
}

func (m *MockRepository[T]) Get(ctx context.Context, id uint64) (*T, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) Create(ctx context.Context, newItem *T) (*T, error) {
	args := m.Called(ctx, newItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) Update(ctx context.Context, id uint64, updatedItem *T) (*T, error) {
	args := m.Called(ctx, id, updatedItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockConverter is a mock for Converter
type MockConverter[DbModel any, PbModel any] struct {
	mock.Mock
}

func (m *MockConverter[DbModel, PbModel]) DataDbToPb(dbItem *DbModel) *PbModel {
	args := m.Called(dbItem)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*PbModel)
}

func (m *MockConverter[DbModel, PbModel]) DataPbToDb(pbItem *PbModel) *DbModel {
	args := m.Called(pbItem)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*DbModel)
}

// MockFilterProvider is a mock for FilterProvider
type MockFilterProvider[PbGetListRequest any] struct {
	mock.Mock
}

func (m *MockFilterProvider[PbGetListRequest]) GetFilters(r *PbGetListRequest) []db_types.FilterFunc {
	args := m.Called(r)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]db_types.FilterFunc)
}

// Define test models
type TestDbModel struct {
	ID   uint64
	Name string
}

type TestPbModel struct {
	ID   uint64
	Name string
}

type TestPbGetListRequest struct {
	ID   uint64
	Name string
}

func TestNewBaseService(t *testing.T) {
	// Arrange
	repo := new(MockRepository[TestDbModel])
	converter := new(MockConverter[TestDbModel, TestPbModel])
	filterProvider := new(MockFilterProvider[TestPbGetListRequest])

	// Act
	service := NewBaseService[TestDbModel, TestPbModel, TestPbGetListRequest](repo, converter, filterProvider)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, repo, service.Repo)
	// Note: converter is private, so we can't directly check it
}

func TestBaseService_GetList(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository[TestDbModel])
	converter := new(MockConverter[TestDbModel, TestPbModel])
	filterProvider := new(MockFilterProvider[TestPbGetListRequest])
	service := NewBaseService[TestDbModel, TestPbModel, TestPbGetListRequest](repo, converter, filterProvider)

	filters := []db_types.FilterFunc{}
	sortBy := "name"
	sortOrder := "asc"
	pageSize := 10
	page := 1
	expectedSort := db_types.SortData{SortField: sortBy, SortOrder: sortOrder}
	expectedPagination := db_types.PaginationData{Page: page, PageSize: pageSize}

	t.Run("Success", func(t *testing.T) {
		// Arrange
		dbItems := []*TestDbModel{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
		}
		pbItems := []*TestPbModel{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
		}
		var total int64 = 2

		repo.On("GetList", ctx, filters, expectedSort, expectedPagination).Return(dbItems, total, nil).Once()

		// Set up the converter expectations for each item
		for i, dbItem := range dbItems {
			converter.On("DataDbToPb", dbItem).Return(pbItems[i]).Once()
		}

		// Act
		result, resultTotal, err := service.GetList(ctx, filters, sortBy, sortOrder, pageSize, page)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, pbItems, result)
		assert.Equal(t, total, resultTotal)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Empty List", func(t *testing.T) {
		// Arrange
		var dbItems []*TestDbModel
		var total int64 = 0

		repo.On("GetList", ctx, filters, expectedSort, expectedPagination).Return(dbItems, total, nil).Once()

		// Act
		result, resultTotal, err := service.GetList(ctx, filters, sortBy, sortOrder, pageSize, page)

		// Assert
		require.NoError(t, err)
		assert.Empty(t, result)
		assert.Equal(t, total, resultTotal)
		repo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		expectedError := errors.New("repository error")
		repo.On("GetList", ctx, filters, expectedSort, expectedPagination).Return(nil, int64(0), expectedError).Once()

		// Act
		result, resultTotal, err := service.GetList(ctx, filters, sortBy, sortOrder, pageSize, page)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
		assert.Equal(t, int64(0), resultTotal)
		repo.AssertExpectations(t)
	})
}

func TestBaseService_GetItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository[TestDbModel])
	converter := new(MockConverter[TestDbModel, TestPbModel])
	filterProvider := new(MockFilterProvider[TestPbGetListRequest])
	service := NewBaseService[TestDbModel, TestPbModel, TestPbGetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		dbItem := &TestDbModel{ID: id, Name: "Item 1"}
		pbItem := &TestPbModel{ID: id, Name: "Item 1"}

		repo.On("Get", ctx, id).Return(dbItem, nil).Once()
		converter.On("DataDbToPb", dbItem).Return(pbItem).Once()

		// Act
		result, err := service.GetItem(ctx, id)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, pbItem, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		expectedError := errors.New("repository error")

		repo.On("Get", ctx, id).Return(nil, expectedError).Once()

		// Act
		result, err := service.GetItem(ctx, id)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
	})
}

func TestBaseService_CreateItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository[TestDbModel])
	converter := new(MockConverter[TestDbModel, TestPbModel])
	filterProvider := new(MockFilterProvider[TestPbGetListRequest])
	service := NewBaseService[TestDbModel, TestPbModel, TestPbGetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		pbItemToCreate := &TestPbModel{Name: "New Item"}
		dbItemToCreate := &TestDbModel{Name: "New Item"}
		createdDbItem := &TestDbModel{ID: 1, Name: "New Item"}
		createdPbItem := &TestPbModel{ID: 1, Name: "New Item"}

		converter.On("DataPbToDb", pbItemToCreate).Return(dbItemToCreate).Once()
		repo.On("Create", ctx, dbItemToCreate).Return(createdDbItem, nil).Once()
		converter.On("DataDbToPb", createdDbItem).Return(createdPbItem).Once()

		// Act
		result, err := service.CreateItem(ctx, pbItemToCreate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, createdPbItem, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		pbItemToCreate := &TestPbModel{Name: "New Item"}
		dbItemToCreate := &TestDbModel{Name: "New Item"}
		expectedError := errors.New("repository error")

		converter.On("DataPbToDb", pbItemToCreate).Return(dbItemToCreate).Once()
		repo.On("Create", ctx, dbItemToCreate).Return(nil, expectedError).Once()

		// Act
		result, err := service.CreateItem(ctx, pbItemToCreate)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})
}

func TestBaseService_UpdateItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository[TestDbModel])
	converter := new(MockConverter[TestDbModel, TestPbModel])
	filterProvider := new(MockFilterProvider[TestPbGetListRequest])
	service := NewBaseService[TestDbModel, TestPbModel, TestPbGetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		pbItemToUpdate := &TestPbModel{ID: id, Name: "Updated Item"}
		dbItemToUpdate := &TestDbModel{ID: id, Name: "Updated Item"}
		updatedDbItem := &TestDbModel{ID: id, Name: "Updated Item"}
		updatedPbItem := &TestPbModel{ID: id, Name: "Updated Item"}

		converter.On("DataPbToDb", pbItemToUpdate).Return(dbItemToUpdate).Once()
		repo.On("Update", ctx, id, dbItemToUpdate).Return(updatedDbItem, nil).Once()
		converter.On("DataDbToPb", updatedDbItem).Return(updatedPbItem).Once()

		// Act
		result, err := service.UpdateItem(ctx, id, pbItemToUpdate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, updatedPbItem, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		pbItemToUpdate := &TestPbModel{ID: id, Name: "Updated Item"}
		dbItemToUpdate := &TestDbModel{ID: id, Name: "Updated Item"}
		expectedError := errors.New("repository error")

		converter.On("DataPbToDb", pbItemToUpdate).Return(dbItemToUpdate).Once()
		repo.On("Update", ctx, id, dbItemToUpdate).Return(nil, expectedError).Once()

		// Act
		result, err := service.UpdateItem(ctx, id, pbItemToUpdate)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})
}

func TestBaseService_DeleteItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository[TestDbModel])
	converter := new(MockConverter[TestDbModel, TestPbModel])
	filterProvider := new(MockFilterProvider[TestPbGetListRequest])
	service := NewBaseService[TestDbModel, TestPbModel, TestPbGetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		repo.On("Delete", ctx, id).Return(nil).Once()

		// Act
		err := service.DeleteItem(ctx, id)

		// Assert
		require.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		expectedError := errors.New("repository error")
		repo.On("Delete", ctx, id).Return(expectedError).Once()

		// Act
		err := service.DeleteItem(ctx, id)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		repo.AssertExpectations(t)
	})
}
