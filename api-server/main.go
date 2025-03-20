package main

import (
	"log"
	"log/slog"

	"github.com/meesooqa/cheque/api"
	"github.com/meesooqa/cheque/api-server/grpc"
	"github.com/meesooqa/cheque/api-server/web"
	"github.com/meesooqa/cheque/api-server/web/handlers"
	"github.com/meesooqa/cheque/api-server/web/middlewares"
	"github.com/meesooqa/cheque/api-server/web/server"
	"github.com/meesooqa/cheque/common/common_log"
	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/db/db_provider"
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
