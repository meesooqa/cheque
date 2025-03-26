package receiptproductss

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/gen/pb/receiptproductpb/v1"
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
	sortBy := "price"
	sortOrder := "desc"
	pageSize := 10
	page := 1
	expectedSort := db_types.SortData{SortField: sortBy, SortOrder: sortOrder}
	expectedPagination := db_types.PaginationData{Page: page, PageSize: pageSize}
	t.Run("Success", func(t *testing.T) {
		// Arrange
		productCodeData1 := `{"code": "12345"}`
		productCodeData2 := `{"code": "67890"}`
		dbItems := []*DbModel{{
			Model:           gorm.Model{ID: 1},
			ProductID:       1,
			ReceiptID:       1,
			Price:           1000,
			Quantity:        2.5,
			Sum:             2500,
			ProductCodeData: &productCodeData1,
		}, {
			Model:           gorm.Model{ID: 2},
			ProductID:       2,
			ReceiptID:       1,
			Price:           500,
			Quantity:        1.0,
			Sum:             500,
			ProductCodeData: &productCodeData2,
		}}
		pbItems := []*pb.Model{{
			Id:        1,
			ProductId: 1,
			ReceiptId: 1,
			Price:     1000,
			Quantity:  2.5,
			Sum:       2500,
			//ProductCodeData: productCodeData1,
		}, {
			Id:        2,
			ProductId: 2,
			ReceiptId: 1,
			Price:     500,
			Quantity:  1.0,
			Sum:       500,
			//ProductCodeData: productCodeData2,
		}}
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
		var dbItems []*DbModel
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
		productCodeData := `{"code": "123456"}`
		dbItem := &DbModel{
			Model:           gorm.Model{ID: 1},
			ProductID:       100,
			ReceiptID:       1,
			Price:           1000,
			Quantity:        2.0,
			Sum:             2000,
			ProductCodeData: &productCodeData,
		}
		pbItem := &pb.Model{
			Id:        1,
			ProductId: 100,
			ReceiptId: 1,
			Price:     1000,
			Quantity:  2.0,
			Sum:       2000,
			//ProductCodeData: productCodeData,
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
		productCodeData := `{"code": "654321"}`
		pbItemToCreate := &pb.Model{
			ProductId: 102,
			ReceiptId: 2,
			Price:     2000,
			Quantity:  1.5,
			Sum:       3000,
			//ProductCodeData: productCodeData,
		}
		dbItemToCreate := &DbModel{
			ProductID:       102,
			ReceiptID:       2,
			Price:           2000,
			Quantity:        1.5,
			Sum:             3000,
			ProductCodeData: &productCodeData,
		}
		createdDbItem := &DbModel{
			Model:           gorm.Model{ID: 3},
			ProductID:       102,
			ReceiptID:       2,
			Price:           2000,
			Quantity:        1.5,
			Sum:             3000,
			ProductCodeData: &productCodeData,
		}
		createdPbItem := &pb.Model{
			Id:        3,
			ProductId: 102,
			ReceiptId: 2,
			Price:     2000,
			Quantity:  1.5,
			Sum:       3000,
			//ProductCodeData: productCodeData,
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
		productCodeData := `{"code": "654321"}`
		pbItemToCreate := &pb.Model{
			ProductId: 102,
			ReceiptId: 2,
			Price:     2000,
			Quantity:  1.5,
			Sum:       3000,
			//ProductCodeData: productCodeData,
		}
		dbItemToCreate := &DbModel{
			ProductID:       102,
			ReceiptID:       2,
			Price:           2000,
			Quantity:        1.5,
			Sum:             3000,
			ProductCodeData: &productCodeData,
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
		productCodeData := `{"code": "123456", "updated": true}`
		pbItemToUpdate := &pb.Model{
			Id:        id,
			ProductId: 100,
			ReceiptId: 1,
			Price:     1100,
			Quantity:  2.5,
			Sum:       2750,
			//ProductCodeData: productCodeData,
		}
		dbItemToUpdate := &DbModel{
			Model:           gorm.Model{ID: 1},
			ProductID:       100,
			ReceiptID:       1,
			Price:           1100,
			Quantity:        2.5,
			Sum:             2750,
			ProductCodeData: &productCodeData,
		}
		updatedDbItem := &DbModel{
			Model:           gorm.Model{ID: 1},
			ProductID:       100,
			ReceiptID:       1,
			Price:           1100,
			Quantity:        2.5,
			Sum:             2750,
			ProductCodeData: &productCodeData,
		}
		updatedPbItem := &pb.Model{
			Id:        id,
			ProductId: 100,
			ReceiptId: 1,
			Price:     1100,
			Quantity:  2.5,
			Sum:       2750,
			//ProductCodeData: productCodeData,
		}

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
		productCodeData := `{"code": "123456", "updated": true}`

		pbItemToUpdate := &pb.Model{
			Id:        id,
			ProductId: 100,
			ReceiptId: 1,
			Price:     1100,
			Quantity:  2.5,
			Sum:       2750,
			//ProductCodeData: productCodeData,
		}

		dbItemToUpdate := &DbModel{
			Model:           gorm.Model{ID: 1},
			ProductID:       100,
			ReceiptID:       1,
			Price:           1100,
			Quantity:        2.5,
			Sum:             2750,
			ProductCodeData: &productCodeData,
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
