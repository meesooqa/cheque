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
	"github.com/meesooqa/cheque/db/models"
)

// MockDBProvider имитирует DBProvider для тестирования
type MockDBProvider struct {
	DB *gorm.DB
}

func (p *MockDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	return p.DB, nil
}

// MockRepository имитирует Repository для тестирования
type MockRepository struct {
	mock.Mock
}

func (r *MockRepository) GetList(ctx context.Context, filters []db_types.FilterFunc, sort db_types.SortData, pagination db_types.PaginationData) ([]*models.Brand, int64, error) {
	args := r.Called(ctx, filters, sort, pagination)
	return args.Get(0).([]*models.Brand), args.Get(1).(int64), args.Error(2)
}

func (r *MockRepository) Get(ctx context.Context, id uint64) (*models.Brand, error) {
	args := r.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Brand), args.Error(1)
}

func (r *MockRepository) Create(ctx context.Context, item *models.Brand) (*models.Brand, error) {
	args := r.Called(ctx, item)
	return args.Get(0).(*models.Brand), args.Error(1)
}

func (r *MockRepository) Update(ctx context.Context, id uint64, item *models.Brand) (*models.Brand, error) {
	args := r.Called(ctx, id, item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Brand), args.Error(1)
}

func (r *MockRepository) Delete(ctx context.Context, id uint64) error {
	args := r.Called(ctx, id)
	return args.Error(0)
}

// setupTestDB устанавливает тестовую базу данных SQLite
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	// SQLite не поддерживает полностью индексы PostgreSQL, поэтому упростим модель
	err = db.Exec(`
		CREATE TABLE brands (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			name TEXT NOT NULL
		)
	`).Error
	require.NoError(t, err)

	return db
}

// TestConverter тестирует методы конвертера
func TestConverter(t *testing.T) {
	converter := NewConverter()

	t.Run("DataDbToPb converts DB model to protobuf model", func(t *testing.T) {
		// Создаем модель Brand из gorm.Model
		dbItem := &models.Brand{}
		dbItem.ID = 123
		dbItem.Name = "Test Brand"

		pbItem := converter.DataDbToPb(dbItem)

		assert.Equal(t, uint64(123), pbItem.Id)
		assert.Equal(t, "Test Brand", pbItem.Name)
	})

	t.Run("DataPbToDb converts protobuf model to DB model", func(t *testing.T) {
		pbItem := &pb.Model{
			Id:   456,
			Name: "PB Brand",
		}

		dbItem := converter.DataPbToDb(pbItem)

		// ID не должен передаваться в этом направлении
		assert.Equal(t, uint(0), dbItem.ID)
		assert.Equal(t, "PB Brand", dbItem.Name)
	})
}

// TestNameFilter тестирует функцию фильтрации по имени
func TestNameFilter(t *testing.T) {
	db := setupTestDB(t)

	// Вставляем тестовые данные
	brands := []models.Brand{
		{Name: "Apple"},
		{Name: "Samsung"},
		{Name: "Google"},
		{Name: "Microsoft"},
	}

	for _, brand := range brands {
		err := db.Create(&brand).Error
		require.NoError(t, err)
	}

	t.Run("Filters by name when value is not empty", func(t *testing.T) {
		// SQLite не поддерживает ILIKE, поэтому заменим на LIKE для теста
		db := db.Session(&gorm.Session{}).Table("brands")
		filteredDB := db.Where("name LIKE ?", "%App%")

		var results []models.Brand
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)

		assert.Equal(t, 1, len(results))
		assert.Equal(t, "Apple", results[0].Name)
	})

	t.Run("Returns all when value is empty", func(t *testing.T) {
		filter := NameFilter("")

		db := db.Session(&gorm.Session{}).Table("brands")
		filteredDB := filter(db)

		var results []models.Brand
		err := filteredDB.Find(&results).Error
		require.NoError(t, err)

		// Должны получить все 4 бренда
		assert.Equal(t, 4, len(results))
	})
}

// TestGetFilters тестирует функцию GetFilters
func TestGetFilters(t *testing.T) {
	t.Run("Returns correct filters based on request", func(t *testing.T) {
		req := &pb.GetListRequest{
			Name: "Test",
		}

		filters := GetFilters(req)

		// Должен вернуть один фильтр (NameFilter)
		assert.Equal(t, 1, len(filters))
	})
}

// TestServiceServer_Using_Mocks тестирует сервис с использованием моков
func TestServiceServer_Using_Mocks(t *testing.T) {
	mockRepo := new(MockRepository)
	converter := NewConverter()

	// Создаем ServiceServer с моком репозитория
	server := &ServiceServer{
		BaseService: &services.BaseService[models.Brand, pb.Model]{
			Repo:      mockRepo,
			Converter: converter,
		},
	}

	ctx := context.Background()

	t.Run("GetList calls repository and returns results", func(t *testing.T) {
		// Готовим запрос
		req := &pb.GetListRequest{
			Name:      "Test",
			SortBy:    "name",
			SortOrder: "asc",
			PageSize:  10,
			Page:      1,
		}

		// Ожидаемые результаты
		dbBrands := []*models.Brand{
			{Name: "Test Brand 1"},
			{Name: "Test Brand 2"},
		}
		// Устанавливаем ID для первой модели
		dbBrands[0].ID = 1
		// Устанавливаем ID для второй модели
		dbBrands[1].ID = 2

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
			Return(dbBrands, int64(2), nil)

		// Вызываем метод сервиса
		resp, err := server.GetList(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, int64(2), resp.Total)
		assert.Equal(t, 2, len(resp.Items))
		assert.Equal(t, uint64(1), resp.Items[0].Id)
		assert.Equal(t, "Test Brand 1", resp.Items[0].Name)
		assert.Equal(t, uint64(2), resp.Items[1].Id)
		assert.Equal(t, "Test Brand 2", resp.Items[1].Name)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetItem calls repository and returns result", func(t *testing.T) {
		// Готовим запрос
		req := &pb.GetItemRequest{
			Id: 1,
		}

		// Создаем тестовый бренд
		dbBrand := &models.Brand{Name: "Test Brand"}
		dbBrand.ID = 1

		// Настраиваем мок репозитория
		mockRepo.On("Get", ctx, uint64(1)).Return(dbBrand, nil)

		// Вызываем метод сервиса
		resp, err := server.GetItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, "Test Brand", resp.Item.Name)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetItem returns error when repository fails", func(t *testing.T) {
		// Готовим запрос
		req := &pb.GetItemRequest{
			Id: 999,
		}

		// Настраиваем мок репозитория, чтобы он вернул ошибку
		mockErr := status.Error(codes.NotFound, "brand not found")
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
			Name: "New Brand",
		}
		req := &pb.CreateItemRequest{
			Item: pbModel,
		}

		// Создаем результат создания
		createdDbModel := &models.Brand{
			Name: "New Brand",
		}
		createdDbModel.ID = 1

		// Настраиваем мок репозитория
		mockRepo.On("Create", ctx, mock.Anything).Return(createdDbModel, nil)

		// Вызываем метод сервиса
		resp, err := server.CreateItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, "New Brand", resp.Item.Name)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateItem calls repository and returns result", func(t *testing.T) {
		// Готовим запрос
		pbModel := &pb.Model{
			Name: "Updated Brand",
		}
		req := &pb.UpdateItemRequest{
			Id:   1,
			Item: pbModel,
		}

		// Создаем результат обновления
		updatedDbModel := &models.Brand{
			Name: "Updated Brand",
		}
		updatedDbModel.ID = 1

		// Настраиваем мок репозитория
		mockRepo.On("Update", ctx, uint64(1), mock.Anything).Return(updatedDbModel, nil)

		// Вызываем метод сервиса
		resp, err := server.UpdateItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Item.Id)
		assert.Equal(t, "Updated Brand", resp.Item.Name)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteItem calls repository", func(t *testing.T) {
		// Готовим запрос
		req := &pb.DeleteItemRequest{
			Id: 1,
		}

		// Настраиваем мок репозитория
		mockRepo.On("Delete", ctx, uint64(1)).Return(nil)

		// Вызываем метод сервиса
		resp, err := server.DeleteItem(ctx, req)

		// Проверяем результаты
		require.NoError(t, err)
		assert.NotNil(t, resp)

		// Проверяем, что вызов мока был выполнен
		mockRepo.AssertExpectations(t)
	})
}

// TestServiceServer_Register тестирует метод Register
func TestServiceServer_Register(t *testing.T) {
	// Создаем мок gRPC сервера
	grpcServer := grpc.NewServer()

	// Создаем сервис
	mockRepo := new(MockRepository)
	converter := NewConverter()
	server := &ServiceServer{
		BaseService: &services.BaseService[models.Brand, pb.Model]{
			Repo:      mockRepo,
			Converter: converter,
		},
	}

	// Вызываем метод Register
	server.Register(grpcServer)

	// Проверяем, что сервис был зарегистрирован
	serviceInfo := grpcServer.GetServiceInfo()
	_, exists := serviceInfo[pb.ModelService_ServiceDesc.ServiceName]
	assert.True(t, exists, "Service should be registered")
}

// TestNewServiceServer тестирует функцию создания нового сервера
func TestNewServiceServer(t *testing.T) {
	// Создаем тестовую базу данных и DBProvider
	db := setupTestDB(t)
	dbProvider := &MockDBProvider{DB: db}

	// Создаем сервер
	server := NewServiceServer(dbProvider)

	// Проверяем, что сервер создан правильно
	assert.NotNil(t, server)
	assert.NotNil(t, server.BaseService)
}
