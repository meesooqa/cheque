package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api"
	"github.com/meesooqa/cheque/api-server/grpc"
	"github.com/meesooqa/cheque/api-server/web"
	"github.com/meesooqa/cheque/api-server/web/handlers"
	"github.com/meesooqa/cheque/api-server/web/middlewares"
	"github.com/meesooqa/cheque/api-server/web/server"
	"github.com/meesooqa/cheque/common/common_log"
	"github.com/meesooqa/cheque/common/config"
)

func main() {
	logger := common_log.InitConsoleLogger(slog.LevelDebug)
	conf, err := config.GetConf()
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
	grpcSrv := grpc.NewServer(logger, conf.GrpcServer, ss)
	err = grpcSrv.Run()
	if err != nil {
		log.Fatal(err)
	}

	hh := []web.Handler{
		// grpc-gateway (REST)
		handlers.NewGrpcGateway(logger, conf.GrpcServer, ss),
		// swagger
		handlers.NewSwagger(logger),
	}

	mws := []web.HandlerMiddleware{
		// order matters
		// middlewares.NewLog(logger),
		middlewares.NewCORS(conf.Server.CORS),
	}
	srv := server.NewServer(logger, conf.Server, hh, mws)
	err = srv.Run()
	logger.Error("server terminated", "err", err)
}
