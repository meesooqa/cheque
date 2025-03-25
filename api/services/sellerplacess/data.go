package sellerplacess

import pb "receipt-002/api/gen/pb/sellerplacepb/v1"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:       uint64(dbItem.ID),
		SellerId: uint64(dbItem.SellerID),
		Name:     dbItem.Name,
		Address:  dbItem.Address,
		Email:    dbItem.Email,
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		SellerID: uint(pbItem.SellerId),
		Name:     pbItem.Name,
		Address:  pbItem.Address,
		Email:    pbItem.Email,
	}
}
