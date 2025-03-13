package processors

import (
	"errors"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/common/models"
)

// processOperator processes Operator. If Operator is filled,
// find or create record by name and save ref to receipt.OperatorId.
func (o *ReceiptProcessor) processOperator(db *gorm.DB, receipt *models.Receipt, cache map[string]uint) error {
	// skip if Operator is nil
	if receipt.Operator == nil || receipt.Operator.Name == "" {
		return nil
	}
	opName := receipt.Operator.Name
	if id, ok := cache[opName]; ok {
		receipt.OperatorId = &id
		receipt.Operator = &models.Operator{Model: gorm.Model{ID: id}, Name: opName}
		return nil
	}
	var existed models.Operator
	if err := db.Where("name = ?", opName).First(&existed).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = db.FirstOrCreate(&existed, models.Operator{Name: opName}).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	cache[opName] = existed.ID
	receipt.OperatorId = &existed.ID
	receipt.Operator = &models.Operator{Model: gorm.Model{ID: existed.ID}, Name: opName}
	return nil
}
