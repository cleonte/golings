# Iterators

Since Go 1.23 you can range over a function. An iterator is a function that is given a `yield` function and calls it once per value: `iter.Seq[V]` is `func(yield func(V) bool)`, and `iter.Seq2[K, V]` gives two values at a time. `for v := range seq` calls the iterator and runs the loop body each time `yield` is called. `yield` returns false when the loop stopped (`break`, `return`): the iterator must stop too. The standard library has iterators in `slices` (`slices.Values`, `slices.All`), `maps` (`maps.Keys`) and `strings`, and `slices.Collect` and `slices.Sorted` turn one back into a slice. Iterators compose: a function can take one and give another.

* [Range over function types](https://go.dev/blog/range-functions)
* [Package iter](https://pkg.go.dev/iter)
* [For statements with range clause](https://go.dev/ref/spec#For_range)
