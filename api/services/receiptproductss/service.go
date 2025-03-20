package receiptproductss

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/meesooqa/cheque/api/pb/receiptproductpb"
	"github.com/meesooqa/cheque/api/services"
	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/repositories"
)

type DbModel = models.ReceiptProduct

type ServiceServer struct {
	*services.BaseService[DbModel, pb.Model]
	pb.UnimplementedModelServiceServer
}

func NewServiceServer(dbProvider db_types.DBProvider) *ServiceServer {
	base := services.NewBaseService[DbModel, pb.Model](repositories.NewReceiptProductRepository(dbProvider), NewConverter())
	return &ServiceServer{BaseService: base}
}

func (o *ServiceServer) Register(grpcServer *grpc.Server) {
	pb.RegisterModelServiceServer(grpcServer, o)
}

func (o *ServiceServer) RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return pb.RegisterModelServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
