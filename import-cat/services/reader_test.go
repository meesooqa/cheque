package services

import (
	"os"
	"path/filepath"
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoogleProductTaxonomyReader_Read(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "test-taxonomy.txt")

	testData := `# Google_Product_Taxonomy_Version: 2021-09-21
5181 - Багаж и сумки
110 - Багаж и сумки > Багажные принадлежности
5652 - Багаж и сумки > Багажные принадлежности > Багажные ремни
# Comment line
invalid line format
`

	err := os.WriteFile(testFilePath, []byte(testData), 0644)
	require.NoError(t, err, "Failed to create test file")

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	// Create the reader with test file
	reader := NewGoogleProductTaxonomyReader(logger, testFilePath)

	// Execute the test
	items, err := reader.Read()

	// Verify results
	require.NoError(t, err, "Read should not return an error")
	assert.Len(t, items, 3, "Should have parsed 3 valid entries")

	// Verify the parsed items
	item1, ok := items[0].(*GoogleProductTaxonomyItem)
	require.True(t, ok, "Item should be of type GoogleProductTaxonomyItem")
	assert.Equal(t, uint(5181), item1.ID)
	assert.Equal(t, "Багаж и сумки", item1.Name)
	assert.Equal(t, "Багаж и сумки", item1.FullPath)
	assert.Equal(t, 0, item1.Level)

	item2, ok := items[1].(*GoogleProductTaxonomyItem)
	require.True(t, ok, "Item should be of type GoogleProductTaxonomyItem")
	assert.Equal(t, uint(110), item2.ID)
	assert.Equal(t, "Багажные принадлежности", item2.Name)
	assert.Equal(t, "Багаж и сумки > Багажные принадлежности", item2.FullPath)
	assert.Equal(t, 1, item2.Level)

	item3, ok := items[2].(*GoogleProductTaxonomyItem)
	require.True(t, ok, "Item should be of type GoogleProductTaxonomyItem")
	assert.Equal(t, uint(5652), item3.ID)
	assert.Equal(t, "Багажные ремни", item3.Name)
	assert.Equal(t, "Багаж и сумки > Багажные принадлежности > Багажные ремни", item3.FullPath)
	assert.Equal(t, 2, item3.Level)
}

func TestGoogleProductTaxonomyReader_FileNotFound(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	// Create reader with non-existent file
	reader := NewGoogleProductTaxonomyReader(logger, "non-existent-file.txt")

	// Execute the test
	items, err := reader.Read()

	// Verify results
	assert.Error(t, err, "Read should return an error for non-existent file")
	assert.Nil(t, items, "Items should be nil when there's an error")
}
