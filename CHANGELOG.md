# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This is a fork of [mauricioabreu/golings](https://github.com/mauricioabreu/golings). The entry for 1.0.0 describes what changed since upstream's last commit (March 2025, Go 1.19).

## [Unreleased]

### Added

- A `modules` topic, 5 exercises, after `layout`, about the files that make a module, each a folder that is a module of its own, built with the network off: a `go` line newer than the installed Go, a `replace` without its `require`, a `require` without its `replace`, a `/v2` module path, and a `go.work` that does not list a module. 262 exercises in 52 topics in all.
- Support for exercises about modules: an exercise can be a module of its own (`module = true` in `info.toml`), with a `go.mod`, a dependency in a folder of it that a `replace` points to, or a `go.work`. It is built and run inside its folder with the network off, no other Go downloaded and the workspace off, so it works the same everywhere. `watch` also notices saves of `go.mod`, `go.sum` and `go.work`, and `golings init`, `reset` and `solution` handle the files of such an exercise (they are stored with a `.txt` added in the repository and in the binary). A symbolic link in the path of the exercises folder does not confuse the workspace of such an exercise.

## [1.6.0] - 2026-11-07

### Added

- A `capstone` topic, at the end, with 5 exercises: a word counter built a step at a time, each exercise a step in the life of the same program, starting from the finished previous step, with the files of one package (and a package of its own) in the folder of the exercise. `capstone1` adds the missing method of a `Counter` type, `capstone2` reads words from a stream (and the limits of `bufio.Scanner`), `capstone3` defines an interface where it is used and tests with a fake, `capstone4` is the whole program with flags and exit codes, and `capstone5` counts in parallel without sharing. 257 exercises in 51 topics in all.
- `e` in the full-screen `watch` opens the exercise in your editor (`$VISUAL`, or else `$EDITOR`), with the main file first and the other Go files of the exercise after it. The view waits until you close the editor. Without an editor set, the key says how to set one.

## [1.5.0] - 2026-11-06

### Added

- **9 new exercises**, 252 exercises in 50 topics in all, and exercises that are more than one file: an exercise can now have several files and packages of its own, and the first ones are here.
- A `layout` topic, 5 exercises, after `packages`, with packages of their own inside the exercise folder: an unexported name, `internal/`, an import cycle, a blank import that registers a plugin, and a package whose name is not its folder (`v2`). 252 exercises in 50 topics in all.
- **4 exercises of several files** in the `packages` topic: a file with the wrong package clause, a leftover second `main`, `init` order across files (set values in declarations), and a registry that you extend by adding a new file. For the last one `golings reset` removes the Go file you added.

### Changed

- `golings init` writes `module github.com/cleonte/golings` in the `go.mod` of the exercises folder (it was `golings-exercises`), the module path of this repository, so that an exercise with several packages imports its own packages by the same path everywhere. A folder made by an older version keeps working for every exercise except the `layout` ones; `golings doctor` tells you to change the first line of its `go.mod`.
- An exercise is now the whole folder, not one file, so an exercise can have several files in the same package. Nothing changes for the exercises that exist: `path` still names the main file, with the `// I AM NOT DONE` marker. `golings reset` restores every file of the folder, removes the Go files you added to it, and `golings solution` shows every file of a solution. This is the first step of the design in `docs/multi-file-exercises.md`.

## [1.4.0] - 2026-11-06

### Added

- **33 new exercises in 7 topics**, 243 exercises in 49 topics in all, each with the three hint levels, a `// TODO:` and a verified solution. This closes the gap with other Go exercise projects on core library topics.
- A `reflect` topic, 5 exercises, before generics: kind against type name, reading struct tags, setting a field through a pointer, `reflect.DeepEqual`, and walking a value by kind.
- An `unsafe` topic, 4 exercises, after `reflect`: struct padding and field order with `unsafe.Sizeof`, a string over a slice with `unsafe.String`, bytes over a slice of numbers with `unsafe.Slice`, and `unsafe.Add` for pointer arithmetic. 
- A `fakes` topic, 5 exercises, after `testing`: injecting a dependency through an interface, a fake that fails to test a retry, a fake that keeps the contract of the real store, an interface defined where it is used, and `t.Cleanup` to put back a replaced global.
- A `net` topic, 5 exercises, after `http`: answering on a `net.Pipe`, reading a message that arrives in pieces, closing a connection so the client sees the end, recognizing a timeout with `errors.Is`, and a goroutine for each connection. Everything stays on the loopback interface.
- A `bench` topic, 5 exercises, after `fakes`: `strings.Join` against `+=`, preallocating a slice, both checked with `testing.AllocsPerRun`, a CPU profile that has to be stopped, the heap profile by name, and the loop of a benchmark (`b.Loop`).
- A `fuzz` topic, 4 exercises, after `bench`: a property over runes, no panic on bad input, seeds of the right types for `f.Add`, and an escape round trip. The exercises run on their seeds, so they are fast and the same every time.
- A `cli` topic, 5 exercises, after `net`: parsing with a `flag.FlagSet`, the arguments left after the flags, exit codes and what goes to stderr, a repeatable flag with `Func`, and subcommands. The programs take their arguments as a parameter, so a test can call them. 243 exercises in 49 topics in all.

## [1.3.0] - 2026-11-06

### Added

- **28 new exercises in 6 topics**, 210 exercises in 42 topics in all, each with the three hint levels, a `// TODO:` and a verified solution. This finishes the list of new topics planned after 1.0:
  - A `memory` topic, 4 exercises, before generics: `unique.Make` handles, reusing buffers with `sync.Pool`, a cache that holds its values with `weak.Pointer`, and `runtime.AddCleanup`.
  - A `rand` topic, 4 exercises, before generics: the new names in `math/rand/v2`, a generator made from a seed, `rand.N` for durations, and `crypto/rand.Text` for secrets.
  - An `encoding` topic, 5 exercises, before generics: URL-safe `base64`, the error of `hex.DecodeString`, reading `csv` with quoted fields, the `Flush` of a `csv.Writer`, and big-endian against little-endian in `encoding/binary`.
  - A `template` topic, 5 exercises, before generics: reading fields, `range`, registering a function with `Funcs`, an index and `if` inside a range, and `html/template` against `text/template`.
  - A `regexp` topic, 5 exercises, before generics: anchors in a pattern, groups with `FindStringSubmatch`, `ReplaceAllString` against the literal version, `FindAllString`, and `QuoteMeta`.
  - A `types` topic, 5 exercises, after structs: `iota` enums, a `String` method for them, `iota` expressions, an alias against a defined type, and conversion between defined types.

## [1.2.0] - 2026-11-06

### Added

- **Graded hints.** Every exercise now has three hints, from small to full: a `nudge` that says what the failure means, a `stronger` hint that says where to change, and the `hint` with the explanation and a link. In `watch`, `h` shows them one at a time (`h` again for the next one, and after the last it closes); `golings hint <name> --level N` prints one (without the flag you get the full hint, as before); and typing `hint` again in the plain watch steps through them.
- **16 new exercises in 3 topics**, 182 exercises in 36 topics in all, with the hint levels from the start:
  - `fs` (5), before generics: `fs.ErrNotExist`, walking an `fs.FS` (`fstest.MapFS` in the tests), `fs.SkipDir`, `os.Root` against paths that leave a folder, and `fs.Sub`.
  - `http` (6), before generics, all in memory with `httptest`: the order of status and body, `GET /items/{id}` patterns, a middleware that stops the request, a bad query parameter, `MaxBytesReader`, and a client timeout.
  - `sorting` (5), after iterators: sorting by several keys with `cmp.Or`, `BinarySearch` and `Insert`, `SortedFunc` over map keys, a `container/heap` priority queue and `container/list`'s `MoveToFront`.

## [1.1.0] - 2026-11-06

### Added

- **18 new exercises in 3 topics**, 166 exercises in 33 topics in all. Each has a hint with a link, a `// TODO:` at the spot to fix and a verified solution, as before.
- A `context` topic, 6 exercises: stopping work on cancel, `WithTimeout`, typed keys for `WithValue`, `WithCancelCause`, `WithoutCancel` and `AfterFunc`.
- A `time` topic, 6 exercises, before concurrency: `Duration` arithmetic, format layouts, `ParseInLocation`, `Truncate` in a zone, `Equal` against `==`, and a debounce with `Timer.Reset`.
- An `io` topic, 6 exercises, before generics: a `Read` that returns less than is left, writing to an `io.Writer`, the error of a `bufio.Scanner`, `LimitReader`, hashing while copying with `MultiWriter`, and a reader of your own.
- A `NOTICE` file with the attribution to upstream, which the release archives now include next to `LICENSE`.

### Changed

- The test libraries (ginkgo and gomega) are up to date. Nothing in the tool or the exercises changes for learners.
- The README says the project is independent now, not a GitHub fork. The credit to the original author and the link to the original repository stay.

## [1.0.0] - 2026-11-04

### Changed

- The project needs **Go 1.26** or newer (it was Go 1.19), and every dependency is current. CI tests Go 1.26, 1.27 and the latest stable release, on Linux (with and without the race detector) and Windows, and runs a `go install` smoke test on Linux, Windows and (every week) macOS.
- The Go module is `github.com/cleonte/golings`; install it with `go install github.com/cleonte/golings/golings@latest`.
- Only the exercises that depend on the race detector are built with `-race` (`concurrent2` and `concurrent10`); the other test exercises run anywhere Go runs.
- Failures are easier to read: compile errors show the source line with a caret under the column, panics and data race reports are cut down to the lines in your exercise (`f` in the full-screen view and `--verbose` show everything), and a test that runs for more than 30 seconds is stopped with a message.

### Added

- **148 exercises in 30 topics**, up from 43 in 13. The new topics are strings, builtins, the slices package, the maps package, packages, pointers, methods, interfaces, closures, defer (with panic and recover), testing, slog, json, iterators, pitfalls and Go 1.27; errors is new too, and concurrency grew from 3 to 18 exercises and generics from 2 to 7. They cover the Go 1.21 to 1.27 features: `min`, `max` and `clear`, `slices` and `maps`, range over an integer and over a function, `log/slog`, `errors.AsType`, `new(expr)`, `sync.WaitGroup.Go`, `testing/synctest`, generic methods, `encoding/json/v2`, `uuid` and `strings.CutLast`.
- **A verified solution for every exercise** in `solutions/`, and a test that proves the unsolved version fails and the solution passes, so a new exercise without one fails the build. `golings solution` shows it once the exercise is done, or on request.
- A `// TODO:` at the spot to fix in every exercise, a hint with a link, and a README for most topics.
- A **full-screen `golings watch`** (Bubble Tea) that runs the exercise again when you save: keys for the hint (`h`), the solution (`s`), reset (`r`), the next exercise (`n`), a list of all exercises with a filter (`l`), checking all of them (`c`), the whole output of a failure (`f`) and scrolling. `--plain` gives line-by-line output, which is also what you get when input or output is not a terminal.
- New commands: `init` (creates the exercises folder from the copy inside the binary), `reset`, `solution`, `doctor` (checks the Go version, the exercises and what the race detector needs), and progress by topic in `list`.
- Exercises can say which Go they need (`go_version`) or that they need the race detector (`race = true`); on a machine that can't run them they are skipped with a message, never failed, and left out of the progress.
- Release binaries for Linux, macOS and Windows (amd64 and arm64) on every `v*` tag, with a checksums file, built with GoReleaser.
- A dev container, contributor documentation (`CONTRIBUTING.md`) and a demo of the full-screen view.

### Fixed

- `watch` no longer leaks a goroutine, notices editors that save by writing a temporary file and renaming it, and no longer deadlocks at start on Windows.
- Completion and progress are right when an exercise needs a newer Go than the one installed.

## [0.9.0 - 0.18.1] - 2026-11-01 to 2026-11-03

The development releases of this fork, on the way to 1.0.0: everything listed above, added a piece at a time. (Versions 0.1.0 to 0.8.0 are upstream's releases.)

[Unreleased]: https://github.com/cleonte/golings/compare/v1.6.0...HEAD
[1.6.0]: https://github.com/cleonte/golings/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/cleonte/golings/compare/v1.4.0...v1.5.0
[1.4.0]: https://github.com/cleonte/golings/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/cleonte/golings/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/cleonte/golings/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/cleonte/golings/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/cleonte/golings/releases/tag/v1.0.0
[0.9.0 - 0.18.1]: https://github.com/cleonte/golings/releases/tag/v0.18.1
