package services

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meesooqa/cheque/db/models"
)

// Integration test that tests the full flow from reading to saving
func TestCategoriesFullImportFlow(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "test-taxonomy.txt")

	testData := `# Google_Product_Taxonomy_Version: 2021-09-21
5181 - Багаж и сумки
110 - Багаж и сумки > Багажные принадлежности
5652 - Багаж и сумки > Багажные принадлежности > Багажные ремни
`

	err := os.WriteFile(testFilePath, []byte(testData), 0644)
	require.NoError(t, err, "Failed to create test file")

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	// Setup real SQLite DB
	db := setupTestDB(t)
	dbProvider := &TestDBProvider{DB: db}

	// Create the components
	reader := NewGoogleProductTaxonomyReader(logger, testFilePath)
	importer := NewGoogleProductTaxonomyImporter(logger, dbProvider)

	// Execute the test
	items, err := reader.Read()
	require.NoError(t, err, "Reading should not fail")
	assert.Len(t, items, 3, "Should have 3 items")

	err = importer.Save(context.Background(), items)
	require.NoError(t, err, "Saving should not fail")

	// Verify the database state
	var categories []models.Category
	result := db.Find(&categories)
	require.NoError(t, result.Error, "Should find categories")
	assert.Len(t, categories, 3, "Should have 3 categories in DB")

	// Verify hierarchy
	var root models.Category
	db.First(&root, "name = ?", "Багаж и сумки")
	assert.Nil(t, root.ParentID, "Root should have nil parent")

	var level1 models.Category
	db.First(&level1, "name = ?", "Багажные принадлежности")
	require.NotNil(t, level1.ParentID, "Level 1 should have parent")
	assert.Equal(t, root.ID, *level1.ParentID)

	var level2 models.Category
	db.First(&level2, "name = ?", "Багажные ремни")
	require.NotNil(t, level2.ParentID, "Level 2 should have parent")
	assert.Equal(t, level1.ID, *level2.ParentID)
}
