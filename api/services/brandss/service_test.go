package brandss

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/gen/pb/brandpb/v1"
	"github.com/meesooqa/cheque/api/services"
	"github.com/meesooqa/cheque/db/db_types"
)

// setupTestDB sets up test database SQLite
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&DbModel{})
	require.NoError(t, err)
	return db
}

// TestConverter tests Converter methods
func TestConverter(t *testing.T) {
	converter := NewConverter()
	t.Run("DataDbToPb converts DB model to protobuf model", func(t *testing.T) {
		dbItem := &DbModel{Name: "Test Item"}
		dbItem.ID = 123
		pbItem := converter.DataDbToPb(dbItem)
		assert.Equal(t, uint64(123), pbItem.Id)
		assert.Equal(t, "Test Item", pbItem.Name)
	})
	t.Run("DataPbToDb converts protobuf model to DB model", func(t *testing.T) {
		pbItem := &pb.Model{Id: 456, Name: "PB Item"}
		dbItem := converter.DataPbToDb(pbItem)
		assert.Equal(t, uint(0), dbItem.ID)
		assert.Equal(t, "PB Item", dbItem.Name)
	})
}

// TestNameFilter tests NameFilter
func TestNameFilter(t *testing.T) {
	db := setupTestDB(t)
	brands := []DbModel{
		{Name: "Apple"},
		{Name: "Samsung"},
		{Name: "Google"},
		{Name: "Microsoft"},
	}
	for _, brand := range brands {
		err := db.Create(&brand).Error
		require.NoError(t, err)
	}
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	t.Run("Filters by name when value is not empty", func(t *testing.T) {
		filter := NameFilter("App")
		db := db.Session(&gorm.Session{}).Table("brands")
		filteredDB := filter(db)
		var results []DbModel
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)
		assert.Equal(t, 1, len(results))
		assert.Equal(t, "Apple", results[0].Name)
	})
	t.Run("Returns all when value is empty", func(t *testing.T) {
		filter := NameFilter("")
		db := db.Session(&gorm.Session{}).Table("brands")
		filteredDB := filter(db)
		var results []DbModel
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)
		assert.Equal(t, 4, len(results))
	})
}

// TestServiceServer_Using_Mocks tests ServiceServer using mocks
func TestServiceServer_Using_Mocks(t *testing.T) {
	mockRepo := new(MockRepository)
	converter := NewConverter()
	filterProvider := NewFilterProvider()
	server := &ServiceServer{
		BaseService: &services.BaseService[DbModel, pb.Model, pb.GetListRequest]{
			Repo:           mockRepo,
			Converter:      converter,
			FilterProvider: filterProvider,
		},
	}
	ctx := context.Background()
	t.Run("GetList calls repository and returns results", func(t *testing.T) {
		req := &pb.GetListRequest{
			Name:      "Test",
			SortBy:    "name",
			SortOrder: "asc",
			PageSize:  10,
			Page:      1,
		}
		dbBrands := []*DbModel{
			{Name: "Test Item 1"},
			{Name: "Test Item 2"},
		}
		dbBrands[0].ID = 1
		dbBrands[1].ID = 2
		expectedSort := db_types.SortData{
			SortField: "name",
			SortOrder: "asc",
		}
		expectedPagination := db_types.PaginationData{
			Page:     1,
			PageSize: 10,
		}
		mockRepo.On("GetList", ctx, mock.Anything, expectedSort, expectedPagination).
			Return(dbBrands, int64(2), nil)
		resp, err := server.GetList(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, int64(2), resp.Total)
		assert.Equal(t, 2, len(resp.Items))
		assert.Equal(t, uint64(1), resp.Items[0].Id)
		assert.Equal(t, "Test Item 1", resp.Items[0].Name)
		assert.Equal(t, uint64(2), resp.Items[1].Id)
		assert.Equal(t, "Test Item 2", resp.Items[1].Name)
		mockRepo.AssertExpectations(t)
	})
	t.Run("GetItem calls repository and returns result", func(t *testing.T) {
		req := &pb.GetItemRequest{
			Id: 1,
		}
		dbBrand := &DbModel{Name: "Test Item"}
		dbBrand.ID = 1
		mockRepo.On("Get", ctx, uint64(1)).Return(dbBrand, nil)
		resp, err := server.GetItem(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, "Test Item", resp.Item.Name)
		mockRepo.AssertExpectations(t)
	})
	t.Run("GetItem returns error when repository fails", func(t *testing.T) {
		req := &pb.GetItemRequest{
			Id: 999,
		}
		mockErr := status.Error(codes.NotFound, "item not found")
		mockRepo.On("Get", ctx, uint64(999)).Return(nil, mockErr)
		_, err := server.GetItem(ctx, req)
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		mockRepo.AssertExpectations(t)
	})
	t.Run("CreateItem calls repository and returns result", func(t *testing.T) {
		pbModel := &pb.Model{
			Name: "New Item",
		}
		req := &pb.CreateItemRequest{
			Item: pbModel,
		}
		createdDbModel := &DbModel{
			Name: "New Item",
		}
		createdDbModel.ID = 1
		mockRepo.On("Create", ctx, mock.Anything).Return(createdDbModel, nil)
		resp, err := server.CreateItem(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, "New Item", resp.Item.Name)
		mockRepo.AssertExpectations(t)
	})
	t.Run("UpdateItem calls repository and returns result", func(t *testing.T) {
		pbModel := &pb.Model{
			Name: "Updated Item",
		}
		req := &pb.UpdateItemRequest{
			Id:   1,
			Item: pbModel,
		}
		updatedDbModel := &DbModel{
			Name: "Updated Item",
		}
		updatedDbModel.ID = 1
		mockRepo.On("Update", ctx, uint64(1), mock.Anything).Return(updatedDbModel, nil)
		resp, err := server.UpdateItem(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, "Updated Item", resp.Item.Name)
		mockRepo.AssertExpectations(t)
	})
	t.Run("DeleteItem calls repository", func(t *testing.T) {
		req := &pb.DeleteItemRequest{
			Id: 1,
		}
		mockRepo.On("Delete", ctx, uint64(1)).Return(nil)
		resp, err := server.DeleteItem(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		mockRepo.AssertExpectations(t)
	})
}

// TestServiceServer_Register tests Register
func TestServiceServer_Register(t *testing.T) {
	grpcServer := grpc.NewServer()
	mockRepo := new(MockRepository)
	converter := NewConverter()
	filterProvider := NewFilterProvider()
	server := &ServiceServer{
		BaseService: &services.BaseService[DbModel, pb.Model, pb.GetListRequest]{
			Repo:           mockRepo,
			Converter:      converter,
			FilterProvider: filterProvider,
		},
	}
	server.Register(grpcServer)
	serviceInfo := grpcServer.GetServiceInfo()
	_, exists := serviceInfo[pb.ModelService_ServiceDesc.ServiceName]
	assert.True(t, exists, "Service should be registered")
}

// TestNewServiceServer tests NewServiceServer
func TestNewServiceServer(t *testing.T) {
	db := setupTestDB(t)
	dbProvider := &MockDBProvider{DB: db}
	server := NewServiceServer(dbProvider)
	assert.NotNil(t, server)
	assert.NotNil(t, server.BaseService)
}
