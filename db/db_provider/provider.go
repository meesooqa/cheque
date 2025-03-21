package db_provider

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/config"
)

// Variable to allow mocking gorm.Open
var gormOpen = func(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(dialector, config)
}

type DefaultDBProvider struct {
	configProvider config.ConfigProvider
}

func NewDefaultDBProvider() *DefaultDBProvider {
	return &DefaultDBProvider{
		configProvider: config.NewDefaultConfigProvider(),
	}
}

func (o *DefaultDBProvider) GetDB(ctx context.Context) *gorm.DB {
	conf, err := o.configProvider.GetConf()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}
	dsn := o.constructDSN(conf)
	db, err := gormOpen(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	if conf.DB.IsDebugMode {
		db = db.Debug()
	}
	return db.WithContext(ctx)
}

// constructDSN creates a PostgreSQL connection string from config
func (o *DefaultDBProvider) constructDSN(conf *config.Conf) string {
	return fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		conf.DB.Host, conf.DB.Port, conf.DB.SslMode, conf.DB.User, conf.DB.Password, conf.DB.DbName)
}
