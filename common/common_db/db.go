package common_db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/config"
)

func GetDB(ctx context.Context) *gorm.DB {
	conf, err := config.GetConf()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s", conf.DB.Host, conf.DB.Port, conf.DB.SslMode, conf.DB.User, conf.DB.Password, conf.DB.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	return db.WithContext(ctx)
}
