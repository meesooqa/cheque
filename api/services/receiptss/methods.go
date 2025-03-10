package receiptss

import (
	"context"

	pb "github.com/meesooqa/cheque/api/pb/receiptpb"
	"github.com/meesooqa/cheque/api/services"
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
		ExternalIdentifierFilter(req.ExternalIdentifier),
		DateTimeFilter(req.DateTimeStart, req.DateTimeEnd),
		OperatorIDFilter(req.OperatorID),
		SellerPlaceIDFilter(req.SellerPlaceID),
		FiscalDocumentNumberFilter(req.FiscalDocumentNumberGt, req.FiscalDocumentNumberLt),
		FiscalDriveNumberFilter(req.FiscalDriveNumber),
		FiscalSignFilter(req.FiscalSignGt, req.FiscalSignLt),
		SumFilter(req.SumGt, req.SumLt),
		KktRegFilter(req.KktReg),
		BuyerPhoneOrAddressFilter(req.BuyerPhoneOrAddress),
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
