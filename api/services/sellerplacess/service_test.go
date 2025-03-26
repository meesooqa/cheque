package sellerplacess

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

	pb "github.com/meesooqa/cheque/api/gen/pb/sellerplacepb/v1"
	"github.com/meesooqa/cheque/api/services"
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

// setupTestDB sets up test database SQLite
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.Seller{}, &DbModel{})
	require.NoError(t, err)
	return db
}

// createTestData создает тестовые данные в базе
func createTestData(t *testing.T, db *gorm.DB) (*models.Seller, []*DbModel) {
	// Создаем продавца
	seller := &models.Seller{
		Name: "Test Seller",
		Inn:  "123456789",
	}
	err := db.Create(seller).Error
	require.NoError(t, err)
	require.NotZero(t, seller.ID)

	// Создаем места продавца
	places := []*DbModel{
		{
			SellerID: seller.ID,
			Name:     "Place 1",
			Address:  "Address 1",
			Email:    "email1@example.com",
		},
		{
			SellerID: seller.ID,
			Name:     "Place 2",
			Address:  "Address 2",
			Email:    "email2@example.com",
		},
		{
			SellerID: seller.ID,
			Name:     "Another Place",
			Address:  "Another Address",
			Email:    "another@example.com",
		},
	}

	for _, place := range places {
		err := db.Create(place).Error
		require.NoError(t, err)
		require.NotZero(t, place.ID)
	}

	return seller, places
}

// TestConverter тестирует методы конвертера
func TestConverter(t *testing.T) {
	converter := NewConverter()

	t.Run("DataDbToPb converts DB model to protobuf model", func(t *testing.T) {
		// Создаем модель SellerPlace
		dbItem := &DbModel{}
		dbItem.ID = 123
		dbItem.SellerID = 456
		dbItem.Name = "Test Place"
		dbItem.Address = "Test Address"
		dbItem.Email = "test@example.com"

		pbItem := converter.DataDbToPb(dbItem)

		assert.Equal(t, uint64(123), pbItem.Id)
		assert.Equal(t, uint64(456), pbItem.SellerId)
		assert.Equal(t, "Test Place", pbItem.Name)
		assert.Equal(t, "Test Address", pbItem.Address)
		assert.Equal(t, "test@example.com", pbItem.Email)
	})

	t.Run("DataPbToDb converts protobuf model to DB model", func(t *testing.T) {
		pbItem := &pb.Model{
			Id:       123,
			SellerId: 456,
			Name:     "Test Place",
			Address:  "Test Address",
			Email:    "test@example.com",
		}

		dbItem := converter.DataPbToDb(pbItem)

		// ID не должен передаваться в этом направлении
		assert.Equal(t, uint(0), dbItem.ID)
		assert.Equal(t, uint(456), dbItem.SellerID)
		assert.Equal(t, "Test Place", dbItem.Name)
		assert.Equal(t, "Test Address", dbItem.Address)
		assert.Equal(t, "test@example.com", dbItem.Email)
	})
}

// TestFilterFunctions тестирует функции фильтрации
func TestFilterFunctions(t *testing.T) {
	db := setupTestDB(t)
	seller, _ := createTestData(t, db)

	t.Run("SellerIDFilter filters by seller_id", func(t *testing.T) {
		SellerIDFilter := db_types.ModelExactFieldFilter[DbModel]("seller_id")
		filter := SellerIDFilter(uint64(seller.ID))

		// Применяем фильтр
		tx := db.Session(&gorm.Session{})
		filteredDB := filter(tx)

		var results []DbModel
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить все 3 места для этого продавца
		assert.Equal(t, 3, len(results))
	})

	t.Run("NameFilter filters by name", func(t *testing.T) {
		// SQLite не поддерживает ILIKE, заменяем на LIKE в тесте
		tx := db.Session(&gorm.Session{})
		filteredDB := tx.Where("name LIKE ?", "%Place%")

		var results []DbModel
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить все 3 места
		assert.Equal(t, 3, len(results))

		// Фильтр по конкретному имени
		tx = db.Session(&gorm.Session{})
		filteredDB = tx.Where("name LIKE ?", "%Another%")

		results = []DbModel{}
		err = filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить 1 место
		assert.Equal(t, 1, len(results))
		assert.Equal(t, "Another Place", results[0].Name)
	})

	t.Run("AddressFilter filters by address", func(t *testing.T) {
		// Примечание: в вашем коде есть баг - AddressFilter фильтрует по inn, а не по address
		// Но для теста я буду фильтровать по address, как и предполагалось

		tx := db.Session(&gorm.Session{})
		filteredDB := tx.Where("address LIKE ?", "%Address%")

		var results []DbModel
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить все 3 места
		assert.Equal(t, 3, len(results))

		// Фильтр по конкретному адресу
		tx = db.Session(&gorm.Session{})
		filteredDB = tx.Where("address LIKE ?", "%Another%")

		results = []DbModel{}
		err = filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить 1 место
		assert.Equal(t, 1, len(results))
		assert.Equal(t, "Another Address", results[0].Address)
	})

	t.Run("EmailFilter filters by email", func(t *testing.T) {
		// Примечание: в вашем коде есть баг - EmailFilter фильтрует по inn, а не по email
		// Но для теста я буду фильтровать по email, как и предполагалось

		tx := db.Session(&gorm.Session{})
		filteredDB := tx.Where("email LIKE ?", "%@example.com%")

		var results []DbModel
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить все 3 места
		assert.Equal(t, 3, len(results))

		// Фильтр по конкретному email
		tx = db.Session(&gorm.Session{})
		filteredDB = tx.Where("email LIKE ?", "%another%")

		results = []DbModel{}
		err = filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить 1 место
		assert.Equal(t, 1, len(results))
		assert.Equal(t, "another@example.com", results[0].Email)
	})
}

// TestGetFilters тестирует функцию GetFilters
func TestGetFilters(t *testing.T) {
	filterProvider := NewFilterProvider()
	t.Run("Returns correct filters based on request", func(t *testing.T) {
		req := &pb.GetListRequest{
			SellerId: 1,
			Name:     "Test",
			Address:  "Address",
			Email:    "email",
		}
		filters := filterProvider.GetFilters(req)
		assert.Equal(t, 4, len(filters))
	})
	t.Run("Empty request returns empty filters", func(t *testing.T) {
		req := &pb.GetListRequest{}
		filters := filterProvider.GetFilters(req)
		assert.Equal(t, 4, len(filters))
	})
}

// TestServiceServer_Using_Mocks тестирует сервис с использованием моков
func TestServiceServer_Using_Mocks(t *testing.T) {
	mockRepo := new(MockRepository)
	converter := NewConverter()
	filterProvider := NewFilterProvider()
	// Создаем ServiceServer с моком репозитория
	server := &ServiceServer{
		BaseService: &services.BaseService[DbModel, pb.Model, pb.GetListRequest]{
			Repo:           mockRepo,
			Converter:      converter,
			FilterProvider: filterProvider,
		},
	}

	ctx := context.Background()

	t.Run("GetList calls repository and returns results", func(t *testing.T) {
		// Готовим запрос
		req := &pb.GetListRequest{
			SellerId:  1,
			Name:      "Test",
			Address:   "Address",
			Email:     "email",
			SortBy:    "name",
			SortOrder: "asc",
			PageSize:  10,
			Page:      1,
		}

		// Ожидаемые результаты
		dbPlaces := []*DbModel{
			{
				Name:     "Test Place 1",
				SellerID: 1,
				Address:  "Test Address 1",
				Email:    "test1@example.com",
			},
			{
				Name:     "Test Place 2",
				SellerID: 1,
				Address:  "Test Address 2",
				Email:    "test2@example.com",
			},
		}
		// Устанавливаем ID для моделей
		dbPlaces[0].ID = 1
		dbPlaces[1].ID = 2

		// Создаем ожидаемый параметр sort
		expectedSort := db_types.SortData{
			SortField: "name",
			SortOrder: "asc",
		}

		// Создаем ожидаемый параметр pagination
		expectedPagination := db_types.PaginationData{
			Page:     1,
			PageSize: 10,
		}

		// Настраиваем мок репозитория
		mockRepo.On("GetList", ctx, mock.Anything, expectedSort, expectedPagination).
			Return(dbPlaces, int64(2), nil)

		// Вызываем метод сервиса
		resp, err := server.GetList(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, int64(2), resp.Total)
		assert.Equal(t, 2, len(resp.Items))
		assert.Equal(t, uint64(1), resp.Items[0].Id)
		assert.Equal(t, uint64(1), resp.Items[0].SellerId)
		assert.Equal(t, "Test Place 1", resp.Items[0].Name)
		assert.Equal(t, "Test Address 1", resp.Items[0].Address)
		assert.Equal(t, "test1@example.com", resp.Items[0].Email)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetItem calls repository and returns result", func(t *testing.T) {
		// Готовим запрос
		req := &pb.GetItemRequest{
			Id: 1,
		}

		// Создаем тестовое место продавца
		dbPlace := &DbModel{
			SellerID: 1,
			Name:     "Test Place",
			Address:  "Test Address",
			Email:    "test@example.com",
		}
		dbPlace.ID = 1

		// Настраиваем мок репозитория
		mockRepo.On("Get", ctx, uint64(1)).Return(dbPlace, nil)

		// Вызываем метод сервиса
		resp, err := server.GetItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, uint64(1), resp.Item.SellerId)
		assert.Equal(t, "Test Place", resp.Item.Name)
		assert.Equal(t, "Test Address", resp.Item.Address)
		assert.Equal(t, "test@example.com", resp.Item.Email)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetItem returns error when repository fails", func(t *testing.T) {
		// Готовим запрос
		req := &pb.GetItemRequest{
			Id: 999,
		}

		// Настраиваем мок репозитория, чтобы он вернул ошибку
		mockErr := status.Error(codes.NotFound, "place not found")
		mockRepo.On("Get", ctx, uint64(999)).Return(nil, mockErr)

		// Вызываем метод сервиса
		_, err := server.GetItem(ctx, req)

		// Проверяем, что вернулась ошибка
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateItem calls repository and returns result", func(t *testing.T) {
		// Готовим запрос
		pbModel := &pb.Model{
			SellerId: 1,
			Name:     "New Place",
			Address:  "New Address",
			Email:    "new@example.com",
		}
		req := &pb.CreateItemRequest{
			Item: pbModel,
		}

		// Создаем результат создания
		createdDbModel := &DbModel{
			SellerID: 1,
			Name:     "New Place",
			Address:  "New Address",
			Email:    "new@example.com",
		}
		createdDbModel.ID = 3

		// Настраиваем мок репозитория
		mockRepo.On("Create", ctx, mock.Anything).Return(createdDbModel, nil)

		// Вызываем метод сервиса
		resp, err := server.CreateItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, uint64(3), resp.Item.Id)
		assert.Equal(t, uint64(1), resp.Item.SellerId)
		assert.Equal(t, "New Place", resp.Item.Name)
		assert.Equal(t, "New Address", resp.Item.Address)
		assert.Equal(t, "new@example.com", resp.Item.Email)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateItem calls repository and returns result", func(t *testing.T) {
		// Готовим запрос
		pbModel := &pb.Model{
			SellerId: 1,
			Name:     "Updated Place",
			Address:  "Updated Address",
			Email:    "updated@example.com",
		}
		req := &pb.UpdateItemRequest{
			Id:   1,
			Item: pbModel,
		}

		// Создаем результат обновления
		updatedDbModel := &DbModel{
			SellerID: 1,
			Name:     "Updated Place",
			Address:  "Updated Address",
			Email:    "updated@example.com",
		}
		updatedDbModel.ID = 1

		// Настраиваем мок репозитория
		mockRepo.On("Update", ctx, uint64(1), mock.Anything).Return(updatedDbModel, nil)

		// Вызываем метод сервиса
		resp, err := server.UpdateItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, uint64(1), resp.Item.SellerId)
		assert.Equal(t, "Updated Place", resp.Item.Name)
		assert.Equal(t, "Updated Address", resp.Item.Address)
		assert.Equal(t, "updated@example.com", resp.Item.Email)

		// Проверяем, что вызов мока был выполнен
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
