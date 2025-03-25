package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/models"
)

// setupTestDB sets up a test database using SQLite in-memory
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err, "Failed to connect to test database")

	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	require.NoError(t, err, "Failed to migrate test database")

	return db
}

// createTempCategoriesFile creates a temporary file with test categories
func createTempCategoriesFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "categories-*.txt")
	require.NoError(t, err, "Failed to create temp file")

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err, "Failed to write to temp file")

	err = tmpFile.Close()
	require.NoError(t, err, "Failed to close temp file")

	return tmpFile.Name()
}

func TestImportFromFile(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)

	// Create a temp file with test data
	testContent := `# Test categories
100 - Electronics
101 - Electronics > Smartphones
102 - Electronics > Smartphones > Android
200 - Books
201 - Books > Fiction
202 - Books > Fiction > Fantasy
`
	tempFileName := createTempCategoriesFile(t, testContent)
	defer os.Remove(tempFileName)

	// Create importer and import from file
	importer := NewCategoryImporter(db)
	err := importer.ImportFromFile(tempFileName)
	assert.NoError(t, err, "Import should not fail")

	// Verify imported categories
	var count int64
	db.Model(&models.Category{}).Count(&count)
	assert.Equal(t, int64(6), count, "Should have imported 6 categories")

	// Check each category
	checkCategory(t, db, uint(100), "Electronics", nil)

	// Find parent ID for Smartphones
	var electronics models.Category
	db.First(&electronics, 100)
	electronicsID := electronics.ID

	checkCategory(t, db, uint(101), "Smartphones", &electronicsID)

	// Find parent ID for Android
	var smartphones models.Category
	db.First(&smartphones, 101)
	smartphonesID := smartphones.ID

	checkCategory(t, db, uint(102), "Android", &smartphonesID)

	// Check Books category
	checkCategory(t, db, uint(200), "Books", nil)

	// Check Fiction category
	var books models.Category
	db.First(&books, 200)
	booksID := books.ID

	checkCategory(t, db, uint(201), "Fiction", &booksID)

	// Check Fantasy category
	var fiction models.Category
	db.First(&fiction, 201)
	fictionID := fiction.ID

	checkCategory(t, db, uint(202), "Fantasy", &fictionID)
}

func TestNestedCategoriesWithSharedPath(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)

	// Create a temp file where IDs for intermediate categories can be found in other lines
	testContent := `100 - Electronics
200 - Home & Garden
300 - Electronics > Smartphones
400 - Home & Garden > Furniture
500 - Electronics > Smartphones > Android
600 - Home & Garden > Furniture > Sofas
`
	tempFileName := createTempCategoriesFile(t, testContent)
	defer os.Remove(tempFileName)

	// Create importer and import from file
	importer := NewCategoryImporter(db)
	err := importer.ImportFromFile(tempFileName)
	assert.NoError(t, err, "Import should not fail")

	// Check ID usage for intermediate categories
	checkCategory(t, db, uint(100), "Electronics", nil)
	checkCategory(t, db, uint(300), "Smartphones", &[]uint{100}[0])
	checkCategory(t, db, uint(500), "Android", &[]uint{300}[0])

	checkCategory(t, db, uint(200), "Home & Garden", nil)
	checkCategory(t, db, uint(400), "Furniture", &[]uint{200}[0])
	checkCategory(t, db, uint(600), "Sofas", &[]uint{400}[0])
}

func TestImportInvalidFile(t *testing.T) {
	db := setupTestDB(t)
	importer := NewCategoryImporter(db)

	err := importer.ImportFromFile("non_existent_file.txt")
	assert.Error(t, err, "Should fail with non-existent file")
}

func TestProcessInvalidLine(t *testing.T) {
	db := setupTestDB(t)
	importer := NewCategoryImporter(db)

	// Create a temp file with invalid data
	testContent := `This is an invalid line
100-Electronics
NotANumber - Category
`
	tempFileName := createTempCategoriesFile(t, testContent)
	defer os.Remove(tempFileName)

	err := importer.ImportFromFile(tempFileName)
	assert.Error(t, err, "Should fail with invalid data")
}

func TestImportDuplicateIDs(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)

	// Create a temp file with duplicate IDs
	testContent := `100 - Electronics
100 - Books  # Same ID as Electronics
`
	tempFileName := createTempCategoriesFile(t, testContent)
	defer os.Remove(tempFileName)

	// Create importer and import from file
	importer := NewCategoryImporter(db)
	err := importer.ImportFromFile(tempFileName)
	assert.NoError(t, err, "Import should not fail")

	// Verify that the second category with the same ID updated the first one
	var category models.Category
	result := db.First(&category, 100)
	assert.NoError(t, result.Error, "Category should exist")
	assert.Equal(t, "Books", category.Name, "Category should be updated to Books")

	// Verify total count is 1
	var count int64
	db.Model(&models.Category{}).Count(&count)
	assert.Equal(t, int64(1), count, "Should have 1 category after update")
}

func TestLineWithComments(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)

	// Create a temp file with comments in the middle of lines
	testContent := `100 - Electronics # Comment here
200 - Books # Another comment
300 - Electronics > Smartphones # Yet another comment
`
	tempFileName := createTempCategoriesFile(t, testContent)
	defer os.Remove(tempFileName)

	// Create importer and import from file
	importer := NewCategoryImporter(db)
	err := importer.ImportFromFile(tempFileName)
	assert.NoError(t, err, "Import should not fail")

	// Verify categories were imported correctly without comments
	checkCategory(t, db, uint(100), "Electronics", nil)
	checkCategory(t, db, uint(200), "Books", nil)
	checkCategory(t, db, uint(300), "Smartphones", &[]uint{100}[0])
}

// Helper function to check a category
func checkCategory(t *testing.T, db *gorm.DB, id uint, name string, parentID *uint) {
	var category models.Category
	result := db.First(&category, id)
	assert.NoError(t, result.Error, "Category should exist")
	assert.Equal(t, name, category.Name, "Category name should match")

	if parentID == nil {
		assert.Nil(t, category.ParentID, "Parent ID should be nil")
	} else {
		assert.Equal(t, *parentID, *category.ParentID, "Parent ID should match")
	}
}
