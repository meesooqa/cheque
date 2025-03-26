package sellerss

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/gen/pb/sellerpb/v1"
	"github.com/meesooqa/cheque/api/services"
	"github.com/meesooqa/cheque/db/db_types"
)

func TestNewSellerService(t *testing.T) {
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

func TestSellerService_GetList(t *testing.T) {
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
		dbSellers := []*DbModel{
			{Model: gorm.Model{ID: 1}, Name: "Seller 1", Inn: "123456789012"},
			{Model: gorm.Model{ID: 2}, Name: "Seller 2", Inn: "987654321098"},
		}
		pbSellers := []*pb.Model{
			{Id: 1, Name: "Seller 1", Inn: "123456789012"},
			{Id: 2, Name: "Seller 2", Inn: "987654321098"},
		}
		var total int64 = 2

		repo.On("GetList", ctx, filters, expectedSort, expectedPagination).Return(dbSellers, total, nil).Once()

		// Set up the converter expectations for each item
		for i, dbSeller := range dbSellers {
			converter.On("DataDbToPb", dbSeller).Return(pbSellers[i]).Once()
		}

		// Act
		result, resultTotal, err := service.GetList(ctx, filters, sortBy, sortOrder, pageSize, page)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, pbSellers, result)
		assert.Equal(t, total, resultTotal)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Empty List", func(t *testing.T) {
		// Arrange
		var dbSellers []*DbModel
		var total int64 = 0

		repo.On("GetList", ctx, filters, expectedSort, expectedPagination).Return(dbSellers, total, nil).Once()

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

func TestSellerService_GetItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		dbSeller := &DbModel{Model: gorm.Model{ID: uint(id)}, Name: "Seller 1", Inn: "123456789012"}
		pbSeller := &pb.Model{Id: id, Name: "Seller 1", Inn: "123456789012"}

		repo.On("Get", ctx, id).Return(dbSeller, nil).Once()
		converter.On("DataDbToPb", dbSeller).Return(pbSeller).Once()

		// Act
		result, err := service.GetItem(ctx, id)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, pbSeller, result)
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

func TestSellerService_CreateItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		pbSellerToCreate := &pb.Model{Name: "New Seller", Inn: "123456789012"}
		dbSellerToCreate := &DbModel{Name: "New Seller", Inn: "123456789012"}
		createdDbSeller := &DbModel{Model: gorm.Model{ID: 1}, Name: "New Seller", Inn: "123456789012"}
		createdPbSeller := &pb.Model{Id: 1, Name: "New Seller", Inn: "123456789012"}

		converter.On("DataPbToDb", pbSellerToCreate).Return(dbSellerToCreate).Once()
		repo.On("Create", ctx, dbSellerToCreate).Return(createdDbSeller, nil).Once()
		converter.On("DataDbToPb", createdDbSeller).Return(createdPbSeller).Once()

		// Act
		result, err := service.CreateItem(ctx, pbSellerToCreate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, createdPbSeller, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		pbSellerToCreate := &pb.Model{Name: "New Seller", Inn: "123456789012"}
		dbSellerToCreate := &DbModel{Name: "New Seller", Inn: "123456789012"}
		expectedError := errors.New("repository error")

		converter.On("DataPbToDb", pbSellerToCreate).Return(dbSellerToCreate).Once()
		repo.On("Create", ctx, dbSellerToCreate).Return(nil, expectedError).Once()

		// Act
		result, err := service.CreateItem(ctx, pbSellerToCreate)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})
}

func TestSellerService_UpdateItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repo := new(MockRepository)
	converter := new(MockConverter)
	filterProvider := new(MockFilterProvider)
	service := services.NewBaseService[DbModel, pb.Model, pb.GetListRequest](repo, converter, filterProvider)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		pbSellerToUpdate := &pb.Model{Id: id, Name: "Updated Seller", Inn: "123456789012"}
		dbSellerToUpdate := &DbModel{Model: gorm.Model{ID: uint(id)}, Name: "Updated Seller", Inn: "123456789012"}
		updatedDbSeller := &DbModel{Model: gorm.Model{ID: uint(id)}, Name: "Updated Seller", Inn: "123456789012"}
		updatedPbSeller := &pb.Model{Id: id, Name: "Updated Seller", Inn: "123456789012"}

		converter.On("DataPbToDb", pbSellerToUpdate).Return(dbSellerToUpdate).Once()
		repo.On("Update", ctx, id, dbSellerToUpdate).Return(updatedDbSeller, nil).Once()
		converter.On("DataDbToPb", updatedDbSeller).Return(updatedPbSeller).Once()

		// Act
		result, err := service.UpdateItem(ctx, id, pbSellerToUpdate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, updatedPbSeller, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Arrange
		var id uint64 = 1
		pbSellerToUpdate := &pb.Model{Id: id, Name: "Updated Seller", Inn: "123456789012"}
		dbSellerToUpdate := &DbModel{Model: gorm.Model{ID: uint(id)}, Name: "Updated Seller", Inn: "123456789012"}
		expectedError := errors.New("repository error")

		converter.On("DataPbToDb", pbSellerToUpdate).Return(dbSellerToUpdate).Once()
		repo.On("Update", ctx, id, dbSellerToUpdate).Return(nil, expectedError).Once()

		// Act
		result, err := service.UpdateItem(ctx, id, pbSellerToUpdate)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
		converter.AssertExpectations(t)
	})
}

func TestSellerService_DeleteItem(t *testing.T) {
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
