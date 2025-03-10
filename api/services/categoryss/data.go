package categoryss

import pb "cheque-04/common/proto/categorypb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:   uint64(dbItem.ID),
		Name: dbItem.Name,
	}
	if dbItem.ParentID != nil {
		pbModel.ParentID = uint64(*dbItem.ParentID)
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	uintItemParentID := uint(pbItem.ParentID)
	dbModel.ParentID = &uintItemParentID
	return &dbModel
}
