package productss

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "receipt-002/api/gen/pb/productpb/v1"
	"receipt-002/api/services"
	"receipt-002/db/db_types"
	"receipt-002/db/models"
	"receipt-002/db/repositories"
)

type DbModel = models.Product

type ServiceServer struct {
	*services.BaseService[DbModel, pb.Model]
	pb.UnimplementedModelServiceServer
}

func NewServiceServer(dbProvider db_types.DBProvider) *ServiceServer {
	base := services.NewBaseService[DbModel, pb.Model](repositories.NewProductRepository(dbProvider), NewConverter())
	return &ServiceServer{BaseService: base}
}

func (o *ServiceServer) Register(grpcServer *grpc.Server) {
	pb.RegisterModelServiceServer(grpcServer, o)
}

func (o *ServiceServer) RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return pb.RegisterModelServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
