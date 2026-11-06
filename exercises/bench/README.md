# Benchmarks and profiling

Measure before you optimize. A benchmark is a function `BenchmarkXxx(b *testing.B)` in a `_test.go` file; its loop is `for b.Loop() { ... }` (Go 1.24), and `go test -bench . -benchmem` prints the time and the allocations of each round. Allocations are often the cost that matters, and `testing.AllocsPerRun` counts them in a plain test, with no clock, so a test can say "this function allocates once" and keep saying it.

Once you know something is slow, a profile tells you where. `go test -cpuprofile cpu.out -memprofile mem.out` writes profiles, `runtime/pprof` writes them from your own code, `net/http/pprof` serves them from a running program, and `go tool pprof` reads them.

The usual first fixes are the ones in these exercises: build strings with a `strings.Builder` (or `strings.Join`) instead of `+=`, and give slices and maps their size with `make` when you know it. Don't do either blindly: measure first, and keep the change if the numbers improve.

* [Package testing: benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)
* [Package runtime/pprof](https://pkg.go.dev/runtime/pprof)
* [Profiling Go Programs](https://go.dev/blog/pprof)
* [Diagnostics](https://go.dev/doc/diagnostics)
