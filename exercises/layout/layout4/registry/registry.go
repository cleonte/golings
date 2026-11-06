// Package registry holds the greetings that plugins add to it.
package registry

import "fmt"

var greetings = map[string]string{}

// Register is called by the plugins, from their init functions.
func Register(language, text string) {
	greetings[language] = text
}

// Greet returns the greeting of a language.
func Greet(language string) string {
	text, ok := greetings[language]
	if !ok {
		return fmt.Sprintf("no greeting for %q", language)
	}
	return text
}
