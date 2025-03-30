package models

import (
	"time"

	"gorm.io/gorm"
)

// Seller – seller, Name + Inn = unique
type Seller struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_seller_name_inn,where:deleted_at IS NULL;not null"`
	Inn  string `gorm:"uniqueIndex:idx_seller_name_inn;not null"`
}

// SellerPlace – seller place, combination of SellerID, Name and Address is unique
type SellerPlace struct {
	gorm.Model
	SellerID uint   `gorm:"uniqueIndex:idx_seller_place_sna,where:deleted_at IS NULL;not null;"`
	Name     string `gorm:"uniqueIndex:idx_seller_place_sna;not null"`
	Address  string `gorm:"uniqueIndex:idx_seller_place_sna;not null"`
	Email    string
	Seller   Seller `gorm:"foreignKey:SellerID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Category – categories tree
type Category struct {
	gorm.Model
	ParentID *uint  `gorm:"index"`
	Name     string `gorm:"not null"`
	NameEn   string
	Products []Product `gorm:"many2many:product_categories;"`
	Parent   *Category `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:SET NULL;"`
}

// Brand – brand with unique name
type Brand struct {
	gorm.Model
	Name     string    `gorm:"uniqueIndex:idx_brand_name,where:deleted_at IS NULL;not null"`
	Products []Product `gorm:"constraint:OnDelete:RESTRICT;"`
}

// Image – product photos
type Image struct {
	gorm.Model
	ProductID uint `gorm:"index;not null;uniqueIndex:idx_product_is_main,where:(is_main = true AND deleted_at IS NULL)"`
	IsMain    bool `gorm:"uniqueIndex:idx_product_is_main"`
	Name      string
	URL       string
	Order     int
}

// Product – item with unique name
type Product struct {
	gorm.Model
	Name       string     `gorm:"uniqueIndex:idx_product_name,where:deleted_at IS NULL;not null"`
	BrandID    *uint      `gorm:"index"`
	Brand      *Brand     `gorm:"constraint:OnDelete:SET NULL;"`
	Categories []Category `gorm:"many2many:product_categories;"`
	Images     []Image    `gorm:"constraint:OnDelete:CASCADE;"`
}

// ProductCategory stores product category relations
type ProductCategory struct {
	ProductID  uint `gorm:"primarykey"`
	CategoryID uint `gorm:"primarykey"`
}

// Receipt – receipt with unique ExternalIdentifier
type Receipt struct {
	gorm.Model
	ExternalIdentifier   string    `gorm:"uniqueIndex:idx_receipt_external_identifier,where:deleted_at IS NULL;not null"`
	DateTime             time.Time `gorm:"not null"`
	FiscalDocumentNumber string
	FiscalDriveNumber    string
	FiscalSign           string
	Sum                  int `gorm:"not null"`
	KktReg               string
	BuyerPhoneOrAddress  string
	Operator             string
	SellerPlaceID        *uint `gorm:"index"`
	Comment              string
	SellerPlace          *SellerPlace     `gorm:"constraint:OnDelete:SET NULL;"`
	ReceiptProducts      []ReceiptProduct `gorm:"foreignKey:ReceiptID;constraint:OnDelete:CASCADE;"`
}

// ReceiptProduct receipt item
type ReceiptProduct struct {
	gorm.Model
	ProductID       uint `gorm:"index;not null"`
	ReceiptID       uint `gorm:"index;not null"`
	Price           int
	Quantity        float64
	Sum             int
	ProductCodeData *string `gorm:"type:jsonb"`
	Product         Product
	Receipt         Receipt
}
