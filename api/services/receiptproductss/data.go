package receiptproductss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/pb/productpb"
	pb "github.com/meesooqa/cheque/api/pb/receiptproductpb"
	"github.com/meesooqa/cheque/db/models"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:        uint64(dbItem.ID),
		ProductId: uint64(dbItem.ProductID),
		ReceiptId: uint64(dbItem.ReceiptID),
		Price:     int32(dbItem.Price),
		Quantity:  dbItem.Quantity,
		Sum:       int32(dbItem.Sum),
	}
	pbModel.Product = &productpb.Summary{
		Id:   uint64(dbItem.Product.ID),
		Name: dbItem.Product.Name,
	}
	if dbItem.Product.BrandID != nil {
		pbModel.Product.BrandId = uint64(*dbItem.Product.BrandID)
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		ProductID: uint(pbItem.ProductId),
		ReceiptID: uint(pbItem.ReceiptId),
		Price:     int(pbItem.Price),
		Quantity:  pbItem.Quantity,
		Sum:       int(pbItem.Sum),
	}
	dbModel.Product = models.Product{
		Model: gorm.Model{
			ID: uint(pbItem.Product.Id),
		},
		Name: pbItem.Product.Name,
	}
	if pbItem.Product.BrandId != 0 {
		pbItemProductBrandId := uint(pbItem.Product.BrandId)
		dbModel.Product.BrandID = &pbItemProductBrandId
	}
	return &dbModel
}
