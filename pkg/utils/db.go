package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// ValidateDB checks if the SQLite database exists, and creates it with the necessary schema if not.
func ValidateDB() (*sql.DB, error) {
	if _, err := os.Stat(DbPath); os.IsNotExist(err) {
		LogEvent(1, fmt.Sprintf("Database %s does not exist. Creating...", DbPath))

		// Open a connection
		db, err := sql.Open("sqlite3", DbPath)
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

		LogEvent(1, "Database %s initialized successfully.", DbPath)
		return db, nil
	}

	// If database exists, open a connection
	db, err := sql.Open("sqlite3", DbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open existing database: %v", err)
	}

	log.Printf("Database %s exists. Connection established.", DbPath)
	return db, nil
}
