package models

import (
	"time"

	"gorm.io/gorm"
)

// Operator – operator with unique name
type Operator struct {
	gorm.Model
	Name string `gorm:"index:idx_operator_name,unique,where:deleted_at IS NULL;not null"`
}

// Seller – seller, Name + Inn = unique
type Seller struct {
	gorm.Model
	Name         string        `gorm:"uniqueIndex:idx_seller_name_inn,where:deleted_at IS NULL;not null"`
	Inn          string        `gorm:"uniqueIndex:idx_seller_name_inn;not null"`
	SellerPlaces []SellerPlace `gorm:"constraint:OnDelete:CASCADE;"`
}

// SellerPlace – seller place, combination of SellerId, Name and Address is unique
type SellerPlace struct {
	gorm.Model
	SellerId uint   `gorm:"uniqueIndex:idx_seller_place,where:deleted_at IS NULL;not null;"`
	Name     string `gorm:"uniqueIndex:idx_seller_place;not null"`
	Address  string `gorm:"uniqueIndex:idx_seller_place;not null"`
	Email    string
	Seller   Seller
}

// Category – categories tree
type Category struct {
	gorm.Model
	ParentId *uint      `gorm:"index"`
	Name     string     `gorm:"not null"`
	Products []Product  `gorm:"many2many:product_categories;"`
	Parent   *Category  `gorm:"foreignKey:ParentId"`
	Children []Category `gorm:"foreignKey:ParentId"`
}

// Brand – brand with unique name
type Brand struct {
	gorm.Model
	Name     string    `gorm:"uniqueIndex:idx_brand_name,where:deleted_at IS NULL;not null"`
	Products []Product `gorm:"constraint:OnDelete:RESTRICT;"`
}

// Product – item with unique name
type Product struct {
	gorm.Model
	Name       string     `gorm:"uniqueIndex:idx_product_name,where:deleted_at IS NULL;not null"`
	BrandId    *uint      `gorm:"index"`
	Brand      *Brand     `gorm:"foreignKey:BrandId;constraint:OnDelete:SET NULL;"`
	Categories []Category `gorm:"many2many:product_categories;"`
	Images     []Image    `gorm:"constraint:OnDelete:CASCADE;"`
}

// Image – product photos
type Image struct {
	gorm.Model
	ProductId uint `gorm:"index;not null;uniqueIndex:idx_product_is_main,where:(is_main = true AND deleted_at IS NULL)"`
	IsMain    bool `gorm:"uniqueIndex:idx_product_is_main"`
	Name      string
	URL       string
	Order     int
}

// Receipt – receipt with unique ExternalIdentifier
type Receipt struct {
	gorm.Model
	ExternalIdentifier   string    `gorm:"uniqueIndex:idx_receipt_external_identifier,where:deleted_at IS NULL;not null"`
	DateTime             time.Time `gorm:"not null"`
	FiscalDocumentNumber int64
	FiscalDriveNumber    string
	FiscalSign           int64
	Sum                  int `gorm:"not null"`
	KktReg               string
	BuyerPhoneOrAddress  string
	OperatorId           *uint `gorm:"index"`
	SellerPlaceId        *uint `gorm:"index"`
	Operator             *Operator
	SellerPlace          *SellerPlace
	ReceiptProducts      []ReceiptProduct `gorm:"foreignKey:ReceiptId;constraint:OnDelete:CASCADE;"`
}

// ReceiptProduct receipt item
type ReceiptProduct struct {
	gorm.Model
	ProductId       uint `gorm:"index;not null"`
	ReceiptId       uint `gorm:"index;not null"`
	Price           int
	Quantity        float64
	Sum             int
	ProductCodeData *string `gorm:"type:jsonb"`
	Product         Product
}
