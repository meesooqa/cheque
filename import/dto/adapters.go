package dto

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/meesooqa/cheque/common/models"
)

type DtoAdapter struct{}

func NewDtoAdapter() *DtoAdapter {
	return &DtoAdapter{}
}

func (a *DtoAdapter) Convert(rawDataDTO RawDataDTO) (*models.Receipt, error) {
	receiptDTO := rawDataDTO.Ticket.Document.Receipt
	result := &models.Receipt{
		ExternalIdentifier:   rawDataDTO.ID,
		DateTime:             a.parseDateTime(receiptDTO.DateTime),
		FiscalDocumentNumber: receiptDTO.FiscalDocumentNumber,
		FiscalDriveNumber:    receiptDTO.FiscalDriveNumber,
		FiscalSign:           receiptDTO.FiscalSign,
		Sum:                  receiptDTO.TotalSum,
		KktReg:               strings.TrimSpace(receiptDTO.KktRegID),
		BuyerPhoneOrAddress:  receiptDTO.BuyerPhoneOrAddress,
	}
	result.Operator = a.getOperator(receiptDTO)
	result.SellerPlace = a.getSellerPlace(receiptDTO)
	result.ReceiptProducts = a.getReceiptProducts(receiptDTO)
	return result, nil
}

func (a *DtoAdapter) getSellerPlace(receiptDTO ReceiptDTO) *models.SellerPlace {
	inn := strings.TrimSpace(receiptDTO.UserInn)
	sellerName := receiptDTO.User
	if sellerName == "" {
		sellerName = inn
	}
	if sellerName == "" {
		sellerName = uuid.NewString()
	}
	if receiptDTO.RetailPlace == "" {
		receiptDTO.RetailPlace = sellerName
	}
	return &models.SellerPlace{
		Name:    receiptDTO.RetailPlace,
		Address: receiptDTO.RetailPlaceAddress,
		Email:   receiptDTO.SellerAddress,
		Seller:  models.Seller{Name: sellerName, Inn: inn},
	}
}

func (a *DtoAdapter) getOperator(receiptDTO ReceiptDTO) *models.Operator {
	var result *models.Operator = nil
	var name = strings.TrimSpace(receiptDTO.Operator)
	if name != "" {
		result = &models.Operator{Name: name}
	}
	return result
}

func (a *DtoAdapter) getReceiptProducts(receiptDTO ReceiptDTO) []models.ReceiptProduct {
	var result []models.ReceiptProduct
	for _, itemDTO := range receiptDTO.Items {
		var productCodeData *string
		if itemDTO.ProductCodeData != nil {
			jsonData, _ := json.Marshal(itemDTO.ProductCodeData)
			strData := string(jsonData)
			productCodeData = &strData
		}

		result = append(result, models.ReceiptProduct{
			Price:           itemDTO.Price,
			Quantity:        itemDTO.Quantity,
			Sum:             itemDTO.Sum,
			ProductCodeData: productCodeData,
			Product:         models.Product{Name: itemDTO.Name},
		})
	}
	return result
}

func (a *DtoAdapter) parseDateTime(dateTimeStr string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05", dateTimeStr)
	return t
}
