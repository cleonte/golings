# Modules: go.mod, require, replace and go.work

A module is a folder tree with a `go.mod` at its root. The file says the module's name (`module example.com/app`, which is the start of every import path inside it), the Go it needs (`go 1.26`), and the other modules it uses (`require example.com/greet v1.2.0`). A dependency is found by its name and version on a module server, unless a `replace` line points the name at a folder on disk instead (`replace example.com/greet => ./greet`), which is how you work on two modules at once. A `go.work` file does the same for a whole set of modules (`use ./app` and `use ./greet`), without touching their `go.mod` files. A module whose major version is 2 or more has `/v2` at the end of its name, in its own `go.mod` and in every import and `require` of it.

In these exercises the folder of the exercise is a module (or a workspace) of its own, and the file to fix is the `go.mod` or `go.work` in it, not a Go file. They run with the network off and no other Go downloaded, so a dependency is always a folder inside the exercise (`greet/`). You write the lines by hand; in a real project `go get` or `go mod edit` would write them for you. A folder from an older `golings init` works as well; `golings doctor` checks its module path.

* [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
* [The go.mod file reference](https://go.dev/doc/modules/gomod-ref)
* [Developing a module with a workspace](https://go.dev/doc/tutorial/workspaces)
* [Go modules: v2 and beyond](https://go.dev/blog/v2-go-modules)
