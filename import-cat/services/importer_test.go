package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/models"
)

// SQLite DB Provider implementation for testing
type TestDBProvider struct {
	DB  *gorm.DB
	Err error
}

func (p *TestDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	if p.Err != nil {
		return nil, p.Err
	}
	return p.DB, nil
}

// Helper to create SQLite in-memory DB for testing
func setupTestDB(t *testing.T) *gorm.DB {
	// Каждый тест получает отдельную SQLite базу данных в памяти
	dbFile := fmt.Sprintf("file:%s?mode=memory", t.Name())
	t.Logf("Using database: %s", dbFile)

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	require.NoError(t, err, "Failed to open SQLite database")

	// Migrate the schema for Category and Product
	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	require.NoError(t, err, "Failed to migrate schema")

	return db
}

// Вспомогательная функция для отладки и проверки состояния БД
func dumpCategories(t *testing.T, db *gorm.DB) {
	var categories []models.Category
	result := db.Find(&categories)
	if result.Error != nil {
		t.Logf("Error fetching categories: %v", result.Error)
		return
	}

	t.Logf("Total categories in DB: %d", len(categories))
	for _, cat := range categories {
		t.Logf("Category: ID=%d, Name=%s, ParentID=%v", cat.ID, cat.Name, cat.ParentID)
	}
}

// Тест исправленной версии Save с правильным type assertion
func TestFixedSave(t *testing.T) {
	// Setup test DB
	db := setupTestDB(t)

	// Создаем тестовые данные - ОБРАТИТЕ ВНИМАНИЕ, мы используем указатель на GoogleProductTaxonomyItem
	items := []CategoriesReaderResultItem{
		&GoogleProductTaxonomyItem{
			ID:       100,
			Name:     "Test Category",
			FullPath: "Test Category",
			Level:    0,
		},
	}

	// Упрощенная версия Save с исправленным type assertion
	var count int64 = 0

	for _, resultItem := range items {
		// Правильное приведение типа к указателю, а не значению
		item, ok := resultItem.(*GoogleProductTaxonomyItem)
		if !ok {
			t.Logf("Item is not a *GoogleProductTaxonomyItem: %T", resultItem)
			continue
		}

		t.Logf("Processing item: %+v", item)

		// Создаем новую категорию
		newCategory := &models.Category{
			Name: item.Name,
		}

		result := db.Create(newCategory)
		if result.Error != nil {
			t.Errorf("Failed to create category: %v", result.Error)
			continue
		}

		t.Logf("Created category with ID: %d", newCategory.ID)
		count++
	}

	// Проверяем, что категории созданы
	assert.Equal(t, int64(1), count, "Should have created 1 category")

	// Проверяем состояние БД
	dumpCategories(t, db)

	var dbCount int64
	db.Model(&models.Category{}).Count(&dbCount)
	assert.Equal(t, int64(1), dbCount, "Should have 1 category in DB")

	// Проверяем, что категория найдена по имени
	var cat models.Category
	result := db.Where("name = ?", "Test Category").First(&cat)
	require.NoError(t, result.Error, "Should find the category")
	assert.Equal(t, "Test Category", cat.Name, "Category name should match")
}

// Тест, который пытается воспроизвести ошибку в Save
func TestBrokenSave(t *testing.T) {
	// Setup test DB
	db := setupTestDB(t)

	// Создаем тестовые данные - ОБРАТИТЕ ВНИМАНИЕ, мы используем указатель на GoogleProductTaxonomyItem
	items := []CategoriesReaderResultItem{
		&GoogleProductTaxonomyItem{
			ID:       100,
			Name:     "Test Category",
			FullPath: "Test Category",
			Level:    0,
		},
	}

	// Воспроизводим ошибку в Save - неправильное приведение типа
	var count int64 = 0

	for _, resultItem := range items {
		// НЕПРАВИЛЬНОЕ приведение типа - к значению, а не указателю!
		item, ok := resultItem.(GoogleProductTaxonomyItem) // <-- Вот проблема!
		if !ok {
			t.Logf("Type assertion failed: %T is not a GoogleProductTaxonomyItem", resultItem)
			continue
		}

		t.Logf("Processing item: %+v", item)

		// Создаем новую категорию
		newCategory := &models.Category{
			Name: item.Name,
		}

		result := db.Create(newCategory)
		if result.Error != nil {
			t.Errorf("Failed to create category: %v", result.Error)
			continue
		}

		t.Logf("Created category with ID: %d", newCategory.ID)
		count++
	}

	// Этот счетчик должен быть 0, так как type assertion не пройдет
	assert.Equal(t, int64(0), count, "Should NOT have created any categories due to type assertion failure")

	// Проверяем состояние БД
	dumpCategories(t, db)

	var dbCount int64
	db.Model(&models.Category{}).Count(&dbCount)
	assert.Equal(t, int64(0), dbCount, "Should have 0 categories in DB")
}

// Тест улучшенной версии Save с правильным type assertion
func TestImprovedSave(t *testing.T) {
	// Setup test DB
	db := setupTestDB(t)

	// Улучшенная фиксированная версия Save
	fixedSave := func(ctx context.Context, items []CategoriesReaderResultItem) error {
		for _, resultItem := range items {
			// Правильное приведение типа к указателю
			item, ok := resultItem.(*GoogleProductTaxonomyItem)
			if !ok {
				t.Logf("Item is not a *GoogleProductTaxonomyItem: %T", resultItem)
				continue
			}

			// Создаем новую категорию
			newCategory := &models.Category{
				Name: item.Name,
			}

			result := db.Create(newCategory)
			if result.Error != nil {
				return fmt.Errorf("failed to create category: %v", result.Error)
			}

			t.Logf("Created category with ID: %d, Name: %s", newCategory.ID, newCategory.Name)
		}

		return nil
	}

	// Создаем тестовые данные
	items := []CategoriesReaderResultItem{
		&GoogleProductTaxonomyItem{
			ID:       100,
			Name:     "Fixed Test Category",
			FullPath: "Fixed Test Category",
			Level:    0,
		},
	}

	// Выполняем импорт через улучшенную функцию
	err := fixedSave(context.Background(), items)
	require.NoError(t, err, "Fixed Save should succeed")

	// Проверяем состояние БД
	dumpCategories(t, db)

	var count int64
	db.Model(&models.Category{}).Count(&count)
	assert.Equal(t, int64(1), count, "Should have 1 category in DB")

	// Проверяем, что категория найдена по имени
	var cat models.Category
	result := db.Where("name = ?", "Fixed Test Category").First(&cat)
	require.NoError(t, result.Error, "Should find the category")
	assert.Equal(t, "Fixed Test Category", cat.Name, "Category name should match")
}
