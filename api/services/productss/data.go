package productss

import pb "github.com/meesooqa/cheque/api/gen/pb/productpb/v1"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:   uint64(dbItem.ID),
		Name: dbItem.Name,
	}
	if dbItem.BrandID != nil {
		pbModel.BrandId = uint64(*dbItem.BrandID)
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	if pbItem.BrandId != 0 {
		uintItemBrandID := uint(pbItem.BrandId)
		dbModel.BrandID = &uintItemBrandID
	}
	return &dbModel
}
