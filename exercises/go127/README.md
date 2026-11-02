# Go 1.27

What is new in Go 1.27, one small exercise for each piece. Every exercise in this topic needs Go 1.27: with an older Go, golings skips it and says so, and the rest of the course is not affected. Run `golings doctor` to see which Go you have.

`strings.CutLast` splits a string around the last occurrence of a separator, the way `strings.Cut` does around the first. The `uuid` package generates and parses UUIDs. Methods can have type parameters of their own now (a generic method), but an interface method can't, and a generic method can't implement one. A struct literal can set a field that comes from an embedded struct by its own name. `encoding/json/v2` is a new JSON package with stricter defaults, and `encoding/json` is now built on it. `synctest.Sleep` sleeps on the fake clock of a test and waits for the other goroutines to settle.

* [Go 1.27 Release Notes](https://go.dev/doc/go1.27)
* [Package strings](https://pkg.go.dev/strings)
* [Package uuid](https://pkg.go.dev/uuid)
* [The Go specification: method declarations](https://go.dev/ref/spec#Method_declarations)
* [Package encoding/json/v2](https://pkg.go.dev/encoding/json/v2)
* [Testing time (and other asynchronicities)](https://go.dev/blog/testing-time)
