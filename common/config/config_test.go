package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	c, err := Load("testdata/config.yml")

	require.NoError(t, err)

	assert.IsType(t, &SystemConfig{}, c.System)
	assert.Equal(t, "test/data", c.System.DataPath)
	assert.Equal(t, "test/upload", c.System.UploadPath)

	assert.IsType(t, &ServerConfig{}, c.Server)
	assert.Equal(t, 4321, c.Server.Port)

	assert.IsType(t, &GrpcServerConfig{}, c.GrpcServer)
	assert.Equal(t, "localhost:11111", c.GrpcServer.Endpoint)

	assert.IsType(t, &DbConfig{}, c.DB)
	assert.Equal(t, "localhost", c.DB.Host)
	assert.Equal(t, 1234, c.DB.Port)
	assert.Equal(t, "disable", c.DB.SslMode)
	assert.Equal(t, "admin", c.DB.User)
	assert.Equal(t, "adminpw", c.DB.Password)
	assert.Equal(t, "app_db", c.DB.DbName)
}

func TestLoadConfigNotFoundFile(t *testing.T) {
	r, err := Load("/tmp/64c6fe8b-4e59-421e-ac37-342f5e1fdaef.txt")
	assert.Nil(t, r)
	assert.EqualError(t, err, "open /tmp/64c6fe8b-4e59-421e-ac37-342f5e1fdaef.txt: no such file or directory")
}

func TestLoadConfigInvalidYaml(t *testing.T) {
	r, err := Load("testdata/file.txt")

	assert.Nil(t, r)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Not Yaml` into config.Conf")
}
