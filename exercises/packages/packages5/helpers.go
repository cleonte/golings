// TODO: Every Go file of a folder has to belong to the same package as the others. Change the package clause below to the one of main.go.
package helpers

import "strings"

func shout(text string) string {
	return strings.ToUpper(text) + "!"
}
