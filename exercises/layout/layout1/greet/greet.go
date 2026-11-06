// Package greet is a package of its own, in a folder inside the exercise.
package greet

// TODO: Only names that start with a capital letter can be used from another package, and main.go calls greet.Hello. Rename the function.
func hello(name string) string {
	return "Hello, " + name + "!"
}
