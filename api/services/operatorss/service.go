package operatorss

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"cheque-04/api/services"
	"cheque-04/common/models"
	pb "cheque-04/common/proto/operatorpb"
)

type DbModel = models.Operator

type ServiceServer struct {
	*services.BaseService[DbModel, pb.Model]
	pb.UnimplementedModelServiceServer
}

func NewServiceServer(log *slog.Logger, db *gorm.DB) *ServiceServer {
	base := services.NewBaseService[DbModel, pb.Model](log, db, NewConverter())
	return &ServiceServer{BaseService: base}
}

func (o *ServiceServer) Register(grpcServer *grpc.Server) {
	pb.RegisterModelServiceServer(grpcServer, o)
}

func (o *ServiceServer) RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return pb.RegisterModelServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
