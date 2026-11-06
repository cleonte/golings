# Unsafe

The `unsafe` package steps around the type system of Go: it tells you how a value sits in memory and lets you treat that memory as something else. It needs no C compiler and no cgo, and the compiler still checks the code, but you take over the promises it normally keeps: a wrong use can read bad memory or break the garbage collector, and the compatibility promise of Go 1 does not cover it.

`unsafe.Sizeof`, `Alignof` and `Offsetof` show the layout of a struct, with the padding the compiler adds so each field starts at a multiple of its alignment. `unsafe.String`, `unsafe.StringData`, `unsafe.Slice` and `unsafe.SliceData` convert between strings, slices and pointers without a copy. `unsafe.Pointer` converts between pointer types, and `unsafe.Add` moves a pointer by a number of bytes.

Reach for it only when a measurement says you must. The `go vet` check `unsafeptr` and the `-d=checkptr` mode of the race detector find many mistakes.

* [Package unsafe](https://pkg.go.dev/unsafe)
* [Go 101: unsafe](https://go101.org/article/unsafe.html)
* [Go Wiki: Compiler And Runtime Optimizations](https://go.dev/wiki/CompilerOptimizations)
