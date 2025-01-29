package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// ValidateDB checks if the SQLite database exists, and creates it with the necessary schema if not.
func ValidateDB(dbPath string) (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Database %s does not exist. Creating...", dbPath)

		// Open a connection
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create database: %v", err)
		}

		// Create the `users` table
		createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
		`
		_, err = db.Exec(createTableQuery)
		if err != nil {
			return nil, fmt.Errorf("failed to create users table: %v", err)
		}

		log.Printf("Database %s initialized successfully.", dbPath)
		return db, nil
	}

	// If database exists, open a connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open existing database: %v", err)
	}

	log.Printf("Database %s exists. Connection established.", dbPath)
	return db, nil
}
