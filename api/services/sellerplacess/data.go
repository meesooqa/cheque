package sellerplacess

import pb "github.com/meesooqa/cheque/api/pb/sellerplacepb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:       uint64(dbItem.ID),
		SellerID: uint64(dbItem.SellerID),
		Name:     dbItem.Name,
		Address:  dbItem.Address,
		Email:    dbItem.Email,
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		SellerID: uint(pbItem.SellerID),
		Name:     pbItem.Name,
		Address:  pbItem.Address,
		Email:    pbItem.Email,
	}
}
