package categoryss

import pb "github.com/meesooqa/cheque/api/pb/categorypb"

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
		pbModel.ParentId = uint64(*dbItem.ParentID)
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	uintItemParentID := uint(pbItem.ParentId)
	dbModel.ParentID = &uintItemParentID
	return &dbModel
}
