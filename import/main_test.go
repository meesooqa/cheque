package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/config"
)

type MockDBProvider struct {
	mock.Mock
}

func (m *MockDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gorm.DB), args.Error(1)
}

type MockConfigProvider struct {
	mock.Mock
}

func (m *MockConfigProvider) GetConf() (*config.Conf, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*config.Conf), args.Error(1)
}

type MockImportService struct {
	mock.Mock
}

func (m *MockImportService) SaveReceipt(db *gorm.DB, data []byte) error {
	args := m.Called(db, data)
	return args.Error(0)
}

func TestRunImport(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "import-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	testDataPath := filepath.Join(tempDir, "data")
	err = os.Mkdir(testDataPath, 0755)
	require.NoError(t, err)

	testDataFile := filepath.Join(testDataPath, "extract.json")
	testData := []byte(`[{"some":"test data"}]`)
	err = os.WriteFile(testDataFile, testData, 0644)
	require.NoError(t, err)

	t.Run("successful case", func(t *testing.T) {
		mockDB := &gorm.DB{}
		mockDBProvider := new(MockDBProvider)
		mockConfigProvider := new(MockConfigProvider)
		mockImportService := new(MockImportService)

		mockDBProvider.On("GetDB", mock.Anything).Return(mockDB, nil)
		mockConfigProvider.On("GetConf").Return(&config.Conf{
			System: &config.SystemConfig{
				DataPath: testDataPath,
			},
		}, nil)
		mockImportService.On("SaveReceipt", mockDB, testData).Return(nil)

		err := runImport(
			mockDBProvider,
			mockConfigProvider,
			mockImportService,
			os.ReadFile,
		)

		assert.NoError(t, err)
		mockDBProvider.AssertExpectations(t)
		mockConfigProvider.AssertExpectations(t)
		mockImportService.AssertExpectations(t)
	})

	t.Run("error getting db", func(t *testing.T) {
		mockDBProvider := new(MockDBProvider)
		mockDBProvider.On("GetDB", mock.Anything).
			Return(nil, assert.AnError)

		err := runImport(
			mockDBProvider,
			nil,
			nil,
			nil,
		)

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockDBProvider.AssertExpectations(t)
	})

	t.Run("error loading configuration", func(t *testing.T) {
		mockDB := &gorm.DB{}
		mockDBProvider := new(MockDBProvider)
		mockConfigProvider := new(MockConfigProvider)

		mockDBProvider.On("GetDB", mock.Anything).Return(mockDB, nil)
		mockConfigProvider.On("GetConf").Return(nil, assert.AnError)

		err := runImport(
			mockDBProvider,
			mockConfigProvider,
			nil,
			nil,
		)

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockDBProvider.AssertExpectations(t)
		mockConfigProvider.AssertExpectations(t)
	})

	t.Run("file read error", func(t *testing.T) {
		mockDB := &gorm.DB{}
		mockDBProvider := new(MockDBProvider)
		mockConfigProvider := new(MockConfigProvider)

		mockDBProvider.On("GetDB", mock.Anything).Return(mockDB, nil)
		mockConfigProvider.On("GetConf").Return(&config.Conf{
			System: &config.SystemConfig{
				DataPath: testDataPath,
			},
		}, nil)

		mockReadFile := func(filename string) ([]byte, error) {
			return nil, assert.AnError
		}

		err := runImport(
			mockDBProvider,
			mockConfigProvider,
			nil,
			mockReadFile,
		)

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockDBProvider.AssertExpectations(t)
		mockConfigProvider.AssertExpectations(t)
	})

	t.Run("import error", func(t *testing.T) {
		mockDB := &gorm.DB{}
		mockDBProvider := new(MockDBProvider)
		mockConfigProvider := new(MockConfigProvider)
		mockImportService := new(MockImportService)

		mockDBProvider.On("GetDB", mock.Anything).Return(mockDB, nil)
		mockConfigProvider.On("GetConf").Return(&config.Conf{
			System: &config.SystemConfig{
				DataPath: testDataPath,
			},
		}, nil)
		mockImportService.On("SaveReceipt", mockDB, testData).Return(assert.AnError)

		err := runImport(
			mockDBProvider,
			mockConfigProvider,
			mockImportService,
			os.ReadFile,
		)

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockDBProvider.AssertExpectations(t)
		mockConfigProvider.AssertExpectations(t)
		mockImportService.AssertExpectations(t)
	})
}
