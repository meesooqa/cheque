package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Conf from config yml
type Conf struct {
	System       *SystemConfig       `yaml:"system"`
	Server       *ServerConfig       `yaml:"server"`
	GrpcServer   *GrpcServerConfig   `yaml:"grpc_server"`
	DB           *DbConfig           `yaml:"db"`
	GormGenProto *GormGenProtoConfig `yaml:"gorm_gen_proto"`
}

// GormGenProtoConfig stores GormGenProto parameters
type GormGenProtoConfig struct {
	PathMaps     string `yaml:"path_maps"`
	PathTmpl     string `yaml:"path_tmpl"`
	ProtoRoot    string `yaml:"proto_root"`
	ProtocRoot   string `yaml:"protoc_root"`
	ServicesRoot string `yaml:"services_root"`
}

// DbConfig - DB parameters
type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	SslMode  string `yaml:"sslmode"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

// ServerConfig - server parameters
type ServerConfig struct {
	Port int `yaml:"port"`
}

// GrpcServerConfig - server parameters
type GrpcServerConfig struct {
	Endpoint string `yaml:"endpoint"`
}

// SystemConfig - system parameters
type SystemConfig struct {
	DataPath   string `yaml:"data_path"`
	UploadPath string `yaml:"upload_path"`
}

// GetConf provides Conf from default config file
func GetConf() (res *Conf, err error) {
	return load("etc/config.yml")
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
