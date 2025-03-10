package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/meesooqa/cheque/common/api"
	"github.com/meesooqa/cheque/common/config"
)

type GrpcHandler struct {
	logger             *slog.Logger
	conf               *config.GrpcServerConfig
	grpcServiceServers []api.ServiceServer
}

func NewGrpcHandler(logger *slog.Logger, conf *config.GrpcServerConfig, grpcServiceServers []api.ServiceServer) *GrpcHandler {
	return &GrpcHandler{
		logger:             logger,
		conf:               conf,
		grpcServiceServers: grpcServiceServers,
	}
}

func (o *GrpcHandler) Handle(mux *http.ServeMux) error {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	// grpc-gateway
	apiMux := runtime.NewServeMux()
	if len(o.grpcServiceServers) > 0 {
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		for _, grpcServiceServer := range o.grpcServiceServers {
			err := grpcServiceServer.RegisterFromEndpoint(ctx, apiMux, o.conf.Endpoint, opts)
			if err != nil {
				o.logger.Error("failed to register grpc gateway endpoint", slog.Any("err", err))
				return err
			}
		}
	}
	// router
	mux.Handle("/api/", apiMux)
	return nil
}
