# Random numbers

`math/rand/v2` (Go 1.22) is the package for random numbers that don't need to be secret. Its top-level functions (`rand.IntN(n)`, `rand.Float64()`, `rand.Shuffle`, `rand.N(d)` for any integer type including `time.Duration`) use a generator that is seeded at random for you. When a sequence has to be reproducible (a game level, a test), make a generator of your own from a seed: `rand.New(rand.NewPCG(seed1, seed2))`; the same seed gives the same sequence every time.

Nothing from `math/rand` is for secrets: its numbers can be predicted. Passwords, session tokens and keys come from `crypto/rand`, which uses the operating system's source; `rand.Text()` (Go 1.24) returns a random string of 26 base32 characters, ready to use.

* [Package math/rand/v2](https://pkg.go.dev/math/rand/v2)
* [Package crypto/rand](https://pkg.go.dev/crypto/rand)
* [Evolving the Go Standard Library with math/rand/v2](https://go.dev/blog/randv2)
