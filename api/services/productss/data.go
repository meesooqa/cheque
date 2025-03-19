package productss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/pb/brandpb"
	"github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/api/pb/imagepb"
	pb "github.com/meesooqa/cheque/api/pb/productpb"
	"github.com/meesooqa/cheque/api/services/categoryss"
	"github.com/meesooqa/cheque/api/services/imagess"
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
	pbModel.Categories = o.getPbCategories(dbItem.Categories)
	pbModel.Images = o.getPbImages(dbItem.Images)
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
	if pbItem.Brand != nil {
		dbModel.Brand = &models.Brand{
			Model: gorm.Model{
				ID: uint(pbItem.Brand.Id),
			},
			Name: pbItem.Brand.Name,
		}
	}
	// TODO Categories
	// TODO Images
	return &dbModel
}

func (o *Converter) getPbCategories(dbItems []models.Category) []*categorypb.Summary {
	if len(dbItems) > 0 {
		pbItems := make([]*categorypb.Summary, len(dbItems))
		converter := categoryss.NewConverter()
		for i, dbItem := range dbItems {
			pbItems[i] = converter.SummaryDbToPb(&dbItem)
		}
		return pbItems
	}
	return nil
}

func (o *Converter) getPbImages(dbItems []models.Image) []*imagepb.Model {
	if len(dbItems) > 0 {
		pbItems := make([]*imagepb.Model, len(dbItems))
		converter := imagess.NewConverter()
		for i, dbItem := range dbItems {
			pbItems[i] = converter.DataDbToPb(&dbItem)
		}
		return pbItems
	}
	return nil
}
