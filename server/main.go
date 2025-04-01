package main

import (
	"log"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/meesooqa/cheque/api"
	pbImage "github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
	"github.com/meesooqa/cheque/common/common_log"
	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/db/db_provider"
	grpcServer "github.com/meesooqa/cheque/server/grpc"
	"github.com/meesooqa/cheque/server/web"
	"github.com/meesooqa/cheque/server/web/handlers"
	"github.com/meesooqa/cheque/server/web/middlewares"
	"github.com/meesooqa/cheque/server/web/server"
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
	grpcSrv := grpcServer.NewServer(logger, conf.GrpcServer, ss)
	err = grpcSrv.Run()
	if err != nil {
		log.Fatal(err)
	}

	// отдельное соединение для Upload
	// @see docs/upload.md
	conn, err := grpc.NewClient(conf.GrpcServer.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	defer conn.Close()
	imageServiceClient := pbImage.NewModelServiceClient(conn)

	hh := []web.Handler{
		// grpc-gateway (REST)
		handlers.NewGrpcGateway(logger, conf.GrpcServer, ss),
		// swagger
		handlers.NewSwagger(logger),
		// save photos
		handlers.NewUpload(logger, conf.System, imageServiceClient),
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
