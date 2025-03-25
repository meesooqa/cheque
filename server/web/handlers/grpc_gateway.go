package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/meesooqa/cheque/common/common_api"
	"github.com/meesooqa/cheque/common/config"
)

type GrpcGateway struct {
	logger         *slog.Logger
	conf           *config.GrpcServerConfig
	serviceServers []common_api.ServiceServer
}

func NewGrpcGateway(logger *slog.Logger, conf *config.GrpcServerConfig, serviceServers []common_api.ServiceServer) *GrpcGateway {
	return &GrpcGateway{
		logger:         logger,
		conf:           conf,
		serviceServers: serviceServers,
	}
}

func (o *GrpcGateway) Handle(mux *http.ServeMux) error {
	ctx := context.Background()
	// grpc-gateway
	apiMux := runtime.NewServeMux()
	if len(o.serviceServers) > 0 {
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		for _, grpcServiceServer := range o.serviceServers {
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
