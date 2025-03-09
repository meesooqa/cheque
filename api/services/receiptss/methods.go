package receiptss

import (
	"context"

	"cheque-04/api/services"
	pb "cheque-04/common/proto/receiptpb"
)

func (o *ServiceServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := o.BaseService.GetItem(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemResponse{Item: item}, nil
}

func (o *ServiceServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	item, err := o.BaseService.CreateItem(req.Item)
	if err != nil {
		return nil, err
	}
	return &pb.CreateItemResponse{Item: item}, nil
}

func (o *ServiceServer) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	item, err := o.BaseService.UpdateItem(req.Id, req.Item)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateItemResponse{Item: item}, nil
}

func (o *ServiceServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	err := o.BaseService.DeleteItem(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteItemResponse{}, nil
}

func (o *ServiceServer) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {
	filters := []services.FilterFunc{
		ExampleFilter(""),
	}
	items, total, err := o.BaseService.GetList(filters, req.SortBy, req.SortOrder, int(req.PageSize), int(req.Page))
	if err != nil {
		return nil, err
	}
	return &pb.GetListResponse{
		Total: total,
		Items: items,
	}, nil
}
