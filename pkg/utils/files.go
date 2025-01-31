package utils

import (
	"bufio"
	"os"
	"strings"
)

// Check if file contains "exp" string
func FileContainsString(fileName string, exp string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trim whitespace/newlines
		if line == exp {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		return false
	}

	return false
}
