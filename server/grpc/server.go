package grpc

import (
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"receipt-002/common/common_api"
	"receipt-002/common/config"
)

type Server struct {
	logger         *slog.Logger
	conf           *config.GrpcServerConfig
	serviceServers []common_api.ServiceServer
}

func NewServer(logger *slog.Logger, conf *config.GrpcServerConfig, serviceServers []common_api.ServiceServer) *Server {
	return &Server{
		logger:         logger,
		conf:           conf,
		serviceServers: serviceServers,
	}
}

func (o *Server) Run() error {
	// listener
	grpcLis, err := net.Listen("tcp", o.conf.Endpoint)
	if err != nil {
		return err
	}
	// init
	grpcServer := grpc.NewServer()
	// register services
	if len(o.serviceServers) > 0 {
		for _, ss := range o.serviceServers {
			ss.Register(grpcServer)
		}
	}
	// run
	go func() {
		o.logger.Info("gRPC server running", slog.String("endpoint", o.conf.Endpoint))
		err := grpcServer.Serve(grpcLis)
		if err != nil {
			o.logger.Error("failed to serve gRPC", slog.Any("err", err))
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()
	return nil
}
