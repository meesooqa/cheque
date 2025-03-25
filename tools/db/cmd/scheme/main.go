package main

import (
	"context"
	"log"

	"github.com/meesooqa/cheque/db/db_provider"
	"github.com/meesooqa/cheque/db/models"
)

func main() {
	dbProvider := db_provider.NewDefaultDBProvider()
	db, err := dbProvider.GetDB(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
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
