package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meesooqa/cheque/common/config"
)

func TestNewGormProtoDataProvider(t *testing.T) {
	conf := &config.GormGenProtoConfig{PathMaps: "/test/maps"}
	provider := NewGormProtoDataProvider(conf)
	assert.NotNil(t, provider)
	assert.Equal(t, conf, provider.conf)
}

func TestGormProtoDataProvider_GetGormProtoMap(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(t *testing.T) *config.GormGenProtoConfig
		expected    map[string]string
		expectError bool
	}{
		{
			name: "successful type map loading",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				tmpDir := t.TempDir()
				typeFile := filepath.Join(tmpDir, "type.json")
				require.NoError(t, os.WriteFile(typeFile, []byte(`{"uint":"uint64"}`), 0644))
				return &config.GormGenProtoConfig{PathMaps: tmpDir}
			},
			expected:    map[string]string{"uint": "uint64"},
			expectError: false,
		},
		{
			name: "missing type file",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				return &config.GormGenProtoConfig{PathMaps: t.TempDir()}
			},
			expectError: true,
		},
		{
			name: "invalid type file content",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				tmpDir := t.TempDir()
				typeFile := filepath.Join(tmpDir, "type.json")
				require.NoError(t, os.WriteFile(typeFile, []byte(`invalid`), 0644))
				return &config.GormGenProtoConfig{PathMaps: tmpDir}
			},
			expectError: true,
		},
		{
			name: "path with trailing slash",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				tmpDir := t.TempDir() + string(filepath.Separator)
				typeFile := filepath.Join(tmpDir, "type.json")
				require.NoError(t, os.WriteFile(typeFile, []byte(`{"string":"string"}`), 0644))
				return &config.GormGenProtoConfig{PathMaps: tmpDir}
			},
			expected:    map[string]string{"string": "string"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := tt.setup(t)
			provider := NewGormProtoDataProvider(conf)

			result, err := provider.GetGormProtoMap()

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGormProtoDataProvider_GetProtoImportsMap(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(t *testing.T) *config.GormGenProtoConfig
		expected    map[string]string
		expectError bool
	}{
		{
			name: "successful imports loading",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				tmpDir := t.TempDir()
				importFile := filepath.Join(tmpDir, "import.json")
				require.NoError(t, os.WriteFile(importFile, []byte(`{"timestamp":"google/protobuf/timestamp.proto"}`), 0644))
				return &config.GormGenProtoConfig{PathMaps: tmpDir}
			},
			expected:    map[string]string{"timestamp": "google/protobuf/timestamp.proto"},
			expectError: false,
		},
		{
			name: "missing import file",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				return &config.GormGenProtoConfig{PathMaps: t.TempDir()}
			},
			expectError: true,
		},
		{
			name: "invalid import file content",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				tmpDir := t.TempDir()
				importFile := filepath.Join(tmpDir, "import.json")
				require.NoError(t, os.WriteFile(importFile, []byte(`{invalid}`), 0644))
				return &config.GormGenProtoConfig{PathMaps: tmpDir}
			},
			expectError: true,
		},
		{
			name: "empty imports file",
			setup: func(t *testing.T) *config.GormGenProtoConfig {
				tmpDir := t.TempDir()
				importFile := filepath.Join(tmpDir, "import.json")
				require.NoError(t, os.WriteFile(importFile, []byte(""), 0644))
				return &config.GormGenProtoConfig{PathMaps: tmpDir}
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := tt.setup(t)
			provider := NewGormProtoDataProvider(conf)

			result, err := provider.GetProtoImportsMap()

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
