# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This is a fork of [mauricioabreu/golings](https://github.com/mauricioabreu/golings). The entry for 1.0.0 describes what changed since upstream's last commit (March 2025, Go 1.19).

## [Unreleased]

### Added

- A `context` topic, 6 exercises: stopping work on cancel, `WithTimeout`, typed keys for `WithValue`, `WithCancelCause`, `WithoutCancel` and `AfterFunc`.
- A `time` topic, 6 exercises, before concurrency: `Duration` arithmetic, format layouts, `ParseInLocation`, `Truncate` in a zone, `Equal` against `==`, and a debounce with `Timer.Reset`.

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

[Unreleased]: https://github.com/cleonte/golings/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/cleonte/golings/releases/tag/v1.0.0
[0.9.0 - 0.18.1]: https://github.com/cleonte/golings/releases/tag/v0.18.1
