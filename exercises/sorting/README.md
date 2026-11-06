# Sorting and ordered collections

`slices.Sort` sorts a slice of ordered values, and `slices.SortFunc(s, cmp)` sorts with a comparison function that returns a negative number, zero or a positive number. `cmp.Compare(a, b)` is that comparison for any ordered type, and `cmp.Or(x, y, ...)` returns the first argument that is not zero, which is how you sort by several keys: compare the first key, and the second one only when the first is equal. `slices.Sorted(maps.Keys(m))` and `slices.SortedFunc` collect the keys of a map in order, since a map has none of its own.

In a sorted slice, `slices.BinarySearch` finds a value (or the place where it belongs) in logarithmic time, and `slices.Insert` puts a value there. The standard library has two containers for other shapes. `container/heap` keeps a slice as a heap, so that `heap.Pop` always returns the smallest element according to your `Less` method, which makes a priority queue. `container/list` is a doubly linked list whose elements can be moved to the front or the back in constant time, as a least-recently-used cache needs.

* [Package cmp](https://pkg.go.dev/cmp)
* [Package slices](https://pkg.go.dev/slices)
* [Package container/heap](https://pkg.go.dev/container/heap)
* [Package container/list](https://pkg.go.dev/container/list)
