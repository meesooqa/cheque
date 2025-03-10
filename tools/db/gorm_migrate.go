package main

import (
	"log"

	"github.com/meesooqa/cheque/common/models"
)

func gormMigrate() {
	db := getDd()
	err := db.AutoMigrate(
		&models.Operator{},
		&models.Seller{},
		&models.SellerPlace{},
		&models.Category{},
		&models.Brand{},
		&models.Product{},
		&models.ProductCategory{},
		&models.Image{},
		&models.Receipt{},
		&models.ReceiptProduct{},
	)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
