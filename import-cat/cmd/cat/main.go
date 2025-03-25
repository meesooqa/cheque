package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/db/db_provider"
	"github.com/meesooqa/cheque/db/models"
)

// CategoryImporter handles the process of importing categories
type CategoryImporter struct {
	db *gorm.DB
}

// CategoryPathEntry represents a parsed line from the file
type CategoryPathEntry struct {
	ID   uint
	Path string
}

// NewCategoryImporter creates a new instance of CategoryImporter
func NewCategoryImporter(db *gorm.DB) *CategoryImporter {
	return &CategoryImporter{db: db}
}

// cleanLine removes comments from a line
func cleanLine(line string) string {
	// Check if there's a comment in the middle of the line
	if commentPos := strings.Index(line, "#"); commentPos >= 0 {
		line = line[:commentPos]
	}
	return strings.TrimSpace(line)
}

// ImportFromFile imports categories from a file
func (ci *CategoryImporter) ImportFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse all category path entries, preserving order
	var categoryEntries []CategoryPathEntry
	categoryPaths := make(map[string]uint) // maps full path to ID (for lookup)

	// Process each line to extract category paths and their IDs
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Clean line from comments and whitespace
		line = cleanLine(line)

		// Parse the line
		parts := strings.SplitN(line, " - ", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		// Parse ID
		id, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 32)
		if err != nil {
			return fmt.Errorf("invalid ID: %s", parts[0])
		}

		path := strings.TrimSpace(parts[1])
		// Add to ordered slice and map
		entry := CategoryPathEntry{
			ID:   uint(id),
			Path: path,
		}
		categoryEntries = append(categoryEntries, entry)
		categoryPaths[path] = uint(id)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	// Clear the database to ensure clean import
	if err := ci.db.Exec("DELETE FROM categories").Error; err != nil {
		return fmt.Errorf("failed to clear categories table: %w", err)
	}

	// Process categories in the order they appeared in the file
	processedPaths := make(map[string]uint)    // Maps category path to ID
	processedCategories := make(map[uint]bool) // Tracks which IDs have been processed

	// First process intermediate paths
	for _, entry := range categoryEntries {
		parts := strings.Split(entry.Path, " > ")
		if len(parts) > 1 {
			// Process intermediate paths first
			var currentPath string
			var parentID *uint

			for i, part := range parts[:len(parts)-1] {
				part = strings.TrimSpace(part)

				// Build the current path
				if i == 0 {
					currentPath = part
				} else {
					currentPath += " > " + part
				}

				// Check if this path already has an assigned ID
				if id, exists := processedPaths[currentPath]; exists {
					tempID := id // Create a temporary variable to get the address
					parentID = &tempID
					continue
				}

				// Find ID for this intermediate path if it exists
				var categoryID uint
				if id, exists := categoryPaths[currentPath]; exists {
					categoryID = id
				} else {
					// Create with auto ID
					category := models.Category{
						Name:     part,
						ParentID: parentID,
					}
					if err := ci.db.Create(&category).Error; err != nil {
						return fmt.Errorf("failed to create intermediate category: %w", err)
					}
					processedPaths[currentPath] = category.ID
					processedCategories[category.ID] = true
					tempID := category.ID
					parentID = &tempID
					continue
				}

				// Check if this ID is already processed
				if _, processed := processedCategories[categoryID]; processed {
					tempID := categoryID
					parentID = &tempID
					continue
				}

				// Create or update with ID from the path map
				var category models.Category
				result := ci.db.First(&category, categoryID)

				if result.Error == nil {
					// Update
					category.Name = part
					category.ParentID = parentID
					if err := ci.db.Save(&category).Error; err != nil {
						return fmt.Errorf("failed to update category: %w", err)
					}
				} else if result.Error == gorm.ErrRecordNotFound {
					// Create
					category = models.Category{
						Name:     part,
						ParentID: parentID,
					}
					category.ID = categoryID
					if err := ci.db.Create(&category).Error; err != nil {
						return fmt.Errorf("failed to create category: %w", err)
					}
				} else {
					return fmt.Errorf("failed to check if category exists: %w", result.Error)
				}

				processedPaths[currentPath] = categoryID
				processedCategories[categoryID] = true
				tempID := categoryID
				parentID = &tempID
			}
		}
	}

	// Then process the final category of each path in the original order
	for _, entry := range categoryEntries {
		parts := strings.Split(entry.Path, " > ")
		var parentPath string

		// If there are parent categories, find the parent ID
		var parentID *uint
		if len(parts) > 1 {
			parentPath = strings.Join(parts[:len(parts)-1], " > ")
			if id, exists := processedPaths[parentPath]; exists {
				tempID := id
				parentID = &tempID
			}
		}

		// Get the final category name
		finalName := strings.TrimSpace(parts[len(parts)-1])

		// Create or update the final category with the ID from the file
		var category models.Category
		result := ci.db.First(&category, entry.ID)

		if result.Error == nil {
			// Update existing category
			category.Name = finalName
			category.ParentID = parentID
			if err := ci.db.Save(&category).Error; err != nil {
				return fmt.Errorf("failed to update final category: %w", err)
			}
		} else if result.Error == gorm.ErrRecordNotFound {
			// Create new category
			category = models.Category{
				Name:     finalName,
				ParentID: parentID,
			}
			category.ID = entry.ID
			if err := ci.db.Create(&category).Error; err != nil {
				return fmt.Errorf("failed to create final category: %w", err)
			}
		} else {
			return fmt.Errorf("failed to check if final category exists: %w", result.Error)
		}

		// Record the full path
		processedPaths[entry.Path] = entry.ID
		processedCategories[entry.ID] = true
	}

	return nil
}

func main() {
	dbp := db_provider.NewDefaultDBProvider()
	db, err := dbp.GetDB(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create a category importer
	importer := NewCategoryImporter(db)

	// Import categories from file
	configProvider := config.NewDefaultConfigProvider()
	conf, err := configProvider.GetConf()
	if err != nil {
		log.Fatal(err)
	}
	filePath := conf.System.DataPath + "/cat/taxonomy-with-ids.ru-RU.txt"
	err = importer.ImportFromFile(filePath)
	if err != nil {
		log.Fatalf("failed to import categories: %v", err)
	}

	log.Println("categories imported successfully!")
}
