package operatorss

import pb "cheque-04/common/proto/operatorpb"

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
