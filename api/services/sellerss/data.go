package sellerss

import pb "receipt-002/api/gen/pb/sellerpb/v1"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:   uint64(dbItem.ID),
		Name: dbItem.Name,
		Inn:  dbItem.Inn,
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		Name: pbItem.Name,
		Inn:  pbItem.Inn,
	}
}
