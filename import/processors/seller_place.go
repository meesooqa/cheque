package processors

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/models"
)

// processSellerPlace processes Seller Place using unique combination seller_id, name and address
func (o *ReceiptProcessor) processSellerPlace(db *gorm.DB, receipt *models.Receipt, cache map[string]uint) error {
	key := fmt.Sprintf("%d_%s_%s", receipt.SellerPlace.SellerID, receipt.SellerPlace.Name, receipt.SellerPlace.Address)

	if id, ok := cache[key]; ok {
		receipt.SellerPlaceID = &id
		receipt.SellerPlace = nil
		return nil
	}

	var existed models.SellerPlace
	if err := db.Where("seller_id = ? AND name = ? AND address = ?",
		receipt.SellerPlace.SellerID, receipt.SellerPlace.Name, receipt.SellerPlace.Address).
		First(&existed).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			item := models.SellerPlace{
				SellerID: receipt.SellerPlace.SellerID,
				Name:     receipt.SellerPlace.Name,
				Address:  receipt.SellerPlace.Address,
				Email:    receipt.SellerPlace.Email,
			}
			if err := db.FirstOrCreate(&item, item).Error; err != nil {
				return err
			}
			existed = item
		} else {
			return err
		}
	}
	cache[key] = existed.ID
	receipt.SellerPlaceID = &existed.ID
	receipt.SellerPlace = nil
	return nil
}
