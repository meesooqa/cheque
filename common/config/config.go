package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Conf from config yml
type Conf struct {
	System     *SystemConfig     `yaml:"system"`
	Server     *ServerConfig     `yaml:"server"`
	GrpcServer *GrpcServerConfig `yaml:"grpc_server"`
	DB         *DbConfig         `yaml:"db"`
	// TODO Log.Level, Log.OtputFormat (text, json)...
}

// DbConfig - DB parameters
type DbConfig struct {
	IsDebugMode bool   `yaml:"is_debug_mode"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	SslMode     string `yaml:"sslmode"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DbName      string `yaml:"dbname"`
}

// ServerConfig - server parameters
type ServerConfig struct {
	Host string   `yaml:"host"`
	Port int      `yaml:"port"`
	CORS []string `yaml:"cors"`
}

// GrpcServerConfig - server parameters
type GrpcServerConfig struct {
	Endpoint string `yaml:"endpoint"`
}

// SystemConfig - system parameters
type SystemConfig struct {
	DataPath          string `yaml:"data_path"`
	UploadPath        string `yaml:"upload_path"`
	MaxUploadFileSize int    `yaml:"max_upload_file_size"`
}

// load config from file
func load(fname string) (res *Conf, err error) {
	res = &Conf{}
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}
