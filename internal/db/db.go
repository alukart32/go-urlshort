package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db connection pool
var db *sql.DB

// GetDB returns a new sql.DB if it was not installed earlier,
// otherwise it returns an existing instance. An error occurs
// if the database connection was initialized.
func GetDB(cfg string) (*sql.DB, error) {
	if db == nil {
		return initDB("postgres", cfg)
	}
	return db, nil
}

func initDB(driver, cfg string) (*sql.DB, error) {
	log.Printf("init the new db connection pool...")
	db, err := sql.Open(driver, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s database open error %w", driver, err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Printf("db connection pool was created...")
	return db, nil
}
