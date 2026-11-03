# Built-in functions

Some functions are part of the language and need no import: `len`, `cap`, `append`, `make`, `new`, `copy`, `delete`, `panic`... Go 1.21 added three more. `min(a, b, ...)` and `max(a, b, ...)` give the smallest and the largest of two or more values of one ordered type (numbers or strings). `clear(x)` empties a map, deleting every entry but keeping the map, and sets every element of a slice to its zero value, keeping the length.

* [The Go specification: built-in functions](https://go.dev/ref/spec#Built-in_functions)
* [Min and max](https://go.dev/ref/spec#Min_and_max)
* [Clear](https://go.dev/ref/spec#Clear)
* [Go 1.21 Release Notes](https://go.dev/doc/go1.21#language)
