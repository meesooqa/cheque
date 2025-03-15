package main

import (
	"log"

	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

func main() {
	db := common_db.GetDB()
	err := db.AutoMigrate(
		&models.Operator{},
		&models.Seller{},
		&models.SellerPlace{},
		&models.Category{},
		&models.Brand{},
		&models.Product{},
		&models.Image{},
		&models.Receipt{},
		&models.ReceiptProduct{},
	)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
