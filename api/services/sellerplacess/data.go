package sellerplacess

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/pb/sellerpb"
	pb "github.com/meesooqa/cheque/api/pb/sellerplacepb"
	"github.com/meesooqa/cheque/common/models"
)

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
		Seller: &sellerpb.Summary{
			Id:   uint64(dbItem.Seller.ID),
			Name: dbItem.Seller.Name,
			Inn:  dbItem.Seller.Inn,
		},
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		SellerID: uint(pbItem.SellerId),
		Name:     pbItem.Name,
		Address:  pbItem.Address,
		Email:    pbItem.Email,
		Seller: models.Seller{
			Model: gorm.Model{
				ID: uint(pbItem.Seller.Id),
			},
			Name: pbItem.Seller.Name,
			Inn:  pbItem.Seller.Inn,
		},
	}
}
