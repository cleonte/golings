package tools

import (
	"github.com/cleonte/golings/exercises/layout/layout2/lib"
)

// Describe says what the tools are working with.
func Describe() string {
	return "the token is " + lib.Token()
}
