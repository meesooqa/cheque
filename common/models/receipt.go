package models

import (
	"time"

	"gorm.io/gorm"
)

// Operator – operator with unique name
type Operator struct {
	gorm.Model
	Name string `gorm:"index:idx_operator_name,unique,where:deleted_at IS NULL;not null" json:"name"`
}

// Seller – seller, Name + Inn = unique
type Seller struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_seller_name_inn,where:deleted_at IS NULL;not null" json:"name"`
	Inn  string `gorm:"uniqueIndex:idx_seller_name_inn;not null" json:"inn"`
}

// SellerPlace – seller place, combination of SellerID, Name and Address is unique
type SellerPlace struct {
	gorm.Model
	SellerID uint   `gorm:"uniqueIndex:idx_seller_place,where:deleted_at IS NULL;not null" json:"seller_id"`
	Name     string `gorm:"uniqueIndex:idx_seller_place;not null" json:"name"`
	Address  string `gorm:"uniqueIndex:idx_seller_place;not null" json:"address"`
	Email    string `json:"email"`
	Seller   Seller `gorm:"foreignKey:SellerID;constraint:OnDelete:CASCADE;" json:"seller"`
}

// Category – categories tree
type Category struct {
	gorm.Model
	ParentID *uint      `gorm:"index" json:"parent_id"`
	Name     string     `gorm:"not null" json:"name"`
	Products []Product  `gorm:"many2many:product_categories;" json:"products"`
	Parent   *Category  `gorm:"foreignKey:ParentID" json:"parent"`
	Children []Category `gorm:"foreignKey:ParentID" json:"children"`
}

// Brand – brand with unique name
type Brand struct {
	gorm.Model
	Name     string    `gorm:"uniqueIndex:idx_brand_name,where:deleted_at IS NULL;not null" json:"name"`
	Products []Product `gorm:"constraint:OnDelete:RESTRICT;" json:"products"`
}

// Product – item with unique name
type Product struct {
	gorm.Model
	Name       string     `gorm:"uniqueIndex:idx_product_name,where:deleted_at IS NULL;not null" json:"name"`
	BrandID    *uint      `gorm:"index" json:"brand_id"`
	Brand      *Brand     `gorm:"foreignKey:BrandID;constraint:OnDelete:SET NULL;" json:"brand"`
	Categories []Category `gorm:"many2many:product_categories;" json:"categories"`
	Images     []Image    `gorm:"constraint:OnDelete:CASCADE;" json:"images"`
}

// ProductCategory – product categories
type ProductCategory struct {
	gorm.Model
	ProductID  uint `gorm:"uniqueIndex:idx_product_category,where:deleted_at IS NULL;not null" json:"product_id"`
	CategoryID uint `gorm:"uniqueIndex:idx_product_category;not null" json:"category_id"`
}

// Image – product photos
type Image struct {
	gorm.Model
	ProductID uint   `gorm:"index;not null;uniqueIndex:idx_product_is_main,where:(is_main = true AND deleted_at IS NULL)" json:"product_id"`
	IsMain    bool   `gorm:"uniqueIndex:idx_product_is_main" json:"is_main"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Order     int    `json:"order"`
}

// Receipt – receipt with unique ExternalIdentifier
type Receipt struct {
	gorm.Model
	ExternalIdentifier   string           `gorm:"uniqueIndex:idx_receipt_external_identifier,where:deleted_at IS NULL;not null" json:"external_identifier"`
	DateTime             time.Time        `gorm:"not null" json:"date_time"`
	FiscalDocumentNumber int64            `json:"fiscal_document_number"`
	FiscalDriveNumber    string           `json:"fiscal_drive_number"`
	FiscalSign           int64            `json:"fiscal_sign"`
	Sum                  int              `gorm:"not null" json:"sum"`
	KktReg               string           `json:"kkt_reg"`
	BuyerPhoneOrAddress  string           `json:"buyer_phone_or_address"`
	OperatorID           *uint            `gorm:"index" json:"operator_id"`
	SellerPlaceID        *uint            `gorm:"index" json:"seller_place_id"`
	Operator             *Operator        `json:"operator"`
	SellerPlace          *SellerPlace     `json:"seller_place"`
	ReceiptProducts      []ReceiptProduct `gorm:"foreignKey:ReceiptID;constraint:OnDelete:CASCADE;" json:"receipt_products"`
}

// ReceiptProduct receipt item
type ReceiptProduct struct {
	gorm.Model
	ProductID       uint    `gorm:"index;not null" json:"product_id"`
	ReceiptID       uint    `gorm:"index;not null" json:"receipt_id"`
	Price           int     `json:"price"`
	Quantity        float64 `json:"quantity"`
	Sum             int     `json:"sum"`
	ProductCodeData *string `gorm:"type:jsonb" json:"product_code_data"`
	Product         Product `json:"product"`
}
