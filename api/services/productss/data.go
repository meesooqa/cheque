package productss

import (
	"github.com/meesooqa/cheque/api/pb/categorypb"
	"github.com/meesooqa/cheque/api/pb/imagepb"
	pb "github.com/meesooqa/cheque/api/pb/productpb"
	"github.com/meesooqa/cheque/api/services/brandss"
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
		converter := brandss.NewConverter()
		pbModel.Brand = converter.DataDbToPb(dbItem.Brand)
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
		brandConverter := brandss.NewConverter()
		dbModel.Brand = brandConverter.DataPbToDb(pbItem.Brand)
	}
	if pbItem.Categories != nil {
		categoryConverter := categoryss.NewConverter()
		for _, categorySummary := range pbItem.Categories {
			dbModel.Categories = append(dbModel.Categories, *categoryConverter.SummaryPbToDb(categorySummary))
		}
	}
	//if pbItem.Images != nil {
	//	imageConverter := imagess.NewConverter()
	//	for _, image := range pbItem.Images {
	//		dbModel.Images = append(dbModel.Images, *imageConverter.DataPbToDb(image))
	//	}
	//}
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
