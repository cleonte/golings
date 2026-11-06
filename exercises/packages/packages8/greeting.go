package greetings

import "fmt"

var greetings = map[string]string{}

// register adds the greeting of a language. Each language lives in its own
// file and registers itself from an init function.
func register(language, text string) {
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
