# Project layout: several packages

Real programs are split into packages, one folder each. The path in an import is the folder: with `module github.com/cleonte/golings` in `go.mod`, the folder `exercises/layout/layout1/greet` is imported as `"github.com/cleonte/golings/exercises/layout/layout1/greet"`. The name you use in code is the one in the `package` clause of the files in that folder, which is usually the name of the folder (`v2` folders are the common exception).

Three rules shape a layout. Only capitalized names are visible to other packages. A folder called `internal` can only be imported from the folder that holds it. And packages cannot import each other in a circle: if two need each other, one of them probably needs less than it imports, or the shared part belongs in a third package.

In these exercises the folder of the exercise is the program (`main.go`, or the test), and the folders inside it are its packages. The `module` line of your `go.mod` must be `github.com/cleonte/golings`, which `golings init` writes (`golings doctor` checks it). A folder made by an older golings says `golings-exercises` there: change that line.

* [How to write Go code](https://go.dev/doc/code)
* [Package names](https://go.dev/blog/package-names)
* [Organizing a Go module](https://go.dev/doc/modules/layout)
* [Go 1.4: internal packages](https://go.dev/doc/go1.4#internalpackages)
