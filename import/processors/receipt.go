package processors

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/meesooqa/cheque/db/models"
)

type ReceiptProcessor struct{}

func NewReceiptProcessor() *ReceiptProcessor {
	return &ReceiptProcessor{}
}

// Process processes Receipt, including related entities
func (o *ReceiptProcessor) Process(db *gorm.DB, receipt *models.Receipt, sellerCache, sellerPlaceCache map[string]uint) error {
	if err := o.processSeller(db, receipt, sellerCache); err != nil {
		return err
	}
	if err := o.processSellerPlace(db, receipt, sellerPlaceCache); err != nil {
		return err
	}
	if err := o.processProducts(db, receipt); err != nil {
		return err
	}

	var existed models.Receipt
	if err := db.Where("external_identifier = ?", receipt.ExternalIdentifier).First(&existed).Error; err == nil {
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return db.Session(&gorm.Session{FullSaveAssociations: true}).
		Clauses(clause.OnConflict{DoNothing: true}).Create(receipt).Error
}
