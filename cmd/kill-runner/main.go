package main

import (
	"log"

	"github.com/excircle/kill-runner/pkg/utils"
)

func main() {
	// Define paths for database and config file
	dbPath := "killdb.sqlite"
	configPath := "kill.config"

	// Validate and initialize the SQLite database
	db, err := utils.ValidateDB(dbPath)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Validate and initialize the config file
	err = utils.ValidateConfig(configPath)
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	log.Println("kill-runner is ready to use.")
}
