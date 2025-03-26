package categoryss

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/gen/pb/categorypb/v1"
	"github.com/meesooqa/cheque/api/services"
	"github.com/meesooqa/cheque/db/db_types"
)

func TestNewService(t *testing.T) {
	// Arrange
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)

	// Act
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, repo, service.Repo)
	assert.Equal(t, converter, service.Converter)
	assert.Equal(t, filterProvider, service.FilterProvider)
}

func TestService_GetList(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	filters := []db_types.FilterFunc{}
	sortBy := "name"
	sortOrder := "asc"
	pageSize := 10
	page := 1
	expectedSort := db_types.SortData{SortField: sortBy, SortOrder: sortOrder}
	expectedPagination := db_types.PaginationData{Page: page, PageSize: pageSize}

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var parentID1 uint = 0
		dbItems := []*DbModel{
			{
				Model:    gorm.Model{ID: 1},
				ParentID: nil,
				Name:     "Item 1",
			},
			{
				Model:    gorm.Model{ID: 2},
				ParentID: &parentID1,
				Name:     "Item 2",
			},
		}
		pbItems := []*pb.Model{
			{
				Id:       1,
				ParentId: 0,
				Name:     "Item 1",
			},
			{
				Id:       2,
				ParentId: 1,
				Name:     "Item 2",
			},
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
		var dbCategories []*DbModel
		var total int64 = 0

		repo.On("GetList", ctx, filters, expectedSort, expectedPagination).Return(dbCategories, total, nil).Once()

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

func TestService_GetItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		dbItem := &DbModel{
			Model:    gorm.Model{ID: 1},
			ParentID: nil,
			Name:     "Item 1",
		}
		pbItem := &pb.Model{
			Id:       1,
			ParentId: 0,
			Name:     "Item 1",
		}

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

func TestService_CreateItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var parentID uint = 1
		pbItemToCreate := &pb.Model{
			ParentId: uint64(parentID),
			Name:     "New Item",
		}
		dbItemToCreate := &DbModel{
			ParentID: &parentID,
			Name:     "New Item",
		}
		createdDbItem := &DbModel{
			Model:    gorm.Model{ID: 2},
			ParentID: &parentID,
			Name:     "New Item",
		}
		createdPbItem := &pb.Model{
			Id:       2,
			ParentId: uint64(parentID),
			Name:     "New Item",
		}

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
		var parentID uint = 1
		pbItemToCreate := &pb.Model{
			ParentId: uint64(parentID),
			Name:     "New Item",
		}
		dbItemToCreate := &DbModel{
			ParentID: &parentID,
			Name:     "New Item",
		}
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

func TestService_UpdateItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		var parentID uint = 2
		pbItemToUpdate := &pb.Model{
			Id:       id,
			ParentId: uint64(parentID),
			Name:     "Updated Item",
		}
		dbItemToUpdate := &DbModel{
			Model:    gorm.Model{ID: 1},
			ParentID: &parentID,
			Name:     "Updated Item",
		}
		updateddbItem := &DbModel{
			Model:    gorm.Model{ID: 1},
			ParentID: &parentID,
			Name:     "Updated Item",
		}
		updatedpbItem := &pb.Model{
			Id:       id,
			ParentId: uint64(parentID),
			Name:     "Updated Item",
		}

		converter.On("DataPbToDb", pbItemToUpdate).Return(dbItemToUpdate).Once()
		repo.On("Update", ctx, id, dbItemToUpdate).Return(updateddbItem, nil).Once()
		converter.On("DataDbToPb", updateddbItem).Return(updatedpbItem).Once()

		// Act
		result, err := service.UpdateItem(ctx, id, pbItemToUpdate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, updatedpbItem, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		var parentID uint = 2
		pbItemToUpdate := &pb.Model{
			Id:       id,
			ParentId: uint64(parentID),
			Name:     "Updated Item",
		}
		dbItemToUpdate := &DbModel{
			Model:    gorm.Model{ID: 1},
			ParentID: &parentID,
			Name:     "Updated Item",
		}
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

func TestService_DeleteItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

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
