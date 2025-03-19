package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type ProductRepository struct {
	common_db.BaseRepository[models.Product]
}

func NewProductRepository() *ProductRepository {
	repo := &ProductRepository{common_db.BaseRepository[models.Product]{
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
