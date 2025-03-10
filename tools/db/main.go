package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/config"
)

func main() {
	cmdGormMigrate := flag.NewFlagSet("gorm:migrate", flag.ExitOnError)
	cmdCleanup := flag.NewFlagSet("cleanup", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'gorm:migrate' or 'cleanup' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gorm:migrate":
		cmdGormMigrate.Parse(os.Args[2:])
		gormMigrate()
	case "cleanup":
		cmdCleanup.Parse(os.Args[2:])
		cleanup()
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func getDd() *gorm.DB {
	// TODO env.CHEQUE_CONF
	conf, err := config.Load("etc/config.yml")
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s", conf.DB.Host, conf.DB.Port, conf.DB.SslMode, conf.DB.User, conf.DB.Password, conf.DB.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db.WithContext(context.TODO())
}
