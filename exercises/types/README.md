# Types

`type Celsius float64` defines a new type: it has the representation of `float64`, but it is a different type, so a `Celsius` and a `float64` can't be mixed in an expression without a conversion (`Celsius(x)`). That is what makes defined types useful for units, IDs and anything where mixing two numbers up would be a bug. `type A = B` is different: an alias is just another name for `B`, the same type, which helps move code around but protects nothing.

Go has no `enum` keyword. An enum is a defined type with a block of constants, and `iota` numbers them: it is 0 on the first line of a `const` block and goes up by one on each line, and a line without a value repeats the line above (with `iota` one larger), so it also works for flags and sizes (`1 << (10 * iota)`). Give the type a `String() string` method and `fmt` prints the names. Remember that a `Direction` can hold any number, not only the named ones.

* [A Tour of Go: Type conversions](https://go.dev/tour/basics/13)
* [Go spec: Iota](https://go.dev/ref/spec#Iota)
* [Go spec: Type definitions and aliases](https://go.dev/ref/spec#Type_declarations)
* [Go by Example: Enums](https://gobyexample.com/enums)
