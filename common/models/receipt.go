package models

import (
	"time"

	"gorm.io/gorm"
)

type Operator struct {
	gorm.Model
	Name string `gorm:"unique;not null"` // operator
}

type Seller struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_seller_name_inn"` // user, if empty: userInn
	Inn  string `gorm:"uniqueIndex:idx_seller_name_inn"` // userInn
}

type SellerPlace struct {
	gorm.Model
	SellerID uint   `gorm:"uniqueIndex:idx_seller_place"` // seller ID
	Name     string `gorm:"uniqueIndex:idx_seller_place"` // retailPlace
	Address  string `gorm:"uniqueIndex:idx_seller_place"` // retailPlaceAddress
	Email    string // sellerAddress
	Seller   Seller
}

type Category struct {
	gorm.Model
	ParentID *uint      `gorm:"index"`    // parent
	Name     string     `gorm:"not null"` // not unique
	Products []Product  `gorm:"many2many:product_categories;"`
	Parent   *Category  `gorm:"foreignKey:ParentID"`
	Children []Category `gorm:"foreignKey:ParentID"`
}

type Brand struct {
	gorm.Model
	Name     string    `gorm:"unique;not null"`
	Products []Product `gorm:"constraint:OnDelete:RESTRICT;"`
}

type Product struct {
	gorm.Model
	Name       string `gorm:"unique;not null"` // items.name
	BrandID    *uint  `gorm:"index"`
	Brand      *Brand
	Categories []Category `gorm:"many2many:product_categories;"`
	Images     []Image    `gorm:"constraint:OnDelete:CASCADE;"`
}

type ProductCategory struct {
	gorm.Model
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}

type Image struct {
	gorm.Model
	ProductID uint `gorm:"index"`
	Name      string
	URL       string
	Order     int
	IsMain    bool
}

type Receipt struct {
	gorm.Model
	ExternalIdentifier   string    `gorm:"unique;not null"` // _id
	DateTime             time.Time `gorm:"not null"`        // dateTime 2017-08-04T10:09:00
	FiscalDocumentNumber int64
	FiscalDriveNumber    string
	FiscalSign           int64
	Sum                  int    `gorm:"not null"` // totalSum
	KktReg               string // kktRegId
	BuyerPhoneOrAddress  string // buyerPhoneOrAddress
	OperatorID           *uint  `gorm:"index"`
	SellerPlaceID        *uint  `gorm:"index"`
	Operator             *Operator
	SellerPlace          *SellerPlace
	ReceiptProducts      []ReceiptProduct `gorm:"foreignKey:ReceiptID;constraint:OnDelete:CASCADE;"`
}

type ReceiptProduct struct {
	gorm.Model
	ProductID       uint    `gorm:"index"`
	ReceiptID       uint    `gorm:"index"`
	Price           int     // price
	Quantity        float64 // quantity
	Sum             int     // sum
	ProductCodeData *string `gorm:"type:jsonb"` // productCodeData
	Product         Product
}
