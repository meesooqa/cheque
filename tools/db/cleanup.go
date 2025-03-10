package main

import "fmt"

func cleanup() {
	db := getDd()
	var tables []string
	db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Pluck("tablename", &tables)
	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
	}
}
