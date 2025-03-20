package main

import (
	"log"
	"os"

	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/db/db_provider"
	"github.com/meesooqa/cheque/import/services"
)

func main() {
	dbProvider := &db_provider.DefaultDBProvider{}
	db := dbProvider.GetDB(nil)
	configProvider := config.NewDefaultConfigProvider()
	conf, err := configProvider.GetConf()
	if err != nil {
		log.Fatal(err)
	}

	filename := conf.System.DataPath + "/extract.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	service := services.NewImportService()
	err = service.SaveReceipt(db, data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data imported")
}
