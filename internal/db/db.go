package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const DRIVER = "postgres"

// db connection pool
var (
	db   *sql.DB
	once sync.Once
)

// GetDB returns a new sql.DB if it was not installed earlier,
// otherwise it returns an existing instance. An error occurs
// if the database connection was initialized.
func GetDB(cfg string) (_ *sql.DB, err error) {
	once.Do(func() {
		log.Printf("init the new db connection pool...")
		db, err = sql.Open(DRIVER, cfg)
		if err != nil {
			err = fmt.Errorf("%s database open error %w", DRIVER, err)
		}

		if err = db.Ping(); err == nil {
			log.Printf("db connection pool was created...")
		}
	})
	return db, nil
}
