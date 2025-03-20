package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockLoad is a test implementation of the config loader function
func MockLoad(fileName string) (*Conf, error) {
	// Check the file name and return appropriate test data
	if fileName == "etc/config.yml" {
		// Return test data for default config file
		return &Conf{
			DB: &DbConfig{
				Host:     "localhost",
				Port:     5432,
				SslMode:  "disable",
				User:     "testuser",
				Password: "testpass",
				DbName:   "testdb",
			},
		}, nil
	} else if fileName == "etc/invalid.yml" {
		// Return error for invalid file
		return nil, errors.New("failed to read config file")
	} else if fileName == "etc/custom.yml" {
		// Return test data for custom file
		return &Conf{
			DB: &DbConfig{
				Host:     "customhost",
				Port:     5433,
				SslMode:  "prefer",
				User:     "customuser",
				Password: "custompass",
				DbName:   "customdb",
			},
		}, nil
	}

	// For other file names - return error
	return nil, errors.New("unknown config file")
}

func TestNewDefaultConfigProvider(t *testing.T) {
	// Test that the constructor creates an object with the correct default filename
	provider := NewDefaultConfigProvider()

	// Verify provider is not nil
	assert.NotNil(t, provider)

	// Verify filename is set correctly
	assert.Equal(t, "etc/config.yml", provider.fname)

	// Verify load function is set
	assert.NotNil(t, provider.loadFunc)
}

func TestNewConfigProviderWithCustomLoader(t *testing.T) {
	// Test that the constructor with custom loader works
	customFilename := "custom/path.yml"
	provider := NewDefaultConfigProviderWithCustomLoader(customFilename, MockLoad)

	// Verify provider is not nil
	assert.NotNil(t, provider)

	// Verify filename is set correctly
	assert.Equal(t, customFilename, provider.fname)

	// Verify load function is set to our mock
	assert.NotNil(t, provider.loadFunc)
}

func TestDefaultConfigProvider_GetConf(t *testing.T) {
	// Create provider with mock loader
	provider := NewDefaultConfigProviderWithCustomLoader("etc/config.yml", MockLoad)

	// Call GetConf and check the result
	conf, err := provider.GetConf()

	// Verify no error occurred
	require.NoError(t, err)
	// Verify config is not nil
	require.NotNil(t, conf)
	// Verify DB is not nil
	require.NotNil(t, conf.DB)

	// Verify config values
	assert.Equal(t, "localhost", conf.DB.Host)
	assert.Equal(t, 5432, conf.DB.Port)
	assert.Equal(t, "disable", conf.DB.SslMode)
	assert.Equal(t, "testuser", conf.DB.User)
	assert.Equal(t, "testpass", conf.DB.Password)
	assert.Equal(t, "testdb", conf.DB.DbName)
}

func TestDefaultConfigProvider_GetConf_WithCustomFile(t *testing.T) {
	// Create provider with custom filename and mock loader
	provider := NewDefaultConfigProviderWithCustomLoader("etc/custom.yml", MockLoad)

	// Call GetConf and check the result
	conf, err := provider.GetConf()

	// Verify no error occurred
	require.NoError(t, err)
	// Verify config is not nil
	require.NotNil(t, conf)
	// Verify DB is not nil
	require.NotNil(t, conf.DB)

	// Verify custom config values
	assert.Equal(t, "customhost", conf.DB.Host)
	assert.Equal(t, 5433, conf.DB.Port)
	assert.Equal(t, "prefer", conf.DB.SslMode)
	assert.Equal(t, "customuser", conf.DB.User)
	assert.Equal(t, "custompass", conf.DB.Password)
	assert.Equal(t, "customdb", conf.DB.DbName)
}

func TestDefaultConfigProvider_GetConf_Error(t *testing.T) {
	// Create provider with invalid filename and mock loader
	provider := NewDefaultConfigProviderWithCustomLoader("etc/invalid.yml", MockLoad)

	// Call GetConf and check the result
	conf, err := provider.GetConf()

	// Verify error occurred
	require.Error(t, err)
	// Verify error message
	assert.Contains(t, err.Error(), "failed to read config file")
	// Verify config is nil
	assert.Nil(t, conf)
}

func TestDefaultConfigProvider_GetConf_UnknownFile(t *testing.T) {
	// Create provider with unknown filename and mock loader
	provider := NewDefaultConfigProviderWithCustomLoader("etc/unknown.yml", MockLoad)

	// Call GetConf and check the result
	conf, err := provider.GetConf()

	// Verify error occurred
	require.Error(t, err)
	// Verify error message
	assert.Contains(t, err.Error(), "unknown config file")
	// Verify config is nil
	assert.Nil(t, conf)
}
