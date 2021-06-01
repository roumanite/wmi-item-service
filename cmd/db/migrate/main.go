package main

import (
	"database/sql"
	"fmt"
	"flag"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate"
	"wmi-item-service/config"
	_ "github.com/golang-migrate/migrate/source/file"
	"strconv"
)

func main() {
	directionPtr := flag.String("direction", "up", "up or down")
	schemaPtr := flag.String("schema", "", "schema to be migrated")
	stepPtr := flag.String("step", "", "steps of migration")
	flag.Parse()

	direction := *directionPtr
	schema := *schemaPtr
	steps := 0
	if len(*stepPtr) > 0 {
		stepsInt, err := strconv.Atoi(*stepPtr)
		if err != nil {
			fmt.Println("Must provide step argument as number")
			return
		}
		steps = stepsInt
	}

	if len(schema) == 0 {
		fmt.Println("Must provide db schema")
		return
	}

	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Port,
	)
	if *schemaPtr != "public" {
		dsn += " search_path=" + *schemaPtr
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		cfg.Database.DbName,
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(*stepPtr) == 0 {
		if direction == "down" {
			err = m.Down()
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = m.Up()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *directionPtr == "down" {
		steps = steps * -1
	}

	m.Steps(steps)
}
