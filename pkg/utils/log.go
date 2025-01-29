package utils

import (
	"fmt"
	"log"
	"os"
)

var TempLogBuffer []string
var logFile *os.File
var severityLevel = map[int]string{
	0: "[INFO]",
	1: "[WARNING]",
	2: "[ERROR]",
}

// InitLog sets up logging based on the config file.
func InitLog() error {

	// Check if log file exists, create if not
	if _, err := os.Stat(LoggingFile); os.IsNotExist(err) {
		file, err := os.Create(LoggingFile)
		if err != nil {
			return err
		}
		file.Close()
	}

	// Open the log file in append mode
	file, err := os.OpenFile(LoggingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logFile = file
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("%s Logging initialized at level %d", severityLevel[0], LogLevel)
	return nil
}

// ClearLogBuffer writes the log buffer to the log file.
func ClearTempLogBuffer() {
	for _, message := range TempLogBuffer {
		switch LogLevel {
		case 0: // No logging
			return
		case 1: // Print to stdout
			fmt.Println(severityLevel[0] + " " + message + "\n")
		case 2: // Log to file only
			log.Printf(message)
		case 3: // Log to both stdout and file
			fmt.Println(severityLevel[0] + " " + message)
			log.Printf(message)
		}
	}
}

// LogEvent logs a message based on the log level settings.
func LogEvent(severity int, event string, v ...interface{}) {
	if severity == 0 {
		severity = 0
	}

	message := fmt.Sprintf(event, v...)

	switch LogLevel {
	case 0: // No logging
		return
	case 1: // Print to stdout
		fmt.Println(severityLevel[severity] + " " + message)
	case 2: // Log to file only
		log.Printf("%s %s", severityLevel[severity], message)
	case 3: // Log to both stdout and file
		fmt.Println(severityLevel[severity] + " " + message)
		log.Printf("%s %s", severityLevel[severity], message)
	}
}
