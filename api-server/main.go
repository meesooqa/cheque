package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cheque-04/api"
	"cheque-04/api-server/grpc_server"
	"cheque-04/api-server/web"
	"cheque-04/api-server/web/handlers"
	"cheque-04/api-server/web/middlewares"
	"cheque-04/api-server/web/server"
	"cheque-04/common/common_log"
	"cheque-04/common/config"
)

func main() {
	logger := common_log.InitConsoleLogger(slog.LevelDebug)
	conf, err := config.Load("etc/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s", conf.DB.Host, conf.DB.Port, conf.DB.SslMode, conf.DB.User, conf.DB.Password, conf.DB.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.WithContext(context.TODO())

	ss := api.GetServiceServers(logger, db)
	grpcSrv := grpc_server.NewServer(logger, conf.GrpcServer, ss)
	err = grpcSrv.Run()
	if err != nil {
		log.Fatal(err)
	}

	hh := []web.Handler{
		// grpc-gateway (REST)
		handlers.NewGrpcHandler(logger, conf.GrpcServer, ss),
	}

	// TODO add to config
	middleware := middlewares.NewCORS([]string{
		"http://localhost:5173", // frontendrefine
	})
	srv := server.NewServer(logger, conf.Server, hh, middleware)
	err = srv.Run()
	logger.Error("server terminated", "err", err)
}
