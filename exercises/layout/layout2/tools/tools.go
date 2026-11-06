package tools

import (
	"github.com/cleonte/golings/exercises/layout/layout2/lib/internal/secret"
)

// Describe says what the tools are working with.
// TODO: tools is outside lib/, so Go does not let it import lib/internal/secret. Use what lib offers instead: import lib and call lib.Token().
func Describe() string {
	return "the token starts with " + secret.Value[:4]
}
