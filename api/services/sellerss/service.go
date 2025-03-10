package sellerss

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	pb "cheque-04/api/pb/sellerpb"
	"cheque-04/api/services"
	"cheque-04/common/models"
)

type DbModel = models.Seller

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
