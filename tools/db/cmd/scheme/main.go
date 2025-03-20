package main

import (
	"log"

	"github.com/meesooqa/cheque/common/models"
	"github.com/meesooqa/cheque/db/db_provider"
)

func main() {
	dbProvider := &db_provider.DefaultDBProvider{}
	db := dbProvider.GetDB(nil)
	err := db.AutoMigrate(
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
