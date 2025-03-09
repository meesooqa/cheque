package receiptproductss

import pb "cheque-04/common/proto/receiptproductpb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id: uint64(dbItem.ID),
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{}
}
