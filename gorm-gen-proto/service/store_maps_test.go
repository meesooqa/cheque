package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStore(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(*testing.T) string
		wantErr      bool
		expectedData map[string]string
	}{
		{
			name: "successful load with valid JSON",
			setup: func(t *testing.T) string {
				tmpFile, err := os.CreateTemp(t.TempDir(), "valid-*.json")
				require.NoError(t, err)
				_, err = tmpFile.WriteString(`{"key1":"value1","key2":"value2"}`)
				require.NoError(t, err)
				return tmpFile.Name()
			},
			wantErr: false,
			expectedData: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "file not found",
			setup: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "non-existent.json")
			},
			wantErr: true,
		},
		{
			name: "invalid JSON syntax",
			setup: func(t *testing.T) string {
				tmpFile, err := os.CreateTemp(t.TempDir(), "invalid-*.json")
				require.NoError(t, err)
				_, err = tmpFile.WriteString(`{invalid-json}`)
				require.NoError(t, err)
				return tmpFile.Name()
			},
			wantErr: true,
		},
		{
			name: "empty file",
			setup: func(t *testing.T) string {
				tmpFile, err := os.CreateTemp(t.TempDir(), "empty-*.json")
				require.NoError(t, err)
				return tmpFile.Name()
			},
			wantErr: true,
		},
		{
			name: "JSON array instead of object",
			setup: func(t *testing.T) string {
				tmpFile, err := os.CreateTemp(t.TempDir(), "array-*.json")
				require.NoError(t, err)
				_, err = tmpFile.WriteString(`["item1","item2"]`)
				require.NoError(t, err)
				return tmpFile.Name()
			},
			wantErr: true,
		},
		{
			name: "non-string values in JSON",
			setup: func(t *testing.T) string {
				tmpFile, err := os.CreateTemp(t.TempDir(), "nonstring-*.json")
				require.NoError(t, err)
				_, err = tmpFile.WriteString(`{"key":123}`)
				require.NoError(t, err)
				return tmpFile.Name()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := tt.setup(t)
			store, err := NewStore(filename)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, store)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, store)
			assert.Equal(t, tt.expectedData, store.Data)
		})
	}
}
