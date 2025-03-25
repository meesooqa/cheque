package services

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoogleProductTaxonomyAdapter_Convert(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	adapter := NewGoogleProductTaxonomyAdapter(logger)

	tests := []struct {
		name     string
		input    string
		expected *GoogleProductTaxonomyItem
		wantErr  bool
	}{
		{
			name:  "Root category",
			input: "5181 - Багаж и сумки",
			expected: &GoogleProductTaxonomyItem{
				ID:       5181,
				Name:     "Багаж и сумки",
				FullPath: "Багаж и сумки",
				Level:    0,
			},
			wantErr: false,
		},
		{
			name:  "Second level category",
			input: "110 - Багаж и сумки > Багажные принадлежности",
			expected: &GoogleProductTaxonomyItem{
				ID:       110,
				Name:     "Багажные принадлежности",
				FullPath: "Багаж и сумки > Багажные принадлежности",
				Level:    1,
			},
			wantErr: false,
		},
		{
			name:  "Third level category",
			input: "5652 - Багаж и сумки > Багажные принадлежности > Багажные ремни",
			expected: &GoogleProductTaxonomyItem{
				ID:       5652,
				Name:     "Багажные ремни",
				FullPath: "Багаж и сумки > Багажные принадлежности > Багажные ремни",
				Level:    2,
			},
			wantErr: false,
		},
		{
			name:     "Comment line",
			input:    "# Google_Product_Taxonomy_Version: 2021-09-21",
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Empty line",
			input:    "",
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Invalid format (no dash)",
			input:    "5181 Багаж и сумки",
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Invalid ID",
			input:    "abc - Багаж и сумки",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := adapter.Convert(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				item, ok := result.(*GoogleProductTaxonomyItem)
				require.True(t, ok, "Result should be of type *GoogleProductTaxonomyItem")

				assert.Equal(t, tt.expected.ID, item.ID)
				assert.Equal(t, tt.expected.Name, item.Name)
				assert.Equal(t, tt.expected.FullPath, item.FullPath)
				assert.Equal(t, tt.expected.Level, item.Level)
			}
		})
	}
}
