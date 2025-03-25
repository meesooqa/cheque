package imagess

import (
	"context"

	pb "github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
)

func (o *ServiceServer) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {
	items, total, err := o.BaseService.GetList(ctx, GetFilters(req), req.SortBy, req.SortOrder, int(req.PageSize), int(req.Page))
	if err != nil {
		return nil, err
	}
	return &pb.GetListResponse{
		Total: total,
		Items: items,
	}, nil
}

func (o *ServiceServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := o.BaseService.GetItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemResponse{Item: item}, nil
}

func (o *ServiceServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	item, err := o.BaseService.CreateItem(ctx, req.Item)
	if err != nil {
		return nil, err
	}
	return &pb.CreateItemResponse{Item: item}, nil
}

func (o *ServiceServer) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	item, err := o.BaseService.UpdateItem(ctx, req.Id, req.Item)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateItemResponse{Item: item}, nil
}

func (o *ServiceServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	err := o.BaseService.DeleteItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteItemResponse{}, nil
}
