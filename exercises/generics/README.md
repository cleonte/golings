# Generics

A function or a type can take type parameters: `func Sum[T Number](values []T) T`, `type Stack[T any] struct{...}`. The constraint after the name says which types are allowed: `any` is every type, `comparable` supports `==` (and map keys), `cmp.Ordered` supports `<` and `>`, and an interface such as `~int | ~float64` lists types (the `~` also accepts types built on them, like `type Celsius float64`). The compiler infers the type arguments from the call most of the time: `Sum([]int{1, 2})`. A generic type alias (Go 1.24) can have type parameters too, a type may refer to itself in its own constraint (Go 1.26), and from Go 1.27 a method can have type parameters of its own (see the `go127` topic).

* [An Introduction To Generics](https://go.dev/blog/intro-generics)
* [When To Use Generics](https://go.dev/blog/when-generics)
* [Package cmp](https://pkg.go.dev/cmp)
* [The Go specification: type parameter declarations](https://go.dev/ref/spec#Type_parameter_declarations)
