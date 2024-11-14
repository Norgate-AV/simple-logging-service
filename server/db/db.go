package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDatabase() (*sql.DB, error) {
	var err error

	if db == nil {
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	return db, err
}

func Initialize() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	log.Println("Successfully connected to database")

	if err := createTables(db); err != nil {
		return err
	}

	return nil
}

func createTables(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		room TEXT NOT NULL,
		level TEXT NOT NULL,
		message TEXT NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(statement); err != nil {
		return err
	}

	return nil
}
