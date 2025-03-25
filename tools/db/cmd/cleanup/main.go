package main

import (
	"context"
	"fmt"
	"log"

	"github.com/meesooqa/cheque/db/db_provider"
)

func main() {
	dbProvider := db_provider.NewDefaultDBProvider()
	db, err := dbProvider.GetDB(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	var tables []string
	db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Pluck("tablename", &tables)
	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
	}
}
