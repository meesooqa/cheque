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
}

func NewConverter() *Converter {
	dbProvider := db_provider.NewDefaultDBProvider()
	return &Converter{
		categoryRepo: repositories.NewCategoryRepository(dbProvider),
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
	return &dbModel
}
