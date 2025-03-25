package main

import (
	"context"
	"log"
	"os"

	"receipt-002/common/config"
	"receipt-002/db/db_provider"
	"receipt-002/db/db_types"
	"receipt-002/import/services"
)

func main() {
	err := runImport(
		db_provider.NewDefaultDBProvider(),
		config.NewDefaultConfigProvider(),
		services.NewImportService(),
		os.ReadFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data imported")
}

// runImport imports data from a file into a database
func runImport(
	dbProvider db_types.DBProvider,
	configProvider config.ConfigProvider,
	importService services.ImporterService,
	readFile func(string) ([]byte, error),
) error {
	db, err := dbProvider.GetDB(context.TODO())
	if err != nil {
		return err
	}

	conf, err := configProvider.GetConf()
	if err != nil {
		return err
	}

	// TODO flags
	filename := conf.System.DataPath + "/extract.json"
	data, err := readFile(filename)
	if err != nil {
		return err
	}

	err = importService.SaveReceipt(db, data)
	if err != nil {
		return err
	}

	return nil
}
