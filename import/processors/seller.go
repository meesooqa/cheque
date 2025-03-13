package processors

import (
	"errors"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/models"
)

// processSeller processes Seller and uses cache to avoid duplicates
func (o *ReceiptProcessor) processSeller(db *gorm.DB, receipt *models.Receipt, cache map[string]uint) error {
	item := &receipt.SellerPlace.Seller
	key := item.Name + "_" + item.Inn

	if id, ok := cache[key]; ok {
		receipt.SellerPlace.SellerId = id
		receipt.SellerPlace.Seller = models.Seller{}
		return nil
	}

	var existedItem models.Seller
	if err := db.Where("name = ? AND inn = ?", item.Name, item.Inn).First(&existedItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = db.FirstOrCreate(&existedItem, models.Seller{Name: item.Name, Inn: item.Inn}).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	receipt.SellerPlace.SellerId = existedItem.ID
	cache[key] = existedItem.ID
	receipt.SellerPlace.Seller = models.Seller{}
	return nil
}
