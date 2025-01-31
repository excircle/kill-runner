package questions

import (
	"fmt"
	"log"
)

func highlight(text string, color string) string {
	var newtext string
	if color == "red" {
		newtext = fmt.Sprintf("%s%s%s", red, text, reset)
	}
	if color == "green" {
		newtext = fmt.Sprintf("%s%s%s", green, text, reset)
	}
	if len(newtext) == 0 {
		log.Fatalf("Invalid color '%s' provided to highlight() function", color)
	}
	return newtext
}
