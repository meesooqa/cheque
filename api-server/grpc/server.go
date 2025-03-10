package grpc

import (
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/meesooqa/cheque/common/api"
	"github.com/meesooqa/cheque/common/config"
)

type Server struct {
	logger         *slog.Logger
	conf           *config.GrpcServerConfig
	serviceServers []api.ServiceServer
}

func NewServer(logger *slog.Logger, conf *config.GrpcServerConfig, serviceServers []api.ServiceServer) *Server {
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
