# Files and file systems

`os` is how a program reaches the operating system: `os.ReadFile` and `os.WriteFile` for a whole file, `os.Open` and `os.Create` for a stream (close it with `defer f.Close()`), `os.MkdirAll`, `os.ReadDir`. When a call fails, the error says why, and `errors.Is(err, fs.ErrNotExist)` tells "no such file" from every other problem.

The `io/fs` package describes a read-only file system as an interface, `fs.FS`. `os.DirFS(dir)` is one backed by a folder, `embed.FS` is one inside your binary, and `testing/fstest.MapFS` is one in memory, which makes code that takes an `fs.FS` easy to test. `fs.WalkDir` visits a tree (return `fs.SkipDir` to skip a folder), `fs.Glob` matches names, and `fs.Sub` gives a file system rooted at a sub-folder. `filepath` works with paths of the running system: `filepath.Join`, `filepath.WalkDir`, `filepath.Base`.

A path that comes from outside can contain `..` or a symbolic link and lead out of the folder you meant. `os.Root` (Go 1.24) opens a folder and refuses any name that escapes it. In tests, `t.TempDir()` gives a fresh folder that is removed afterwards.

* [Package os](https://pkg.go.dev/os)
* [Package io/fs](https://pkg.go.dev/io/fs)
* [Traversal-resistant file APIs (os.Root)](https://go.dev/blog/osroot)
* [Go by Example: Directories](https://gobyexample.com/directories)
