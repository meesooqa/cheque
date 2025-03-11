package main

import (
	"fmt"

	"github.com/meesooqa/cheque/common/common_db"
)

func cleanup() {
	db := common_db.GetDB()
	var tables []string
	db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Pluck("tablename", &tables)
	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
	}
}
