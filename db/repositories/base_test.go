package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/db_types"
)

// Mock models
type TestModel struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string
	Age       int
	TestItems []TestItem `gorm:"foreignKey:TestModelID"`
}

type TestItem struct {
	ID          uint64 `gorm:"primarykey"`
	Description string
	TestModelID uint64
}

// Mock repository with HasAssociations
type TestRepository struct {
	BaseRepository[TestModel]
}

func (r *TestRepository) UpdateAssociations(db *gorm.DB, item *TestModel, updatedData *TestModel) error {
	// Clear old associations and replace with new ones
	if err := db.Where("test_model_id = ?", item.ID).Delete(&TestItem{}).Error; err != nil {
		return err
	}
	// Create new associations
	for i := range updatedData.TestItems {
		updatedData.TestItems[i].TestModelID = item.ID
		if err := db.Create(&updatedData.TestItems[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// Mock DBProvider
type TestDBProvider struct {
	DB *gorm.DB
}

func (p *TestDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	return p.DB, nil
}

// Setup function for tests
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	// Auto migrate the tables
	err = db.AutoMigrate(&TestModel{}, &TestItem{})
	require.NoError(t, err)

	return db
}

func setupTestRepo(db *gorm.DB) *TestRepository {
	repo := &TestRepository{
		BaseRepository: BaseRepository[TestModel]{
			DBProvider: &TestDBProvider{DB: db},
			Preloads:   []string{"TestItems"},
		},
	}
	repo.Self = repo
	return repo
}

// Filter for testing
func nameFilter(name string) db_types.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}

// Test for GetList method
func TestBaseRepository_GetList(t *testing.T) {
	db := setupTestDB(t)
	repo := setupTestRepo(db)

	// Create test data
	testData := []TestModel{
		{Name: "Alice", Age: 30, TestItems: []TestItem{{Description: "Item 1"}, {Description: "Item 2"}}},
		{Name: "Bob", Age: 25, TestItems: []TestItem{{Description: "Item 3"}}},
		{Name: "Charlie", Age: 35, TestItems: []TestItem{{Description: "Item 4"}, {Description: "Item 5"}}},
	}

	for i := range testData {
		_, err := repo.Create(context.Background(), &testData[i])
		require.NoError(t, err)
	}

	t.Run("GetList with no filters", func(t *testing.T) {
		items, count, err := repo.GetList(context.Background(), nil, db_types.SortData{}, db_types.PaginationData{})
		require.NoError(t, err)
		assert.Equal(t, int64(3), count)
		assert.Len(t, items, 3)
	})

	t.Run("GetList with filter", func(t *testing.T) {
		items, count, err := repo.GetList(context.Background(), []db_types.FilterFunc{nameFilter("Alice")}, db_types.SortData{}, db_types.PaginationData{})
		require.NoError(t, err)
		assert.Equal(t, int64(1), count)
		assert.Len(t, items, 1)
		assert.Equal(t, "Alice", items[0].Name)
	})

	t.Run("GetList with sort", func(t *testing.T) {
		items, count, err := repo.GetList(context.Background(), nil, db_types.SortData{SortField: "age", SortOrder: "desc"}, db_types.PaginationData{})
		require.NoError(t, err)
		assert.Equal(t, int64(3), count)
		assert.Len(t, items, 3)
		assert.Equal(t, "Charlie", items[0].Name) // Charlie has the highest age (35)
	})

	t.Run("GetList with pagination", func(t *testing.T) {
		items, count, err := repo.GetList(context.Background(), nil, db_types.SortData{SortField: "name", SortOrder: "asc"}, db_types.PaginationData{Page: 1, PageSize: 2})
		require.NoError(t, err)
		assert.Equal(t, int64(3), count) // Total count should still be 3
		assert.Len(t, items, 2)          // But only 2 items returned due to pagination
		assert.Equal(t, "Alice", items[0].Name)
		assert.Equal(t, "Bob", items[1].Name)
	})
}

// Test for Get method
func TestBaseRepository_Get(t *testing.T) {
	db := setupTestDB(t)
	repo := setupTestRepo(db)

	// Create test data
	testModel := TestModel{Name: "Alice", Age: 30, TestItems: []TestItem{{Description: "Item 1"}, {Description: "Item 2"}}}
	created, err := repo.Create(context.Background(), &testModel)
	require.NoError(t, err)

	t.Run("Get existing item", func(t *testing.T) {
		item, err := repo.Get(context.Background(), created.ID)
		require.NoError(t, err)
		assert.Equal(t, created.ID, item.ID)
		assert.Equal(t, "Alice", item.Name)
		assert.Equal(t, 30, item.Age)
		assert.Len(t, item.TestItems, 2)
	})

	t.Run("Get non-existing item", func(t *testing.T) {
		_, err := repo.Get(context.Background(), 9999)
		assert.Error(t, err)
	})
}

// Test for Create method
func TestBaseRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := setupTestRepo(db)

	t.Run("Create simple item", func(t *testing.T) {
		testModel := TestModel{Name: "Alice", Age: 30}
		created, err := repo.Create(context.Background(), &testModel)
		require.NoError(t, err)
		assert.NotZero(t, created.ID)
		assert.Equal(t, "Alice", created.Name)
		assert.Equal(t, 30, created.Age)
	})

	t.Run("Create item with associations", func(t *testing.T) {
		testModel := TestModel{
			Name: "Bob",
			Age:  25,
			TestItems: []TestItem{
				{Description: "Item 1"},
				{Description: "Item 2"},
			},
		}
		created, err := repo.Create(context.Background(), &testModel)
		require.NoError(t, err)
		assert.NotZero(t, created.ID)
		assert.Equal(t, "Bob", created.Name)
		assert.Len(t, created.TestItems, 2)

		// Verify that associations have the correct parent ID
		for _, item := range created.TestItems {
			assert.Equal(t, created.ID, item.TestModelID)
		}
	})
}

// Test for Update method
func TestBaseRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := setupTestRepo(db)

	// Create initial test data
	testModel := TestModel{
		Name: "Alice",
		Age:  30,
		TestItems: []TestItem{
			{Description: "Item 1"},
			{Description: "Item 2"},
		},
	}
	created, err := repo.Create(context.Background(), &testModel)
	require.NoError(t, err)

	t.Run("Update basic fields", func(t *testing.T) {
		updatedModel := TestModel{
			Name:      "Alice Updated",
			Age:       31,
			TestItems: created.TestItems, // Сохраняем существующие ассоциации
		}
		updated, err := repo.Update(context.Background(), created.ID, &updatedModel)
		require.NoError(t, err)
		assert.Equal(t, created.ID, updated.ID)
		assert.Equal(t, "Alice Updated", updated.Name)
		assert.Equal(t, 31, updated.Age)
		assert.Len(t, updated.TestItems, 2) // Associations should remain
	})

	t.Run("Update with associations", func(t *testing.T) {
		updatedModel := TestModel{
			Name: "Alice Changed Again",
			Age:  32,
			TestItems: []TestItem{
				{Description: "New Item 1"},
				{Description: "New Item 2"},
				{Description: "New Item 3"},
			},
		}
		updated, err := repo.Update(context.Background(), created.ID, &updatedModel)
		require.NoError(t, err)
		assert.Equal(t, created.ID, updated.ID)
		assert.Equal(t, "Alice Changed Again", updated.Name)
		assert.Equal(t, 32, updated.Age)
		assert.Len(t, updated.TestItems, 3) // Should have 3 new associations

		// Check descriptions to verify associations were updated
		descriptions := make([]string, len(updated.TestItems))
		for i, item := range updated.TestItems {
			descriptions[i] = item.Description
		}
		assert.Contains(t, descriptions, "New Item 1")
		assert.Contains(t, descriptions, "New Item 2")
		assert.Contains(t, descriptions, "New Item 3")
	})

	t.Run("Update non-existing item", func(t *testing.T) {
		updatedModel := TestModel{Name: "Should Not Update"}
		_, err := repo.Update(context.Background(), 9999, &updatedModel)
		assert.Error(t, err)
	})
}

// Test for Delete method
func TestBaseRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := setupTestRepo(db)

	// Create test data
	testModel := TestModel{Name: "To Be Deleted", Age: 40}
	created, err := repo.Create(context.Background(), &testModel)
	require.NoError(t, err)

	t.Run("Delete existing item", func(t *testing.T) {
		err := repo.Delete(context.Background(), created.ID)
		require.NoError(t, err)

		// Verify it's deleted
		_, err = repo.Get(context.Background(), created.ID)
		assert.Error(t, err)
	})

	t.Run("Delete non-existing item", func(t *testing.T) {
		err := repo.Delete(context.Background(), 9999)
		assert.Error(t, err)
	})
}

// Test helper methods
func TestBaseRepository_HelperMethods(t *testing.T) {
	db := setupTestDB(t)
	repo := setupTestRepo(db)

	// Create test data
	for i := 0; i < 5; i++ {
		_, err := repo.Create(context.Background(), &TestModel{
			Name: "Test",
			Age:  20 + i,
		})
		require.NoError(t, err)
	}

	t.Run("addSort asc", func(t *testing.T) {
		sort := db_types.SortData{SortField: "age", SortOrder: "asc"}
		items, _, err := repo.GetList(context.Background(), nil, sort, db_types.PaginationData{})
		require.NoError(t, err)
		assert.Equal(t, 20, items[0].Age) // Youngest first
		assert.Equal(t, 24, items[4].Age) // Oldest last
	})

	t.Run("addSort desc", func(t *testing.T) {
		sort := db_types.SortData{SortField: "age", SortOrder: "desc"}
		items, _, err := repo.GetList(context.Background(), nil, sort, db_types.PaginationData{})
		require.NoError(t, err)
		assert.Equal(t, 24, items[0].Age) // Oldest first
		assert.Equal(t, 20, items[4].Age) // Youngest last
	})

	t.Run("addPagination first page", func(t *testing.T) {
		pagination := db_types.PaginationData{Page: 1, PageSize: 2}
		items, total, err := repo.GetList(context.Background(), nil, db_types.SortData{SortField: "age", SortOrder: "asc"}, pagination)
		require.NoError(t, err)
		assert.Equal(t, int64(5), total) // Total count should be 5
		assert.Len(t, items, 2)          // But only 2 items per page
		assert.Equal(t, 20, items[0].Age)
		assert.Equal(t, 21, items[1].Age)
	})

	t.Run("addPagination second page", func(t *testing.T) {
		pagination := db_types.PaginationData{Page: 2, PageSize: 2}
		items, total, err := repo.GetList(context.Background(), nil, db_types.SortData{SortField: "age", SortOrder: "asc"}, pagination)
		require.NoError(t, err)
		assert.Equal(t, int64(5), total) // Total count should be 5
		assert.Len(t, items, 2)          // But only 2 items per page
		assert.Equal(t, 22, items[0].Age)
		assert.Equal(t, 23, items[1].Age)
	})
}

// Test preload functionality
func TestBaseRepository_Preload(t *testing.T) {
	db := setupTestDB(t)

	// Create a repo without preloads
	repoWithoutPreloads := &TestRepository{
		BaseRepository: BaseRepository[TestModel]{
			DBProvider: &TestDBProvider{DB: db},
			Preloads:   []string{}, // No preloads
		},
	}
	repoWithoutPreloads.Self = repoWithoutPreloads

	// Create a repo with preloads
	repoWithPreloads := &TestRepository{
		BaseRepository: BaseRepository[TestModel]{
			DBProvider: &TestDBProvider{DB: db},
			Preloads:   []string{"TestItems"}, // With preloads
		},
	}
	repoWithPreloads.Self = repoWithPreloads

	// Create test data
	testModel := TestModel{
		Name: "Preload Test",
		Age:  25,
		TestItems: []TestItem{
			{Description: "Preload Item 1"},
			{Description: "Preload Item 2"},
		},
	}
	created, err := repoWithPreloads.Create(context.Background(), &testModel)
	require.NoError(t, err)

	t.Run("Get without preloads", func(t *testing.T) {
		item, err := repoWithoutPreloads.Get(context.Background(), created.ID)
		require.NoError(t, err)
		assert.Equal(t, "Preload Test", item.Name)
		assert.Len(t, item.TestItems, 0) // No items should be preloaded
	})

	t.Run("Get with preloads", func(t *testing.T) {
		item, err := repoWithPreloads.Get(context.Background(), created.ID)
		require.NoError(t, err)
		assert.Equal(t, "Preload Test", item.Name)
		assert.Len(t, item.TestItems, 2) // Items should be preloaded
	})
}
