# Pitfalls

Mistakes that compile, look fine and still go wrong. Each exercise here has a test that shows one of them, and you fix it.

A slice shares its array with the slices made from it, and `append` writes into that array when there is room: copy (`slices.Clone`) when you need a slice of your own. `for _, p := range items` gives a copy of each element, so change `items[i]`. `:=` inside a block declares new variables that hide the outer ones with the same name. A nil map can be read but not written to, a map has no order (sort the keys when the output must be the same every time), closing a closed channel panics (`sync.Once` makes a close happen once), and a struct with a `sync.Mutex` must not be copied: use a pointer receiver. `go vet` finds some of these.

* [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
* [The Go specification: for statements with range clause](https://go.dev/ref/spec#For_range)
* [The Go specification: declarations and scope](https://go.dev/ref/spec#Declarations_and_scope)
* [Go maps in action](https://go.dev/blog/maps)
* [Package sync](https://pkg.go.dev/sync)
* [Command vet](https://pkg.go.dev/cmd/vet)
