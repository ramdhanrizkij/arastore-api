package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	appmigrations "github.com/ramdhanrizkij/arastore-api/migrations"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet("migrate", flag.ContinueOnError)
	action := flags.String("action", "up", "migration action: up, down, status, or redo")
	databaseURL := flags.String("database", os.Getenv("DATABASE_URL"), "PostgreSQL connection URL")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *databaseURL == "" {
		return errors.New("database URL is required; pass -database or set DATABASE_URL")
	}

	db, err := sql.Open("pgx", *databaseURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	goose.SetBaseFS(appmigrations.EmbedMigrations())

	switch *action {
	case "up":
		err = goose.Up(db, ".")
	case "down":
		err = goose.Down(db, ".")
	case "redo":
		err = goose.Redo(db, ".")
	case "status":
		err = goose.Status(db, ".")
	case "create":
		return errors.New("use 'goose create' command directly to create migrations")
	default:
		return fmt.Errorf("unsupported migration action %q", *action)
	}

	if err != nil {
		return err
	}

	fmt.Printf("migration action %s completed successfully\n", *action)
	return nil
}
