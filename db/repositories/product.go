package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/db_types"
	"github.com/meesooqa/cheque/db/models"
)

type ProductRepository struct {
	BaseRepository[models.Product]
}

func NewProductRepository(dbProvider db_types.DBProvider) *ProductRepository {
	repo := &ProductRepository{BaseRepository[models.Product]{
		DBProvider: dbProvider,
		Preloads: []string{
			"Brand",
			"Categories",
			"Images",
		},
	}}
	repo.BaseRepository.Self = repo
	return repo
}

func (o *ProductRepository) UpdateAssociations(db *gorm.DB, item *models.Product, updatedData *models.Product) error {
	if updatedData.Categories != nil {
		// TODO category with ID == 0 (added new category)
		if err := db.Model(item).Association("Categories").Replace(updatedData.Categories); err != nil {
			return fmt.Errorf("failed to update product categories: %w", err)
		}
	}
	// TODO Images
	//if updatedData.Images != nil {
	//	if err := db.Model(item).Association("Images").Replace(updatedData.Images).Error; err != nil {
	//		return fmt.Errorf("failed to update product images: %w", err)
	//	}
	//}
	return nil
}
