package categoryss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/categorypb"
)

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
	if dbItem.Parent != nil {
		pbModel.Parent = o.SummaryDbToPb(dbItem.Parent)
	}
	return &pbModel
}

func (o *Converter) SummaryDbToPb(dbItem *DbModel) *pb.Summary {
	var pbParentId uint64
	if dbItem.ParentID != nil {
		pbParentId = uint64(*dbItem.ParentID)
	}
	return &pb.Summary{
		Id:       uint64(dbItem.ID),
		Name:     dbItem.Name,
		ParentId: pbParentId,
	}
}

func (o *Converter) SummaryPbToDb(pbItem *pb.Summary) *DbModel {
	var dbParentId *uint
	if pbItem.ParentId == 0 {
		dbParentId = nil
	} else {
		dbParentIdTmp := uint(pbItem.ParentId)
		dbParentId = &dbParentIdTmp
	}
	return &DbModel{
		Model: gorm.Model{
			ID: uint(pbItem.Id),
		},
		ParentID: dbParentId,
		Name:     pbItem.Name,
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	uintItemParentID := uint(pbItem.ParentId)
	dbModel.ParentID = &uintItemParentID
	dbModel.Parent = o.SummaryPbToDb(pbItem.Parent)
	return &dbModel
}
