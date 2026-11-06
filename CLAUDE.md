# CLAUDE.md — golings modernization

## What this project is
A fork of https://github.com/mauricioabreu/golings ("rustlings but for Go"), an interactive CLI that teaches Go through small broken exercises the learner fixes. Upstream is dormant (last commit 2025-03-12, Go 1.19, 2022-era deps). Goal: modernize it for current Go and fill the missing core topics. See ROADMAP.md for the ordered plan.

Upstream is Apache-2.0: keep LICENSE, keep attribution to the original author in README.

## Setup
- Fork upstream on GitHub, clone the fork into this folder (or move these two files into the clone). Add upstream as a remote: `git remote add upstream https://github.com/mauricioabreu/golings`.
- Work on a branch per roadmap phase, e.g. `phase-1-housekeeping`. Never commit straight to main.

## Repo layout
- `golings/` — the CLI (Go module `github.com/cleonte/golings`, the `go.mod` is at the root): `golings.go` (main), `cmd/`, `exercises/` (the manifest, the runner, how failures are shortened), `tui/` (the full-screen `watch`), `ui/`, `watcher/`, `workspace/` (`init` and `reset`), `doctor/`, `fixtures/`.
- `embed.go` (package `golings`, at the root) — embeds `info.toml`, `exercises/` and `solutions/` in the binary.
- `exercises/<topic>/<name>/` — the learner-facing exercises (43 topics, 215 exercises), with a `README.md` per topic.
- `solutions/<topic>/<name>/` — a verified solution for every exercise.
- `info.toml` — the exercise manifest: `name`, `path`, `mode` (`compile` or `test`), `hint`, and optionally `go_version` and `race`. Every exercise must be listed here, in learning order.
- `Makefile` — `make test`, `make watch`, `make run`, `make smoke`. `scripts/` — `smoke.sh` and `release-notes.sh`.
- `.github/workflows/` — `test.yml` (Linux, Go 1.26, 1.27 and stable), `environments.yml` (no race detector, Windows, the `go install` smoke test on Linux and Windows), `macos-smoke.yml` (the same smoke test on macOS, weekly and on demand: not part of the checks of a pull request), `lint.yml` and `release.yml` (GoReleaser on a `v*` tag; the release body is the section of `CHANGELOG.md`).
- Inspect one existing exercise and its `info.toml` entry before writing new ones, and copy that format exactly.

## Commands
- Test: `make test` (runs `go test` in `golings/`, excluding `fixtures/error1`, which fails on purpose). It runs every exercise through the real runner (`TestSolutions`: the unsolved version must fail, the solution must pass), so it takes a minute or two. `-short` skips that and the end-to-end tests.
- Smoke test: `make smoke` installs golings into a temporary folder and runs `init`, `list`, `doctor` and `run` on the first exercise.
- Run the tool: `make run` / `make watch`, or `go run golings/golings.go <command>`.
- Update deps: `go get -u ./... && go mod tidy` (from the root).
- Modernize code: `go fix ./...` (Go 1.26+).
- Check formatting and vet: `gofmt -l .` and `go vet` on the CLI package.

## Rules for exercises
- Each exercise is a small, single-concept, broken program. The learner fixes it until it compiles and, where applicable, passes its checks or prints the expected output.
- Keep the Rustlings style: a short comment at the top says what to learn, an `// I AM NOT DONE`-style marker only if the existing format uses one, and a `hint` in `info.toml`.
- Use only the standard library in exercises.
- Idiomatic modern Go only: `any` not `interface{}`, `slices`/`maps` packages, `for i := range n`, `errors.Is/As/Join`, `log/slog`. No `ioutil`, no `rand.Seed`.
- For every new exercise, also create a working solution (outside the learner path, in a `solutions/` tree or whatever mechanism Phase 5 settles on) and verify the solution compiles and passes, and that the unsolved version fails in the intended way.
- Exercises that need a newer Go than the project floor (1.26) must declare it with `go_version` in `info.toml`, so the tool skips them instead of failing confusingly; one that uses a *language* feature of the newer Go also needs a `//go:build go1.X` line at the top. Only an exercise whose unsolved version does not fail reliably without the race detector gets `race = true`. CONTRIBUTING.md has the details, and a manifest test checks them.
- The unsolved version must fail every time, the same way. Waiting for time in a test: use `testing/synctest` where goroutines only wait on channels, timers and `WaitGroup`s (a goroutine waiting for a `sync.Mutex` does not let the fake clock move).
- Tests must pass on Windows too: no `/tmp`, no relative paths from another drive (make folders in the working directory), forward slashes in the TOML paths, no file permissions.
- Order matters: add new exercises to `info.toml` where they belong in the learning progression, not at the end.

## Working style
- Small commits, one logical change each. Run `make test` before every commit.
- Don't rewrite existing exercises' intent; modernize idioms and fix wording only.
- Don't add dependencies to the CLI without a clear reason; prefer stdlib.
- When unsure about the exercise format or manifest schema, read the code in `golings/exercises/` and an existing example, then follow it.
