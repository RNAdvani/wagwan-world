package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [up|down]")
		return
	}

	cmd := os.Args[1]
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:postgres@localhost:55432/eventguests?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		fmt.Println("Rolled back last migration successfully")
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
