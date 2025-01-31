package utils

import (
	"os"
)

// Check if file contains "exp" string
func FileContains(fileName string, exp string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer file.Close()

	stat, _ := file.Stat()
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return false
	}

	return string(bs) == exp
}
