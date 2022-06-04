package greeter

import (
	"fmt"
	"strings"
)

const (
	gn  string = "Good night"
	gm  string = "Good morning"
	ge  string = "Good evening"
	def string = "Hello"
)

// Greet prints greetings for each times of day
func Greet(name string, hour int) string {
	if len(name) == 0 {
		return "<name>: empty!"
	}
	trimmedName := strings.Trim(name, " ")
	onlyName := strings.Split(trimmedName, " ")[0]
	if len([]rune(trimmedName)) > 256 {
		return "<name>: too much!"
	}

	var greeting string
	switch {
	case hour >= 0 && hour < 6:
		greeting = gn
	case hour >= 6 && hour < 12:
		greeting = gm
	case hour >= 12 && hour < 18:
		greeting = def
	case hour >= 18 && hour < 22:
		greeting = ge
	case hour >= 22 && hour < 24:
		greeting = gn
	default:
		return "<hour>: range error!"
	}

	return fmt.Sprintf("%s %s!", greeting, strings.Title(onlyName))
}
