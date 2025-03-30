package productss

import (
	"context"

	pb "github.com/meesooqa/cheque/api/gen/pb/productpb/v1"
	"github.com/meesooqa/cheque/db/db_provider"
	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
	"github.com/meesooqa/cheque/db/repositories"
)

type Converter struct {
	categoryRepo db_types.Repository[models.Category]
	imageRepo    db_types.Repository[models.Image]
}

func NewConverter() *Converter {
	dbProvider := db_provider.NewDefaultDBProvider()
	return &Converter{
		categoryRepo: repositories.NewCategoryRepository(dbProvider),
		imageRepo:    repositories.NewImageRepository(dbProvider),
	}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:   uint64(dbItem.ID),
		Name: dbItem.Name,
	}
	if dbItem.BrandID != nil {
		pbModel.BrandId = uint64(*dbItem.BrandID)
	}
	if len(dbItem.Categories) > 0 {
		var categoriesId []uint64
		for _, category := range dbItem.Categories {
			categoriesId = append(categoriesId, uint64(category.ID))
		}
		pbModel.CategoriesId = categoriesId
	}
	if len(dbItem.Images) > 0 {
		var imagesId []uint64
		for _, image := range dbItem.Images {
			imagesId = append(imagesId, uint64(image.ID))
		}
		pbModel.ImagesId = imagesId
	}
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
	if len(pbItem.CategoriesId) > 0 {
		var categories []*models.Category
		for _, categoryId := range pbItem.CategoriesId {
			category, _ := o.categoryRepo.Get(context.TODO(), categoryId)
			categories = append(categories, category)
		}
		dbModel.Categories = categories
	}
	if len(pbItem.ImagesId) > 0 {
		var images []*models.Image
		for _, imageId := range pbItem.ImagesId {
			image, _ := o.imageRepo.Get(context.TODO(), imageId)
			images = append(images, image)
		}
		dbModel.Images = images
	}
	return &dbModel
}
