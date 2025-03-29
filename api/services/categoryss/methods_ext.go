package categoryss

import (
	"context"

	pb "github.com/meesooqa/cheque/api/gen/pb/categorypb/v1"
)

func (o *ServiceServer) GetChildren(ctx context.Context, req *pb.GetChildrenRequest) (*pb.GetChildrenResponse, error) {
	// TODO tests
	items, total, err := o.BaseService.GetList(ctx, o.FilterProvider.(*FilterProvider).GetChildrenFilters(req), req.SortBy, req.SortOrder, int(req.PageSize), int(req.Page))
	if err != nil {
		return nil, err
	}
	return &pb.GetChildrenResponse{
		Total: total,
		Items: items,
	}, nil
}
