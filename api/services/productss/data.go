package productss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/pb/brandpb"
	"github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/api/pb/imagepb"
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
		for i, dbItem := range dbItems {
			var pbParentId uint64
			if dbItem.ParentID != nil {
				pbParentId = uint64(*dbItem.ParentID)
			}
			pbItems[i] = &categorypb.Summary{
				Id:       uint64(dbItem.ID),
				Name:     dbItem.Name,
				ParentId: pbParentId,
			}
		}
		return pbItems
	}
	return nil
}

func (o *Converter) getPbImages(dbItems []models.Image) []*imagepb.Model {
	if len(dbItems) > 0 {
		pbItems := make([]*imagepb.Model, len(dbItems))
		for i, dbItem := range dbItems {
			pbItems[i] = &imagepb.Model{
				Id:        uint64(dbItem.ID),
				ProductId: uint64(dbItem.ProductID),
				Name:      dbItem.Name,
				Url:       dbItem.URL,
				Order:     int32(dbItem.Order),
				IsMain:    dbItem.IsMain,
			}
		}
		return pbItems
	}
	return nil
}
