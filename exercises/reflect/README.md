# Reflection

The `reflect` package lets a program look at values and types it knows nothing about at compile time: `reflect.TypeOf(v)` describes the type, `reflect.ValueOf(v)` holds the value. This is how `encoding/json`, `fmt` and `text/template` work with any struct you give them.

Two ideas matter most. A type has a name and a *kind* (`int`, `slice`, `struct`, `ptr`...): code that walks a value switches on `Kind()`. And a value can only be changed through a pointer: `reflect.ValueOf(&x).Elem()` is settable, `reflect.ValueOf(x)` is a copy. Struct tags (`json:"id"`) are strings that reflection reads with `Tag.Get`.

Reflection is slower than normal code, gives up the checks of the compiler (a wrong call panics at run time) and is hard to read: use it for libraries that must handle any type, and reach for generics or an interface when you know what you need.

* [Package reflect](https://pkg.go.dev/reflect)
* [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
* [Learn Go with Tests: Reflection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection)
