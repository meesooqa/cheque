package operatorss

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/meesooqa/cheque/api/pb/operatorpb"
	"github.com/meesooqa/cheque/api/services"
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/common/repositories"
)

type DbModel = models.Operator

type ServiceServer struct {
	*services.BaseService[DbModel, pb.Model]
	pb.UnimplementedModelServiceServer
}

func NewServiceServer() *ServiceServer {
	base := services.NewBaseService[DbModel, pb.Model](repositories.NewOperatorRepository(), NewConverter())
	return &ServiceServer{BaseService: base}
}

func (o *ServiceServer) Register(grpcServer *grpc.Server) {
	pb.RegisterModelServiceServer(grpcServer, o)
}

func (o *ServiceServer) RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return pb.RegisterModelServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
