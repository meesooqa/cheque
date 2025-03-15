package productss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/pb/brandpb"
	pb "github.com/meesooqa/cheque/api/pb/productpb"
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
	if dbItem.BrandID != nil {
		pbModel.BrandId = uint64(*dbItem.BrandID)
	}
	if dbItem.Brand != nil {
		pbModel.Brand = &brandpb.Model{
			Id:   uint64(*dbItem.BrandID),
			Name: dbItem.Brand.Name,
		}
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		Name: pbItem.Name,
	}
	uintItemBrandID := uint(pbItem.BrandId)
	dbModel.BrandID = &uintItemBrandID
	dbModel.Brand = &models.Brand{
		Model: gorm.Model{
			ID: uint(pbItem.Brand.Id),
		},
		Name: pbItem.Brand.Name,
	}
	return &dbModel
}
