package db_provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/config"
)

// MockConfigProvider is a mock implementation of the ConfigProvider interface
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

type testGormOpener struct{}

func (o *testGormOpener) Open(dsn string, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), config)
}

func TestNewDefaultDBProvider(t *testing.T) {
	// Test the constructor function
	provider := NewDefaultDBProvider()
	assert.NotNil(t, provider)
	assert.NotNil(t, provider.configProvider)
	assert.IsType(t, &config.DefaultConfigProvider{}, provider.configProvider)
}

func TestConstructDSN(t *testing.T) {
	// Test that the DSN string is correctly constructed
	conf := &config.Conf{
		DB: &config.DbConfig{
			Host:     "testhost",
			Port:     5432,
			SslMode:  "disable",
			User:     "testuser",
			Password: "testpass",
			DbName:   "testdb",
		},
	}
	provider := &DefaultDBProvider{configProvider: nil}

	expectedDSN := "host=testhost port=5432 sslmode=disable user=testuser password=testpass dbname=testdb"
	actualDSN := provider.constructDSN(conf)

	assert.Equal(t, expectedDSN, actualDSN)
}

func TestGetDB_ConfigConstruction(t *testing.T) {
	// Create mock config provider
	mockConfig := new(MockConfigProvider)
	dbConfig := &config.DbConfig{
		Host:     "testhost",
		Port:     5432,
		SslMode:  "disable",
		User:     "testuser",
		Password: "testpassword",
		DbName:   "testdb",
	}
	mockConfig.On("GetConf").Return(&config.Conf{DB: dbConfig}, nil)

	// Create our provider with the mock config
	provider := &DefaultDBProvider{
		configProvider: mockConfig,
		gormOpener:     &testGormOpener{},
	}

	_, err := provider.GetDB(context.Background())
	require.NoError(t, err)

	// Verify the mock was called
	mockConfig.AssertExpectations(t)
}

func TestGetDB_Success(t *testing.T) {
	// Create mock config provider
	mockConfig := new(MockConfigProvider)
	dbConfig := &config.DbConfig{
		Host:     "testhost",
		Port:     5432,
		SslMode:  "disable",
		User:     "testuser",
		Password: "testpassword",
		DbName:   "testdb",
	}
	mockConfig.On("GetConf").Return(&config.Conf{DB: dbConfig}, nil)

	// Create provider with mock config
	provider := &DefaultDBProvider{
		configProvider: mockConfig,
		gormOpener:     &testGormOpener{},
	}

	_, err := provider.GetDB(context.Background())
	require.NoError(t, err)

	// Verify our config provider was called as expected
	mockConfig.AssertExpectations(t)
}
