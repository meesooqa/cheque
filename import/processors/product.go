package processors

import (
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/models"
)

// processProducts processes receipt products using receipt local cache.
func (o *ReceiptProcessor) processProducts(db *gorm.DB, receipt *models.Receipt) error {
	cache := make(map[string]uint)
	for i, rp := range receipt.ReceiptProducts {
		productName := rp.Product.Name
		if id, ok := cache[productName]; ok {
			receipt.ReceiptProducts[i].ProductId = id
			receipt.ReceiptProducts[i].Product = models.Product{}
			continue
		}
		var item models.Product
		if err := db.Where("name = ?", productName).
			FirstOrCreate(&item, models.Product{Name: productName}).Error; err != nil {
			return err
		}
		cache[productName] = item.ID
		receipt.ReceiptProducts[i].ProductId = item.ID
		receipt.ReceiptProducts[i].Product = models.Product{}
	}
	return nil
}
