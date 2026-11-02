# The maps package

The `maps` package (Go 1.21) has the helpers for maps: `maps.Clone` makes a separate copy, `maps.Copy(dst, src)` adds the entries of one map to another, `maps.Equal` compares the contents, and `maps.DeleteFunc` removes the entries a function picks. `maps.Keys` and `maps.Values` give iterators (Go 1.23), in no particular order because Go randomises map order: `slices.Sorted(maps.Keys(m))` gives the keys in order. A map is a reference: `b := a` is the same map, so clone it when you need your own.

* [Package maps](https://pkg.go.dev/maps)
* [Go maps in action](https://go.dev/blog/maps)
* [Range over function types](https://go.dev/blog/range-functions)
