package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"cheque-04/common/config"
)

func TestNewFS(t *testing.T) {
	conf := &config.GormGenProtoConfig{ProtoRoot: "test"}
	fs := NewFS(conf)
	assert.NotNil(t, fs)
	assert.Equal(t, conf, fs.conf)
}

func TestFS_GetProtoDir(t *testing.T) {
	tests := []struct {
		name      string
		protoRoot string
		dirName   string
		expected  string
	}{
		{
			name:      "normal directory",
			protoRoot: "proto",
			dirName:   "user",
			expected:  "proto/user",
		},
		{
			name:      "empty dirName",
			protoRoot: "proto",
			dirName:   "",
			expected:  "proto",
		},
		{
			name:      "protoRoot with trailing slash",
			protoRoot: "proto/",
			dirName:   "user",
			expected:  "proto/user",
		},
		{
			name:      "empty protoRoot",
			protoRoot: "",
			dirName:   "user",
			expected:  "user",
		},
		{
			name:      "both empty",
			protoRoot: "",
			dirName:   "",
			expected:  "",
		},
		{
			name:      "protoRoot is root directory",
			protoRoot: "/",
			dirName:   "user",
			expected:  "/user",
		},
		{
			name:      "protoRoot is root with empty dirName",
			protoRoot: "/",
			dirName:   "",
			expected:  "/",
		},
		{
			name:      "protoRoot is root and dirName with slash",
			protoRoot: "/",
			dirName:   "/user",
			expected:  "/user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &config.GormGenProtoConfig{ProtoRoot: tt.protoRoot}
			fs := NewFS(conf)
			assert.Equal(t, tt.expected, fs.GetProtoDir(tt.dirName))
		})
	}
}

func TestFS_GetProtoFilePath(t *testing.T) {
	tests := []struct {
		name      string
		protoRoot string
		dirName   string
		fileStem  string
		expected  string
	}{
		{
			name:      "normal file path",
			protoRoot: "proto",
			dirName:   "user",
			fileStem:  "service",
			expected:  "proto/user/service.proto",
		},
		{
			name:      "empty dirName",
			protoRoot: "proto",
			dirName:   "",
			fileStem:  "empty",
			expected:  "proto/empty.proto",
		},
		{
			name:      "empty fileStem",
			protoRoot: "proto",
			dirName:   "user",
			fileStem:  "",
			expected:  "",
		},
		{
			name:      "protoRoot with trailing slash",
			protoRoot: "proto/",
			dirName:   "v1",
			fileStem:  "api",
			expected:  "proto/v1/api.proto",
		},
		{
			name:      "all empty",
			protoRoot: "",
			dirName:   "",
			fileStem:  "",
			expected:  "",
		},
		{
			name:      "root protoRoot with normal inputs",
			protoRoot: "/",
			dirName:   "user",
			fileStem:  "service",
			expected:  "/user/service.proto",
		},
		{
			name:      "root protoRoot with empty dirName",
			protoRoot: "/",
			dirName:   "",
			fileStem:  "config",
			expected:  "/config.proto",
		},
		{
			name:      "root protoRoot with empty fileStem",
			protoRoot: "/",
			dirName:   "system",
			fileStem:  "",
			expected:  "",
		},
		{
			name:      "root protoRoot with all empty",
			protoRoot: "/",
			dirName:   "",
			fileStem:  "",
			expected:  "",
		},
		{
			name:      "root protoRoot with slashes in dirName",
			protoRoot: "/",
			dirName:   "/internal/user",
			fileStem:  "v1_service",
			expected:  "/internal/user/v1_service.proto",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &config.GormGenProtoConfig{ProtoRoot: tt.protoRoot}
			fs := NewFS(conf)
			assert.Equal(t, tt.expected, fs.GetProtoFilePath(tt.dirName, tt.fileStem))
		})
	}
}
