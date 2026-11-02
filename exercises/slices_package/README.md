# The slices package

The `slices` package (Go 1.21) holds the loops you would otherwise write yourself: `slices.Sort` and `slices.SortFunc` (with `cmp.Compare` and `cmp.Or`), `BinarySearch` for a sorted slice, `Contains`, `Index` and `IndexFunc`, `Insert`, `Delete` and `Compact`, and `Clone`. Most of them work in place on the slice you give them, and the ones that change its length return the new slice, so keep the result. A slice shares its array, so sorting a slice you were given also sorts the caller's: clone it first when that matters.

* [Package slices](https://pkg.go.dev/slices)
* [Package cmp](https://pkg.go.dev/cmp)
* [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
* [Robust generic functions on slices](https://go.dev/blog/generic-slice-functions)
