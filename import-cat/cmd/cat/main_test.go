package main

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/meesooqa/cheque/import-cat/services"
)

// MockCategoriesReader mocks the CategoriesReader interface
type MockCategoriesReader struct {
	mock.Mock
}

func (m *MockCategoriesReader) Read() ([]services.CategoriesReaderResultItem, error) {
	args := m.Called()
	return args.Get(0).([]services.CategoriesReaderResultItem), args.Error(1)
}

// MockCategoriesImporter mocks the CategoriesImporter interface
type MockCategoriesImporter struct {
	mock.Mock
}

func (m *MockCategoriesImporter) Save(ctx context.Context, items []services.CategoriesReaderResultItem) error {
	args := m.Called(ctx, items)
	return args.Error(0)
}

func TestRunCategoriesImport_Success(t *testing.T) {
	// Setup mocks
	reader := new(MockCategoriesReader)
	importer := new(MockCategoriesImporter)

	// Mock data
	mockItems := []services.CategoriesReaderResultItem{
		&services.GoogleProductTaxonomyItem{
			ID:       5181,
			Name:     "Багаж и сумки",
			FullPath: "Багаж и сумки",
			Level:    0,
		},
		&services.GoogleProductTaxonomyItem{
			ID:       110,
			Name:     "Багажные принадлежности",
			FullPath: "Багаж и сумки > Багажные принадлежности",
			Level:    1,
		},
	}

	// Set expectations
	reader.On("Read").Return(mockItems, nil)
	importer.On("Save", mock.Anything, mockItems).Return(nil)

	// Execute test
	err := runCategoriesImport(reader, importer)

	// Assert results
	require.NoError(t, err)
	reader.AssertExpectations(t)
	importer.AssertExpectations(t)
}

func TestRunCategoriesImport_ReaderError(t *testing.T) {
	// Setup mocks
	reader := new(MockCategoriesReader)
	importer := new(MockCategoriesImporter)

	// Mock reader error
	expectedErr := errors.New("file not found")
	reader.On("Read").Return([]services.CategoriesReaderResultItem{}, expectedErr)

	// Execute test
	err := runCategoriesImport(reader, importer)

	// Assert results
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	reader.AssertExpectations(t)
	importer.AssertNotCalled(t, "Save")
}

func TestRunCategoriesImport_ImporterError(t *testing.T) {
	// Setup mocks
	reader := new(MockCategoriesReader)
	importer := new(MockCategoriesImporter)

	// Mock data
	mockItems := []services.CategoriesReaderResultItem{
		&services.GoogleProductTaxonomyItem{
			ID:       5181,
			Name:     "Багаж и сумки",
			FullPath: "Багаж и сумки",
			Level:    0,
		},
	}

	// Mock responses
	expectedErr := errors.New("database error")
	reader.On("Read").Return(mockItems, nil)
	importer.On("Save", mock.Anything, mockItems).Return(expectedErr)

	// Execute test
	err := runCategoriesImport(reader, importer)

	// Assert results
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	reader.AssertExpectations(t)
	importer.AssertExpectations(t)
}

func TestRunCategoriesImport_EmptyItems(t *testing.T) {
	// Setup mocks
	reader := new(MockCategoriesReader)
	importer := new(MockCategoriesImporter)

	// Mock empty items list
	var emptyItems []services.CategoriesReaderResultItem
	reader.On("Read").Return(emptyItems, nil)

	// Execute test
	err := runCategoriesImport(reader, importer)

	// Assert results
	assert.NoError(t, err)
	reader.AssertExpectations(t)
	importer.AssertNotCalled(t, "Save", "Save should not be called with empty items")
}

// Для более полного покрытия можно также создать интеграционный тест для функции main,
// но для этого потребуется рефакторинг main, чтобы сделать его более тестируемым.
// Вот пример, как можно это сделать:

// MockMainDependencies содержит все моки для зависимостей main
type MockMainDependencies struct {
	Reader   *MockCategoriesReader
	Importer *MockCategoriesImporter
}

// TestableMain - версия main, которую можно тестировать
func TestableMain(deps MockMainDependencies) error {
	return runCategoriesImport(deps.Reader, deps.Importer)
}

func TestMainIntegration(t *testing.T) {
	// Setup mocks
	reader := new(MockCategoriesReader)
	importer := new(MockCategoriesImporter)

	// Mock data
	mockItems := []services.CategoriesReaderResultItem{
		&services.GoogleProductTaxonomyItem{
			ID:       5181,
			Name:     "Багаж и сумки",
			FullPath: "Багаж и сумки",
			Level:    0,
		},
	}

	// Set expectations
	reader.On("Read").Return(mockItems, nil)
	importer.On("Save", mock.Anything, mockItems).Return(nil)

	// Create dependencies
	deps := MockMainDependencies{
		Reader:   reader,
		Importer: importer,
	}

	// Execute testable version of main
	err := TestableMain(deps)

	// Assert results
	assert.NoError(t, err)
	reader.AssertExpectations(t)
	importer.AssertExpectations(t)
}
