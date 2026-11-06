# Fuzzing

A normal test checks the answers you thought of. A *fuzz test* states a property that must hold for any input ("the result is valid text", "decoding what you encoded gives the original", "it never panics"), and the fuzzer makes up inputs to break it. Write `func FuzzXxx(f *testing.F)`, add seeds with `f.Add`, and give `f.Fuzz` a function whose parameters (after `t`) have the types of the seeds.

`go test` alone only runs the seeds, as ordinary tests, so a fuzz test is cheap to keep in the suite. `go test -fuzz FuzzXxx -fuzztime 30s` runs the fuzzer: it uses coverage to steer toward new code paths, and when it finds an input that fails it saves it in `testdata/fuzz/FuzzXxx/` as a new seed. Commit that file: it is a regression test from then on.

Good properties: a round trip (`Decode(Encode(x)) == x`), an invariant (a clamped number stays in range), "no panic" for anything that parses outside data, and agreement with a slow but obviously correct version. The exercises run on their seeds only, so they are fast and the same on every run.

* [Go Fuzzing](https://go.dev/doc/security/fuzz/)
* [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
* [Package testing: F](https://pkg.go.dev/testing#F)
