package dto

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDtoAdapter(t *testing.T) {
	adapter := NewDtoAdapter()
	assert.NotNil(t, adapter, "NewDtoAdapter should return a non-nil adapter")
}

func TestDtoAdapter_Convert(t *testing.T) {
	adapter := NewDtoAdapter()

	t.Run("complete receipt conversion", func(t *testing.T) {
		// Setup test data
		rawData := createTestRawDataDTO()

		// Convert the data
		receipt, err := adapter.Convert(rawData)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, receipt)

		// Check receipt fields
		assert.Equal(t, "test_id", receipt.ExternalIdentifier)
		assert.Equal(t, time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC), receipt.DateTime)
		assert.Equal(t, "123456789", receipt.FiscalDocumentNumber)
		assert.Equal(t, "FD12345", receipt.FiscalDriveNumber)
		assert.Equal(t, "987654321", receipt.FiscalSign)
		assert.Equal(t, 1500, receipt.Sum)
		assert.Equal(t, "KKT123456", receipt.KktReg)
		assert.Equal(t, "+7123456789", receipt.BuyerPhoneOrAddress)
		assert.Equal(t, "John Doe", receipt.Operator)

		// Check seller place
		require.NotNil(t, receipt.SellerPlace)
		assert.Equal(t, "Shop Name", receipt.SellerPlace.Name)
		assert.Equal(t, "123 Main St", receipt.SellerPlace.Address)
		assert.Equal(t, "shop@example.com", receipt.SellerPlace.Email)

		// Check seller
		assert.Equal(t, "Test Shop", receipt.SellerPlace.Seller.Name)
		assert.Equal(t, "1234567890", receipt.SellerPlace.Seller.Inn)

		// Check receipt products
		require.Len(t, receipt.ReceiptProducts, 2)

		// Check first product
		assert.Equal(t, "Product 1", receipt.ReceiptProducts[0].Product.Name)
		assert.Equal(t, 500, receipt.ReceiptProducts[0].Price)
		assert.Equal(t, 2.0, receipt.ReceiptProducts[0].Quantity)
		assert.Equal(t, 1000, receipt.ReceiptProducts[0].Sum)
		assert.Nil(t, receipt.ReceiptProducts[0].ProductCodeData)

		// Check second product with product code data
		assert.Equal(t, "Product 2", receipt.ReceiptProducts[1].Product.Name)
		assert.Equal(t, 500, receipt.ReceiptProducts[1].Price)
		assert.Equal(t, 1.0, receipt.ReceiptProducts[1].Quantity)
		assert.Equal(t, 500, receipt.ReceiptProducts[1].Sum)
		require.NotNil(t, receipt.ReceiptProducts[1].ProductCodeData)
		assert.Contains(t, *receipt.ReceiptProducts[1].ProductCodeData, "code123")
	})

	t.Run("empty seller name uses inn", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		rawData.Ticket.Document.Receipt.User = ""

		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		require.NotNil(t, receipt)
		assert.Equal(t, "1234567890", receipt.SellerPlace.Seller.Name)
	})

	t.Run("empty seller name and inn generates UUID", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		rawData.Ticket.Document.Receipt.User = ""
		rawData.Ticket.Document.Receipt.UserInn = ""

		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		require.NotNil(t, receipt)
		assert.NotEmpty(t, receipt.SellerPlace.Seller.Name)
		// UUID format validation
		assert.Contains(t, receipt.SellerPlace.Seller.Name, "-")
		assert.Len(t, receipt.SellerPlace.Seller.Name, 36)
	})

	t.Run("empty retail place uses seller name", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		rawData.Ticket.Document.Receipt.RetailPlace = ""

		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		require.NotNil(t, receipt)
		assert.Equal(t, "Test Shop", receipt.SellerPlace.Name)
	})

	t.Run("whitespace trimming", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		rawData.ID = "  test_id_with_spaces  "
		rawData.Ticket.Document.Receipt.FiscalDriveNumber = "  FD12345  "
		rawData.Ticket.Document.Receipt.KktRegID = "  KKT123456  "
		rawData.Ticket.Document.Receipt.BuyerPhoneOrAddress = "  +7123456789  "
		rawData.Ticket.Document.Receipt.Operator = "  John Doe  "
		rawData.Ticket.Document.Receipt.UserInn = "  1234567890  "
		rawData.Ticket.Document.Receipt.Items[0].Name = "  Product 1  "

		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		require.NotNil(t, receipt)
		assert.Equal(t, "test_id_with_spaces", receipt.ExternalIdentifier)
		assert.Equal(t, "FD12345", receipt.FiscalDriveNumber)
		assert.Equal(t, "KKT123456", receipt.KktReg)
		assert.Equal(t, "+7123456789", receipt.BuyerPhoneOrAddress)
		assert.Equal(t, "John Doe", receipt.Operator)
		assert.Equal(t, "1234567890", receipt.SellerPlace.Seller.Inn)
		assert.Equal(t, "Product 1", receipt.ReceiptProducts[0].Product.Name)
	})
}

func TestDtoAdapter_ParseDateTime(t *testing.T) {
	adapter := NewDtoAdapter()

	t.Run("valid datetime", func(t *testing.T) {
		// Using reflection to access unexported method for testing
		// This is a private method, so we're testing it through its effects in Convert
		rawData := createTestRawDataDTO()
		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		expected := time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC)
		assert.Equal(t, expected, receipt.DateTime)
	})

	t.Run("invalid datetime results in zero time", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		rawData.Ticket.Document.Receipt.DateTime = "invalid-datetime"

		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		assert.Equal(t, time.Time{}, receipt.DateTime)
	})
}

func TestDtoAdapter_GetSellerPlace(t *testing.T) {
	adapter := NewDtoAdapter()

	// Testing through Convert since getSellerPlace is private
	t.Run("complete seller place data", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		require.NotNil(t, receipt.SellerPlace)
		assert.Equal(t, "Shop Name", receipt.SellerPlace.Name)
		assert.Equal(t, "123 Main St", receipt.SellerPlace.Address)
		assert.Equal(t, "shop@example.com", receipt.SellerPlace.Email)
		assert.Equal(t, "Test Shop", receipt.SellerPlace.Seller.Name)
		assert.Equal(t, "1234567890", receipt.SellerPlace.Seller.Inn)
	})
}

func TestDtoAdapter_GetReceiptProducts(t *testing.T) {
	adapter := NewDtoAdapter()

	// Testing through Convert since getReceiptProducts is private
	t.Run("products with and without product code data", func(t *testing.T) {
		rawData := createTestRawDataDTO()
		receipt, err := adapter.Convert(rawData)

		require.NoError(t, err)
		require.Len(t, receipt.ReceiptProducts, 2)

		// First product has no product code data
		assert.Equal(t, "Product 1", receipt.ReceiptProducts[0].Product.Name)
		assert.Equal(t, 500, receipt.ReceiptProducts[0].Price)
		assert.Equal(t, 2.0, receipt.ReceiptProducts[0].Quantity)
		assert.Equal(t, 1000, receipt.ReceiptProducts[0].Sum)
		assert.Nil(t, receipt.ReceiptProducts[0].ProductCodeData)

		// Second product has product code data
		assert.Equal(t, "Product 2", receipt.ReceiptProducts[1].Product.Name)
		assert.Equal(t, 500, receipt.ReceiptProducts[1].Price)
		assert.Equal(t, 1.0, receipt.ReceiptProducts[1].Quantity)
		assert.Equal(t, 500, receipt.ReceiptProducts[1].Sum)
		require.NotNil(t, receipt.ReceiptProducts[1].ProductCodeData)

		// Verify product code data JSON
		var productCodeMap map[string]interface{}
		err = json.Unmarshal([]byte(*receipt.ReceiptProducts[1].ProductCodeData), &productCodeMap)
		require.NoError(t, err)
		assert.Equal(t, "code123", productCodeMap["code"])
	})
}

// Helper function to create test data
func createTestRawDataDTO() RawDataDTO {
	productCodeData := map[string]string{"code": "code123"}

	return RawDataDTO{
		ID:        "test_id",
		CreatedAt: time.Now(),
		Ticket: TicketDTO{
			Document: struct {
				Receipt ReceiptDTO `json:"receipt"`
			}{
				Receipt: ReceiptDTO{
					DateTime:             "2023-05-15T14:30:00",
					FiscalDocumentNumber: 123456789,
					FiscalDriveNumber:    "FD12345",
					FiscalSign:           987654321,
					TotalSum:             1500,
					KktRegID:             "KKT123456",
					BuyerPhoneOrAddress:  "+7123456789",
					Operator:             "John Doe",
					User:                 "Test Shop",
					UserInn:              "1234567890",
					RetailPlace:          "Shop Name",
					RetailPlaceAddress:   "123 Main St",
					SellerAddress:        "shop@example.com",
					Items: []TicketItemDTO{
						{
							Name:     "Product 1",
							Price:    500,
							Quantity: 2.0,
							Sum:      1000,
						},
						{
							Name:            "Product 2",
							Price:           500,
							Quantity:        1.0,
							Sum:             500,
							ProductCodeData: productCodeData,
						},
					},
				},
			},
		},
	}
}
