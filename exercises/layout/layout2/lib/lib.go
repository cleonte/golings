// Package lib is the public face of a library whose details are in internal/.
package lib

import "github.com/cleonte/golings/exercises/layout/layout2/lib/internal/secret"

// Token is what the rest of the program may know about the secret.
func Token() string {
	return secret.Value[:4] + "..."
}
