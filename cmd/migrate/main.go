package main

import (
	"flag"
	"fmt"
	"log"

	"ig4llc.com/internal/infrastructure/db"
)

func main() {
	action := flag.String("action", "status", "up | down | status")
	flag.Parse()

	if err := db.InitDB(); err != nil {
		log.Fatal("DB init failed:", err)
	}
	defer db.CloseDB()

	switch *action {
	case "up":
		fmt.Println("Running migrations...")
		if err := db.RunMigrations(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Migrations complete.")

	case "down":
		fmt.Println("Rolling back last migration...")
		if err := db.RollbackLastMigration(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Rollback complete.")

	case "status":
		fmt.Println("Migration status:")
		if err := db.PrintMigrationStatus(); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println("Unknown action:", *action)
	}
}
