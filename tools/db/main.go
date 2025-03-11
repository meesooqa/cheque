package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cmdGormMigrate := flag.NewFlagSet("gorm:migrate", flag.ExitOnError)
	cmdCleanup := flag.NewFlagSet("cleanup", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'gorm:migrate' or 'cleanup' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gorm:migrate":
		cmdGormMigrate.Parse(os.Args[2:])
		gormMigrate()
	case "cleanup":
		cmdCleanup.Parse(os.Args[2:])
		cleanup()
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
