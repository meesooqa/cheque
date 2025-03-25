package processors

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/meesooqa/cheque/db/models"
)

type ReceiptProcessorTestSuite struct {
	suite.Suite
	db        *gorm.DB
	processor *ReceiptProcessor
}

func TestReceiptProcessorSuite(t *testing.T) {
	suite.Run(t, new(ReceiptProcessorTestSuite))
}

func (s *ReceiptProcessorTestSuite) SetupTest() {
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

	// Initialize processor
	s.processor = NewReceiptProcessor()
}

func (s *ReceiptProcessorTestSuite) TearDownTest() {
	// Get the database connection
	sqlDB, err := s.db.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
}

func (s *ReceiptProcessorTestSuite) TestNewReceiptProcessor() {
	// Test the constructor function
	processor := NewReceiptProcessor()
	assert.NotNil(s.T(), processor, "NewReceiptProcessor should return a non-nil processor")
}

func (s *ReceiptProcessorTestSuite) TestProcess_NewReceipt() {
	// Prepare test data
	receipt := createTestReceipt("receipt-001", "Seller1", "123456", "ShopA", "Address1")

	// Initialize empty caches
	sellerCache := make(map[string]uint)
	sellerPlaceCache := make(map[string]uint)

	// Process the receipt
	err := s.processor.Process(s.db, receipt, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err)

	// Verify that the receipt was saved
	var savedReceipt models.Receipt
	err = s.db.Where("external_identifier = ?", "receipt-001").First(&savedReceipt).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), receipt.ExternalIdentifier, savedReceipt.ExternalIdentifier)
	assert.Equal(s.T(), receipt.Sum, savedReceipt.Sum)

	// Verify that seller was created
	var seller models.Seller
	err = s.db.First(&seller).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Seller1", seller.Name)
	assert.Equal(s.T(), "123456", seller.Inn)

	// Verify that seller place was created
	var sellerPlace models.SellerPlace
	err = s.db.First(&sellerPlace).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "ShopA", sellerPlace.Name)
	assert.Equal(s.T(), "Address1", sellerPlace.Address)
	assert.Equal(s.T(), seller.ID, sellerPlace.SellerID)

	// Verify that product was created
	var product models.Product
	err = s.db.First(&product).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Product1", product.Name)

	// Verify that receipt product was created
	var receiptProduct models.ReceiptProduct
	err = s.db.First(&receiptProduct).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), product.ID, receiptProduct.ProductID)
	assert.Equal(s.T(), savedReceipt.ID, receiptProduct.ReceiptID)
	assert.Equal(s.T(), 1000, receiptProduct.Price)
	assert.Equal(s.T(), 2.0, receiptProduct.Quantity)
	assert.Equal(s.T(), 2000, receiptProduct.Sum)

	// Verify cache was populated
	sellerKey := "Seller1_123456"
	assert.Contains(s.T(), sellerCache, sellerKey)
	assert.Equal(s.T(), seller.ID, sellerCache[sellerKey])

	sellerPlaceKey := fmt.Sprintf("%d_ShopA_Address1", seller.ID)
	assert.Contains(s.T(), sellerPlaceCache, sellerPlaceKey)
	assert.Equal(s.T(), sellerPlace.ID, sellerPlaceCache[sellerPlaceKey])
}

func (s *ReceiptProcessorTestSuite) TestProcess_DuplicateReceipt() {
	// Create and process the first receipt
	receipt1 := createTestReceipt("receipt-002", "Seller2", "654321", "ShopB", "Address2")
	sellerCache := make(map[string]uint)
	sellerPlaceCache := make(map[string]uint)

	err := s.processor.Process(s.db, receipt1, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err)

	// Count records before trying to insert duplicate
	var count int64
	s.db.Model(&models.Receipt{}).Count(&count)
	assert.Equal(s.T(), int64(1), count)

	// Try to process the same receipt again
	receipt2 := createTestReceipt("receipt-002", "Seller2", "654321", "ShopB", "Address2")
	err = s.processor.Process(s.db, receipt2, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err) // No error, just skips the duplicate

	// Verify that no new record was created
	s.db.Model(&models.Receipt{}).Count(&count)
	assert.Equal(s.T(), int64(1), count, "No new receipt should be created for duplicate")
}

func (s *ReceiptProcessorTestSuite) TestProcess_UseCache() {
	// Setup initial data and cache
	receipt1 := createTestReceipt("receipt-003", "Seller3", "111111", "ShopC", "Address3")
	sellerCache := make(map[string]uint)
	sellerPlaceCache := make(map[string]uint)

	err := s.processor.Process(s.db, receipt1, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err)

	// Get the seller and seller place IDs that should be in the cache now
	var seller models.Seller
	var sellerPlace models.SellerPlace
	s.db.Where("name = ?", "Seller3").First(&seller)
	s.db.Where("name = ?", "ShopC").First(&sellerPlace)

	// Create a new receipt with the same seller and seller place
	receipt2 := createTestReceipt("receipt-004", "Seller3", "111111", "ShopC", "Address3")

	// Process the new receipt - this should use the cache
	err = s.processor.Process(s.db, receipt2, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err)

	// Verify cache was used by checking that only one seller and one seller place exist
	var sellerCount, sellerPlaceCount int64
	s.db.Model(&models.Seller{}).Count(&sellerCount)
	s.db.Model(&models.SellerPlace{}).Count(&sellerPlaceCount)

	assert.Equal(s.T(), int64(1), sellerCount, "Should be only one seller")
	assert.Equal(s.T(), int64(1), sellerPlaceCount, "Should be only one seller place")

	// Verify that cache entries exist
	sellerKey := "Seller3_111111"
	sellerPlaceKey := fmt.Sprintf("%d_ShopC_Address3", seller.ID)

	assert.Contains(s.T(), sellerCache, sellerKey)
	assert.Contains(s.T(), sellerPlaceCache, sellerPlaceKey)
	assert.Equal(s.T(), seller.ID, sellerCache[sellerKey])
	assert.Equal(s.T(), sellerPlace.ID, sellerPlaceCache[sellerPlaceKey])
}

func (s *ReceiptProcessorTestSuite) TestProcessSeller() {
	// Create test receipt and empty cache
	receipt := createTestReceipt("receipt-005", "Seller4", "222222", "ShopD", "Address4")
	sellerCache := make(map[string]uint)

	// Process the seller
	err := s.processor.Process(s.db, receipt, sellerCache, make(map[string]uint))
	require.NoError(s.T(), err)

	// Verify seller was created
	var seller models.Seller
	err = s.db.Where("name = ? AND inn = ?", "Seller4", "222222").First(&seller).Error
	require.NoError(s.T(), err)

	// Verify cache was populated
	sellerKey := "Seller4_222222"
	assert.Contains(s.T(), sellerCache, sellerKey)
	assert.Equal(s.T(), seller.ID, sellerCache[sellerKey])

	// Create a new receipt with the same seller
	receipt2 := createTestReceipt("receipt-006", "Seller4", "222222", "ShopE", "Address5")

	// Process the seller again - should use cache
	err = s.processor.Process(s.db, receipt2, sellerCache, make(map[string]uint))
	require.NoError(s.T(), err)

	// Verify only one seller exists
	var count int64
	s.db.Model(&models.Seller{}).Where("name = ? AND inn = ?", "Seller4", "222222").Count(&count)
	assert.Equal(s.T(), int64(1), count, "Should still have only one seller")
}

func (s *ReceiptProcessorTestSuite) TestProcessSellerPlace() {
	// Create test receipt
	receipt := createTestReceipt("receipt-007", "Seller5", "333333", "ShopF", "Address6")
	sellerCache := make(map[string]uint)
	sellerPlaceCache := make(map[string]uint)

	// Process the seller place
	err := s.processor.Process(s.db, receipt, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err)

	// Verify seller place was created
	var sellerPlace models.SellerPlace
	err = s.db.Where("name = ? AND address = ?", "ShopF", "Address6").First(&sellerPlace).Error
	require.NoError(s.T(), err)

	// Get the seller ID
	var seller models.Seller
	s.db.Where("name = ?", "Seller5").First(&seller)

	// Verify cache was populated
	sellerPlaceKey := fmt.Sprintf("%d_ShopF_Address6", seller.ID)
	assert.Contains(s.T(), sellerPlaceCache, sellerPlaceKey)
	assert.Equal(s.T(), sellerPlace.ID, sellerPlaceCache[sellerPlaceKey])

	// Create a new receipt with the same seller place
	receipt2 := createTestReceipt("receipt-008", "Seller5", "333333", "ShopF", "Address6")

	// Process again - should use cache
	err = s.processor.Process(s.db, receipt2, sellerCache, sellerPlaceCache)
	require.NoError(s.T(), err)

	// Verify only one seller place exists
	var count int64
	s.db.Model(&models.SellerPlace{}).Where("name = ? AND address = ?", "ShopF", "Address6").Count(&count)
	assert.Equal(s.T(), int64(1), count, "Should still have only one seller place")
}

func (s *ReceiptProcessorTestSuite) TestProcessProducts() {
	// Create test receipt with two products with the same name
	receipt := &models.Receipt{
		ExternalIdentifier: "receipt-009",
		DateTime:           time.Now(),
		Sum:                3000,
		SellerPlace: &models.SellerPlace{
			Name:    "ShopG",
			Address: "Address7",
			Seller: models.Seller{
				Name: "Seller6",
				Inn:  "444444",
			},
		},
		ReceiptProducts: []models.ReceiptProduct{
			{
				Price:    1000,
				Quantity: 1.0,
				Sum:      1000,
				Product:  models.Product{Name: "Product2"},
			},
			{
				Price:    1000,
				Quantity: 2.0,
				Sum:      2000,
				Product:  models.Product{Name: "Product2"}, // Same product name
			},
		},
	}

	// Process the receipt
	err := s.processor.Process(s.db, receipt, make(map[string]uint), make(map[string]uint))
	require.NoError(s.T(), err)

	// Verify only one product was created
	var productCount int64
	s.db.Model(&models.Product{}).Count(&productCount)
	assert.Equal(s.T(), int64(1), productCount, "Should have only one product")

	// But two receipt products
	var receiptProductCount int64
	s.db.Model(&models.ReceiptProduct{}).Count(&receiptProductCount)
	assert.Equal(s.T(), int64(2), receiptProductCount, "Should have two receipt products")

	// Both receipt products should reference the same product ID
	var products []models.ReceiptProduct
	s.db.Find(&products)
	assert.Equal(s.T(), products[0].ProductID, products[1].ProductID, "Both receipt products should reference the same product")
}

func (s *ReceiptProcessorTestSuite) TestProcess_DatabaseError() {
	// Create a test receipt
	receipt := createTestReceipt("receipt-010", "Seller7", "555555", "ShopH", "Address8")

	// Close the database to simulate an error
	sqlDB, _ := s.db.DB()
	err := sqlDB.Close()
	require.NoError(s.T(), err)

	// Try to process the receipt, which should now fail
	err = s.processor.Process(s.db, receipt, make(map[string]uint), make(map[string]uint))
	assert.Error(s.T(), err, "Should return an error when database is closed")
}

// Helper function to create test receipts
func createTestReceipt(externalID, sellerName, inn, shopName, address string) *models.Receipt {
	return &models.Receipt{
		ExternalIdentifier: externalID,
		DateTime:           time.Now(),
		Sum:                2000,
		SellerPlace: &models.SellerPlace{
			Name:    shopName,
			Address: address,
			Seller: models.Seller{
				Name: sellerName,
				Inn:  inn,
			},
		},
		ReceiptProducts: []models.ReceiptProduct{
			{
				Price:    1000,
				Quantity: 2.0,
				Sum:      2000,
				Product:  models.Product{Name: "Product1"},
			},
		},
	}
}
