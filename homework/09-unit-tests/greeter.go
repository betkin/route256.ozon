package greeter

import (
	"fmt"
	"strings"
)

func Greet(name string, hour int) string {
	if len(name) == 0 {
		return "<name>: empty!"
	}

	var greeting string
	switch {
	case hour >= 0 && hour < 6:
		greeting = "Good night"
	case hour >= 6 && hour < 12:
		greeting = "Good morning"
	case hour >= 12 && hour < 18:
		greeting = "Hello"
	case hour >= 18 && hour < 22:
		greeting = "Good evening"
	case hour >= 22 && hour < 24:
		greeting = "Good night"
	default:
		return "<hour>: range error!"
	}

	trimmedName := strings.Trim(name, " ")
	return fmt.Sprintf("%s %s!", greeting, strings.Title(trimmedName))
}
