package main

import (
	"log"

	"github.com/excircle/kill-runner/pkg/utils"
)

func main() {
	// Path to the SQLite database file
	dbPath := "killdb.sqlite"

	// Explicitly check and initialize the database
	db, err := utils.ValidateDB(dbPath)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	log.Println("kill-runner is ready to use.")
}
