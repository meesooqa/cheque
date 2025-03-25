package services

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"receipt-002/db/models"
	"receipt-002/import/dto"
)

// MockProcessor implements the ReceiptProcessorInterface for testing
type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(db *gorm.DB, receipt *models.Receipt, sellerCache, sellerPlaceCache map[string]uint) error {
	args := m.Called(db, receipt, sellerCache, sellerPlaceCache)
	return args.Error(0)
}

// Define a test suite for ImportService
type ImportServiceTestSuite struct {
	suite.Suite
	db      *gorm.DB
	service *ImportService
}

func TestImportServiceSuite(t *testing.T) {
	suite.Run(t, new(ImportServiceTestSuite))
}

func (s *ImportServiceTestSuite) SetupTest() {
	var err error
	// Create in-memory SQLite database for testing
	s.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(s.T(), err)

	// Migrate the schema
	err = s.db.AutoMigrate(
		&models.Seller{},
		&models.SellerPlace{},
		&models.Product{},
		&models.Receipt{},
		&models.ReceiptProduct{},
	)
	require.NoError(s.T(), err)

	// Initialize the import service
	s.service = NewImportService()
}

func (s *ImportServiceTestSuite) TearDownTest() {
	// Get the database connection
	sqlDB, err := s.db.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
}

func (s *ImportServiceTestSuite) TestNewImportService() {
	// Test the constructor function
	service := NewImportService()
	assert.NotNil(s.T(), service, "NewImportService should return a non-nil service")
}

func (s *ImportServiceTestSuite) TestSaveReceipt_InvalidJSON() {
	// Test with invalid JSON data
	invalidJSON := []byte(`not a valid json`)

	err := s.service.SaveReceipt(s.db, invalidJSON)

	assert.Error(s.T(), err, "SaveReceipt should return an error for invalid JSON")
	assert.Contains(s.T(), err.Error(), "invalid", "Error should mention the JSON is invalid")
}

func (s *ImportServiceTestSuite) TestSaveReceipt_EmptyArray() {
	// Test with an empty array
	emptyArray := []byte(`[]`)

	err := s.service.SaveReceipt(s.db, emptyArray)

	assert.NoError(s.T(), err, "SaveReceipt should not return an error for an empty array")

	// Verify no data was inserted
	var count int64
	s.db.Model(&models.Receipt{}).Count(&count)
	assert.Equal(s.T(), int64(0), count, "No receipts should be inserted for an empty array")
}

func (s *ImportServiceTestSuite) TestSaveReceipt_SingleReceipt() {
	// Create test data for a single receipt
	rawData := []dto.RawDataDTO{
		createTestRawDataDTO("receipt-001", "TestSeller", "123456", "TestShop", "Test Address"),
	}

	data, err := json.Marshal(rawData)
	require.NoError(s.T(), err)

	// Process the receipt
	err = s.service.SaveReceipt(s.db, data)
	require.NoError(s.T(), err)

	// Verify receipt was saved
	var receipt models.Receipt
	err = s.db.Where("external_identifier = ?", "receipt-001").First(&receipt).Error
	assert.NoError(s.T(), err)

	// Verify related entities were saved
	var seller models.Seller
	err = s.db.Where("name = ?", "TestSeller").First(&seller).Error
	assert.NoError(s.T(), err)

	var sellerPlace models.SellerPlace
	err = s.db.Where("name = ?", "TestShop").First(&sellerPlace).Error
	assert.NoError(s.T(), err)

	var product models.Product
	err = s.db.Where("name = ?", "TestProduct").First(&product).Error
	assert.NoError(s.T(), err)

	var receiptProduct models.ReceiptProduct
	err = s.db.Where("receipt_id = ?", receipt.ID).First(&receiptProduct).Error
	assert.NoError(s.T(), err)
}

func (s *ImportServiceTestSuite) TestSaveReceipt_MultipleReceipts() {
	// Create test data with multiple receipts
	rawData := []dto.RawDataDTO{
		createTestRawDataDTO("receipt-002", "Seller1", "111111", "Shop1", "Address1"),
		createTestRawDataDTO("receipt-003", "Seller1", "111111", "Shop1", "Address1"), // Same seller and shop
		createTestRawDataDTO("receipt-004", "Seller2", "222222", "Shop2", "Address2"), // Different seller and shop
	}

	data, err := json.Marshal(rawData)
	require.NoError(s.T(), err)

	// Process the receipts
	err = s.service.SaveReceipt(s.db, data)
	require.NoError(s.T(), err)

	// Verify the right number of receipts were saved
	var receiptCount int64
	s.db.Model(&models.Receipt{}).Count(&receiptCount)
	assert.Equal(s.T(), int64(3), receiptCount, "Should have 3 receipts")

	// Verify the right number of sellers were saved (2 unique sellers)
	var sellerCount int64
	s.db.Model(&models.Seller{}).Count(&sellerCount)
	assert.Equal(s.T(), int64(2), sellerCount, "Should have 2 sellers")

	// Verify the right number of seller places were saved (2 unique places)
	var placeCount int64
	s.db.Model(&models.SellerPlace{}).Count(&placeCount)
	assert.Equal(s.T(), int64(2), placeCount, "Should have 2 seller places")
}

func (s *ImportServiceTestSuite) TestSaveReceipt_DuplicateReceipts() {
	// Create test data with duplicate receipts (same external identifier)
	rawData := []dto.RawDataDTO{
		createTestRawDataDTO("receipt-005", "Seller3", "333333", "Shop3", "Address3"),
		createTestRawDataDTO("receipt-005", "Seller3", "333333", "Shop3", "Address3"), // Duplicate receipt ID
	}

	data, err := json.Marshal(rawData)
	require.NoError(s.T(), err)

	// Process the receipts
	err = s.service.SaveReceipt(s.db, data)
	require.NoError(s.T(), err)

	// Verify only one receipt was saved
	var receiptCount int64
	s.db.Model(&models.Receipt{}).Count(&receiptCount)
	assert.Equal(s.T(), int64(1), receiptCount, "Should have only 1 receipt due to duplicate")
}

func (s *ImportServiceTestSuite) TestSaveReceipt_Transaction() {
	// Create a mock processor
	mockProcessor := new(MockProcessor)

	// Create a custom service with our mock processor
	customService := &ImportService{
		adapter:   dto.NewDtoAdapter(),
		processor: mockProcessor,
	}

	// Setup the mock to return error on the second call
	mockProcessor.On("Process", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	mockProcessor.On("Process", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("processing error")).Once()

	// Create test data with two receipts
	rawData := []dto.RawDataDTO{
		createTestRawDataDTO("receipt-tx-1", "SellerTx1", "111111", "ShopTx1", "AddressTx1"),
		createTestRawDataDTO("receipt-tx-2", "SellerTx2", "222222", "ShopTx2", "AddressTx2"),
	}

	data, err := json.Marshal(rawData)
	require.NoError(s.T(), err)

	// Process the receipts - should fail on the second one
	err = customService.SaveReceipt(s.db, data)

	// Verify error was returned
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "processing error")

	// Verify that transaction was rolled back (no receipts were saved)
	var receiptCount int64
	s.db.Model(&models.Receipt{}).Count(&receiptCount)
	assert.Equal(s.T(), int64(0), receiptCount, "No receipts should be saved when transaction is rolled back")

	// Verify the mock processor was called as expected
	mockProcessor.AssertExpectations(s.T())
}

func (s *ImportServiceTestSuite) TestSaveReceipt_DBBeginError() {
	// Create a broken database connection to simulate a Begin error
	brokenDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(s.T(), err)

	// Close the database to force a Begin error
	sqlDB, err := brokenDB.DB()
	require.NoError(s.T(), err)
	err = sqlDB.Close()
	require.NoError(s.T(), err)

	// Try to process a receipt, which should fail at Begin()
	err = s.service.SaveReceipt(brokenDB, []byte(`[]`))
	assert.Error(s.T(), err, "Should return an error when database Begin fails")
}

// Helper function to create test RawDataDTO objects
func createTestRawDataDTO(id, sellerName, inn, shopName, address string) dto.RawDataDTO {
	return dto.RawDataDTO{
		ID:        id,
		CreatedAt: time.Now(),
		Ticket: dto.TicketDTO{
			Document: struct {
				Receipt dto.ReceiptDTO `json:"receipt"`
			}{
				Receipt: dto.ReceiptDTO{
					DateTime:             "2023-05-15T14:30:00",
					FiscalDocumentNumber: 123456789,
					FiscalDriveNumber:    "FD12345",
					FiscalSign:           987654321,
					TotalSum:             1500,
					KktRegID:             "KKT123456",
					BuyerPhoneOrAddress:  "+7123456789",
					Operator:             "John Doe",
					User:                 sellerName,
					UserInn:              inn,
					RetailPlace:          shopName,
					RetailPlaceAddress:   address,
					SellerAddress:        "seller@example.com",
					Items: []dto.TicketItemDTO{
						{
							Name:     "TestProduct",
							Price:    500,
							Quantity: 1.0,
							Sum:      500,
						},
					},
				},
			},
		},
	}
}
