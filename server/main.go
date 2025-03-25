package main

import (
	"log"
	"log/slog"

	"receipt-002/api"
	"receipt-002/common/common_log"
	"receipt-002/common/config"
	"receipt-002/db/db_provider"
	"receipt-002/server/grpc"
	"receipt-002/server/web"
	"receipt-002/server/web/handlers"
	"receipt-002/server/web/middlewares"
	"receipt-002/server/web/server"
)

func main() {
	loggerProvider := common_log.NewConsoleLoggerProvider(slog.LevelDebug)
	logger, cleanup := loggerProvider.GetLogger()
	defer cleanup()
	configProvider := config.NewDefaultConfigProvider()
	conf, err := configProvider.GetConf()
	if err != nil {
		log.Fatal(err)
	}

	ss := api.GetServiceServers(db_provider.NewDefaultDBProvider())
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
		// product photos
		handlers.NewMedia(logger, conf.System),
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
