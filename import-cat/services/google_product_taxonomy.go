package services

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

// GoogleProductTaxonomyItem represents a category from the import file
type GoogleProductTaxonomyItem struct {
	ID       uint
	Name     string
	FullPath string
	Level    int
}

type GoogleProductTaxonomyReader struct {
	logger   *slog.Logger
	filePath string
	adapter  CategoriesAdapter
}

type GoogleProductTaxonomyAdapter struct {
	logger *slog.Logger
}

type GoogleProductTaxonomyImporter struct {
	logger     *slog.Logger
	dbProvider db_types.DBProvider
}

func NewGoogleProductTaxonomyReader(logger *slog.Logger, filePath string) *GoogleProductTaxonomyReader {
	return &GoogleProductTaxonomyReader{
		logger:   logger,
		filePath: filePath,
		adapter:  NewGoogleProductTaxonomyAdapter(logger),
	}
}

func NewGoogleProductTaxonomyAdapter(logger *slog.Logger) *GoogleProductTaxonomyAdapter {
	return &GoogleProductTaxonomyAdapter{logger: logger}
}

func NewGoogleProductTaxonomyImporter(logger *slog.Logger, dbProvider db_types.DBProvider) *GoogleProductTaxonomyImporter {
	return &GoogleProductTaxonomyImporter{
		logger:     logger,
		dbProvider: dbProvider,
	}
}

// Read reads and parses the categories file
func (o *GoogleProductTaxonomyReader) Read() ([]CategoriesReaderResultItem, error) {
	file, err := os.Open(o.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []CategoriesReaderResultItem
	for scanner.Scan() {
		line := scanner.Text()
		item, err := o.adapter.Convert(line)
		if err != nil {
			o.logger.Error("failed to parse line", slog.String("line", line), slog.Any("error", err))
		}
		if item != nil {
			items = append(items, item)
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (o *GoogleProductTaxonomyAdapter) Convert(line string) (CategoriesReaderResultItem, error) {
	// Skip empty lines and comments
	if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
		return nil, nil
	}

	// Parse the category line
	// Format: ID - Path > Subpath > Name
	parts := strings.SplitN(line, "-", 2)
	if len(parts) != 2 {
		o.logger.Warn("line doesn't match expected format", slog.String("line", line))
		return nil, nil
	}

	// Parse ID
	idStr := strings.TrimSpace(parts[0])
	var id uint
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		o.logger.Warn("couldn't parse ID from line", slog.String("line", line), slog.Any("err", err))
		return nil, err
	}

	// Parse full path
	fullPath := strings.TrimSpace(parts[1])

	// Calculate level and get name
	pathParts := strings.Split(fullPath, ">")
	for i := range pathParts {
		pathParts[i] = strings.TrimSpace(pathParts[i])
	}

	level := len(pathParts) - 1
	name := pathParts[level]

	return &GoogleProductTaxonomyItem{
		ID:       id,
		Name:     name,
		FullPath: fullPath,
		Level:    level,
	}, nil
}

// Save imports categories to the database
func (o *GoogleProductTaxonomyImporter) Save(ctx context.Context, items []CategoriesReaderResultItem) error {
	db, err := o.dbProvider.GetDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Create a map to store category ID by full path
	categoryMap := make(map[string]uint)

	// Process categories by levels to ensure parents are created before children
	for level := 0; level <= o.getMaxLevel(items); level++ {
		for _, resultItem := range items {
			item, ok := resultItem.(*GoogleProductTaxonomyItem)
			if !ok {
				continue
			}
			if item.Level != level {
				continue
			}

			var parentID *uint

			// Find parent ID if this is not a root category
			if item.Level > 0 {
				// Get parent path by removing the last part
				pathParts := strings.Split(item.FullPath, ">")
				parentPath := strings.Join(pathParts[:len(pathParts)-1], ">")
				parentPath = strings.TrimSpace(parentPath)

				// Look up parent ID in our map
				if parentCatID, exists := categoryMap[parentPath]; exists {
					parentID = &parentCatID
				} else {
					log.Printf("Warning: parent not found for category: %s", item.FullPath)
				}
			}

			// Check if category already exists in DB
			var existingCategory models.Category
			result := db.Where("name = ?", item.Name).First(&existingCategory)

			if result.Error == nil {
				// Update the existing category
				existingCategory.ParentID = parentID
				if err := db.Save(&existingCategory).Error; err != nil {
					return fmt.Errorf("failed to update category %s: %v", item.Name, err)
				}
				categoryMap[item.FullPath] = existingCategory.ID
			} else if result.Error == gorm.ErrRecordNotFound {
				// Create a new category
				newCategory := &models.Category{
					Name:     item.Name,
					ParentID: parentID,
				}

				if err := db.Create(&newCategory).Error; err != nil {
					return fmt.Errorf("failed to create category %s: %v", item.Name, err)
				}
				categoryMap[item.FullPath] = newCategory.ID
			} else {
				return fmt.Errorf("database error when checking for category %s: %v", item.Name, result.Error)
			}
		}
	}

	return nil
}

// getMaxLevel finds the maximum nesting level in the categories
func (o *GoogleProductTaxonomyImporter) getMaxLevel(items []CategoriesReaderResultItem) int {
	maxLevel := 0
	for _, resultItem := range items {
		item, ok := resultItem.(*GoogleProductTaxonomyItem)
		if !ok {
			continue
		}
		if item.Level > maxLevel {
			maxLevel = item.Level
		}
	}
	return maxLevel
}
