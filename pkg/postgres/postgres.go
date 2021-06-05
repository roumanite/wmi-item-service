package postgres

import (
	"fmt"
	"database/sql"
	"time"
	_ "github.com/golang-migrate/migrate/database/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(host string, username string, password string, dbName string, port int, schema string) (*gorm.DB, error) {
	db, err := connect(host, username, password, dbName, port, schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connect(host string, username string, password string, dbName string, port int, schema string) (*gorm.DB, error) {
	dialect, err := getPostgresDialect(host, username, password, dbName, port, schema)
	if err != nil {
		return nil, err
	}

	dbConn, err := gorm.Open(dialect, &gorm.Config{})
	return dbConn, err
}

func getPostgresDialect(host string, username string, password string, dbName string, port int, schema string) (gorm.Dialector, error) {
	sqlDB, err := Open(host, username, password, dbName, port, schema)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		return nil, err
	}

	var dialect gorm.Dialector
	dialect = postgres.New(postgres.Config{Conn: sqlDB})
	return dialect, nil
}

func Open(host string, username string, password string, dbName string, port int, schema string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, dbName, port)
	if schema != "public" {
		dsn += " search_path=" + schema
	}

	db, err := sql.Open("postgres", dsn)
	return db, err
}
