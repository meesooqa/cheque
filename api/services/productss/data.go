package productss

import pb "github.com/meesooqa/cheque/api/pb/productpb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:   uint64(dbItem.ID),
		Name: dbItem.Name,
	}
	if dbItem.BrandId != nil {
		pbModel.BrandId = uint64(*dbItem.BrandId)
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	uintItemBrandId := uint(pbItem.BrandId)
	dbModel.BrandId = &uintItemBrandId
	return &dbModel
}
