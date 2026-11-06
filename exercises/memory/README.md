# Memory

Go's garbage collector frees what nothing refers to, and a few standard packages let you work with that. `unique.Make(v)` returns a handle for a value: equal values give equal handles, and comparing handles is as cheap as comparing pointers, so values that repeat a lot are stored once. `sync.Pool` keeps objects for reuse (typically buffers) so a busy program doesn't allocate the same thing over and over; always `Reset` what you get, don't use it after `Put`, and remember that the collector may empty the pool at any time.

A `weak.Pointer[T]` refers to a value without keeping it alive: its `Value()` returns the pointer until the collector frees the value, and `nil` after. `runtime.AddCleanup(ptr, f, arg)` runs `f(arg)` some time after `ptr` became unreachable; the function must not refer to `ptr`. Use weak pointers and cleanups for caches and as a safety net, never for what has to happen: close files and connections explicitly.

* [Package unique](https://pkg.go.dev/unique)
* [Package sync: Pool](https://pkg.go.dev/sync#Pool)
* [Package weak](https://pkg.go.dev/weak)
* [runtime.AddCleanup](https://pkg.go.dev/runtime#AddCleanup)
* [A Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)
