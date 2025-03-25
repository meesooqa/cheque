package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/meesooqa/cheque/common/common_log"
	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/db/db_provider"
	"github.com/meesooqa/cheque/import-cat/services"
)

func main() {
	loggerProvider := common_log.NewConsoleLoggerProvider(slog.LevelDebug)
	lgr, cleanup := loggerProvider.GetLogger()
	defer cleanup()
	dbProvider := db_provider.NewDefaultDBProvider()
	configProvider := config.NewDefaultConfigProvider()
	conf, err := configProvider.GetConf()
	if err != nil {
		log.Fatal(err)
	}
	// TODO flags
	//filePath := flag.String("file", "categories.txt", "Path to the categories file")
	//flag.Parse() // *filePath
	filePath := conf.System.DataPath + "/cat/taxonomy-with-ids.ru-RU.txt"

	reader := services.NewGoogleProductTaxonomyReader(lgr, filePath)
	importer := services.NewGoogleProductTaxonomyImporter(lgr, dbProvider)

	err = runCategoriesImport(reader, importer)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Categories imported successfully!")
}

func runCategoriesImport(reader services.CategoriesReader, importer services.CategoriesImporter) error {
	// Parse categories
	items, err := reader.Read()
	if err != nil {
		return err
	}
	// Import categories to the database
	if len(items) > 0 {
		return importer.Save(context.TODO(), items)
	}
	return nil
}
