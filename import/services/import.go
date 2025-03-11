package services

import (
	"encoding/json"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/import/dto"
	"github.com/meesooqa/cheque/import/processors"
)

type ImportService struct {
	adapter   *dto.DtoAdapter
	processor *processors.ReceiptProcessor
}

func NewImportService() *ImportService {
	return &ImportService{
		adapter:   dto.NewDtoAdapter(),
		processor: processors.NewReceiptProcessor(),
	}
}

// SaveReceipt unmarshals the data, processes each receipt and saves them in a transaction.
func (o *ImportService) SaveReceipt(db *gorm.DB, data []byte) error {
	var rawData []dto.RawDataDTO
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// global caches Seller, SellerPlace, Operator
	// (Product uses local cache for every Receipt)
	sellerCache := make(map[string]uint)
	sellerPlaceCache := make(map[string]uint)
	operatorCache := make(map[string]uint)

	for _, item := range rawData {
		receipt, err := o.adapter.Convert(item)
		if err != nil {
			tx.Rollback()
			return err
		}

		if err = o.processor.Process(tx, receipt, sellerCache, sellerPlaceCache, operatorCache); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
