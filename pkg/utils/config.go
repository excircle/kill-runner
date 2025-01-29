package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the configuration file.
type Config struct {
	KillRunner struct {
		Config struct {
			Logging bool `yaml:"logging"`
		} `yaml:"config"`
		User struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"user"`
	} `yaml:"kill-runner"`
}

// DefaultConfig returns a default configuration.
func DefaultConfig() Config {
	return Config{
		KillRunner: struct {
			Config struct {
				Logging bool `yaml:"logging"`
			} `yaml:"config"`
			User struct {
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"user"`
		}{
			Config: struct {
				Logging bool `yaml:"logging"`
			}{Logging: true},
			User: struct {
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			}{Username: "", Password: ""},
		},
	}
}

// ValidateConfig checks if the configuration file exists, and if not, creates it with default values.
func ValidateConfig(configPath string) error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("Configuration file %s does not exist. Creating...", configPath)

		// Generate default config
		defaultConfig := DefaultConfig()
		configData, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			return fmt.Errorf("failed to marshal default config: %v", err)
		}

		// Write to file
		err = os.WriteFile(configPath, configData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write default config: %v", err)
		}

		log.Printf("Configuration file %s initialized successfully.", configPath)
	} else {
		log.Printf("Configuration file %s exists. Skipping initialization.", configPath)
	}
	return nil
}
