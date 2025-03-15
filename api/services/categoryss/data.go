package categoryss

import (
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/common/models"
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
		pbModel.Parent = &pb.Summary{
			Id:       uint64(*dbItem.ParentID),
			Name:     dbItem.Parent.Name,
			ParentId: uint64(*dbItem.ParentID),
		}
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	uintItemParentID := uint(pbItem.ParentId)
	dbModel.ParentID = &uintItemParentID
	dbModel.Parent = &models.Category{
		Model: gorm.Model{
			ID: uint(pbItem.Parent.Id),
		},
		Name: pbItem.Parent.Name,
	}
	return &dbModel
}
