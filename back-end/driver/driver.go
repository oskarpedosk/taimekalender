package driver

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "/data/taimekalender.db?_foreign_keys=on"
const tablesPath = "/data/tables.sql"

var db *sql.DB

func ConnectDB() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	db, err = sql.Open("sqlite3", wd+dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = CreateDBTables(db, wd+tablesPath)
	if err != nil {
		log.Fatalf("Error creating database tables: %v", err)
	}

	defer db.Close()

	log.Println("Database connection established")
}

func CreateDBTables(db *sql.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}
