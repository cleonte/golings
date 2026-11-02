# ROADMAP — modernizing golings

Where things stand (latest release: see the GitHub releases page): 129 exercises in 28 topics, each with a verified solution; a full-screen `watch`; `init`, `reset`, `solution`, `doctor`. Phases 1 to 4 are done; the open work is the rest of Phase 5, condensing the data race report in the TUI and the ideas at the end.

Baseline (upstream HEAD, for the record): last commit 2025-03-12, `go 1.19` in go.mod and CI, `actions/checkout@v3` / `setup-go@v3`, deps from 2022 (cobra 1.6, fsnotify 1.6, go-pretty 6.4, ginkgo 2.5, progressbar 3.12, go-toml 2.0.5, fatih/color 1.13). 43 exercises in 13 topics.

Go version policy: the project floor is **Go 1.26** (latest deps such as x/sys, x/term, x/text require it; raised from the original 1.24 plan). Exercises needing a newer Go than the floor are tagged with a minimum Go version (`go_version` in `info.toml`). CI tests a matrix: the floor and the latest stable.

## Phase 1 — Housekeeping (done)
- [x] Fork, `upstream` remote, branch per phase.
- [x] `go.mod` at `go 1.26`, every dependency updated.
- [x] CI: `actions/checkout`, `actions/setup-go`, version matrix (1.26, stable), cache on, `fixtures/error1` kept out of build and lint.
- [x] `go fix`, `gofmt`, `any` instead of `interface{}`, no deprecated APIs.
- [x] All existing exercises verified: unsolved fails, solved passes.
- [x] README: install, fork notice, supported Go versions. Module renamed to `github.com/cleonte/golings`.
- [x] GoReleaser config (v2) and a release workflow on `v*` tags; binaries for linux, darwin and windows on amd64 and arm64.

## Phase 2 — Tooling and learner experience (done)
- [x] `golings doctor`: Go version against the floor and per-exercise minimums, workspace, race detector prerequisites.
- [x] Per-exercise `go_version` in `info.toml`; the runner skips such exercises when the local Go is too old. Everything now needs Go 1.26 (the floor), so it is for exercises that need a later release, 1.27 and up.
- [x] Per-exercise `race = true` in `info.toml`: only the exercises that depend on the race detector (`concurrent2`, `concurrent10`) are built with `-race`, and they are skipped, like a `go_version` skip, where it is not available (no cgo or no C compiler; macOS needs neither). Every other test exercise runs anywhere Go runs.
- [x] `golings init` (create the exercises folder from the copy embedded in the binary) and `golings reset <exercise>`.
- [x] `golings solution <exercise>`: hidden while the exercise is pending, `--reveal` to look anyway.
- [x] Clearer failure output: compile errors with the source line and caret, test failures and condensed panics.
- [x] Progress by topic; completion-state edge cases fixed.

## Phase 2.5 — Terminal UI (rustlings-style `watch`)
- [x] Runner: build first, run the program under a timeout, `RunContext(ctx)` for cancellation, `ErrTimeout`.
- [x] Watcher: its own package, debounced, sees atomic saves; the goroutine leak in `watch` is fixed.
- [x] Full-screen view (Bubble Tea v2, Lip Gloss v2): progress, current topic, result pane, re-run on save, `q` to quit.
- [x] Keys: hint (`h`), solution (`s`, asks first), reset (`r`, asks first), scrolling, `esc`.
- [x] `--plain` line mode, and automatic fallback when input or output is not a terminal.
- [x] Key: `n` to move on to the next exercise (takes out the marker once the exercise works).
- [x] Key: `c` to check all (runs every exercise four at a time, lists which work, which do not and which were skipped, and flags finished exercises that broke).
- [x] Key: `l` lists every exercise (Bubbles list, with a filter by name or topic) and opens the one picked; `n` goes back to the first pending one.
- [ ] Condense the data race report like panics are condensed (the `WARNING: DATA RACE` block is long and has absolute paths).

## Phase 3 — Missing core topics (done)
Added in learning order, each with hints, a `// TODO:` at the fix site and verified solutions.
- [x] strings (6): runes vs bytes, `strings.Builder`, `strings.Fields`, `SplitSeq`/`Lines` (1.24).
- [x] pointers (5): address and dereference, pointer vs value parameters, nil pointers, `new`.
- [x] methods (5): value vs pointer receivers, method sets.
- [x] interfaces (6): implicit satisfaction, `any`, type assertions and switches, the nil-interface gotcha, embedding.
- [x] errors (5): sentinel errors, wrapping with `%w`, `errors.Is/As/Join`, custom error types.
- [x] closures (5): captured variables, per-iteration loop variables (Go 1.22).
- [x] defer, panic and recover (6): ordering, argument evaluation time, `recover`, named results.
- [x] packages (4): exported names, `init`, initialization order, import aliases.
- [x] testing (5): subtests, `t.Helper`, `t.TempDir`/`t.Cleanup`, examples, `b.Loop` (1.24).
- [ ] Modules, `internal` packages and multi-package layouts. These need an exercise that spans several files with its own `go.mod`; the tool, `reset` and the embedded copy assume one file per exercise. A feature of its own (see Ideas).

## Phase 4 — Modern Go features (done)
Already covered along the way: per-iteration loop variables (closures4), `strings.SplitSeq` and `strings.Lines` (strings6), `b.Loop` (testing5), `errors.Join` (errors5), `slices.Reverse` (strings3).
- [x] Go 1.21: `min`/`max`/`clear` builtins (builtins, 3); the `slices` package: sorting, `BinarySearch`, `Contains`/`Index`, `Insert`/`Delete`/`Compact`, `Clone` (slices_package, 5); the `maps` package: `Keys` with `slices.Sorted`, `Clone`/`Copy`/`Equal`, `DeleteFunc` (maps_package, 3).
- [x] Go 1.22: `for i := range 10` (range4).
- [x] Go 1.23: iterators (iterators, 5): writing an `iter.Seq`, ranging over one, `iter.Seq2`, stopping when `yield` returns false, composing with `slices.Values`, `maps.Keys` and `slices.Collect`.
- [x] Go 1.24: generic type aliases (generics6); the `omitzero` JSON tag (json2, with json1 on struct tags).
- [x] Go 1.25: `sync.WaitGroup.Go` (concurrent11); `testing/synctest` (concurrent12).
- [x] Go 1.26: `new(expr)` (pointers6); `errors.AsType` (errors6); self-referential generic constraints (generics7).
- [x] `log/slog` structured logging basics (slog, 4): key/value pairs, the JSON handler and levels, `With` and groups, `ReplaceAttr`.
- [x] Expand **generics** beyond 2 exercises (now 7): `cmp.Ordered` and `~`, a generic struct, `comparable`, generic type aliases, self-referential constraints. (The "methods can't have type parameters" rule is left out on purpose: generic methods arrive in Go 1.27, so it would not hold on the newest Go.)

## Phase 5 — Concurrency and pitfalls
- [x] Concurrency from 3 to 10 exercises: goroutines, `WaitGroup`, `Mutex`, unbuffered and buffered channels and deadlocks, `select` with a timeout, `context` cancellation, worker pool, fan-in, `sync.Once`, and the race detector.
- [ ] More concurrency: `context` deadlines (`WithTimeout`), fan-out, an `errgroup`-style pattern with the standard library only, `RWMutex`, `sync/atomic`.
- [ ] "Pitfalls" track: slice aliasing and append capacity, map iteration order, shadowing with `:=`, defer in loops, closing channels twice, copying a mutex. (The nil-interface gotcha, closure capture and shadowing are already covered in interfaces5, closures3 and closures4.)
- [x] Solutions tree: a verified solution for every exercise, used by CI to prove each exercise is solvable.

## Phase 6 — Quality and release
- [x] CI check that every unsolved exercise fails and every solution passes (`TestSolutions`, part of the normal test run; a new exercise without a solution fails the build).
- [x] Contributor docs: how to add an exercise, its solution and its `TODO` (CONTRIBUTING.md).
- [x] Releases with binaries on every `v*` tag.
- [ ] Decide: open PRs back to upstream (small, per-topic) or keep the fork independent. Upstream is dormant, so the fork is maintained on its own for now.

## Ideas, not scheduled
- Exercises that span several files (modules, `internal`, multi-package), which also needs `init`, `reset` and the embedded copy to handle a folder instead of a file.
- More topics from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) (MIT licensed, used as a topic list only): reading files with `io/fs`, `encoding/json`, `net/http` with `httptest`, `text/template`, `time` with `testing/synctest`, `reflect`.
- Fuzzing as a testing exercise.
- A web version, as the upstream README mentions.

## Decided
- Module path: `github.com/cleonte/golings`.
- Where solutions live: a top-level `solutions/` tree that mirrors `exercises/`, embedded in the binary separately so `golings init` never writes it, shown only by `golings solution`.
- Exercise format: one file per exercise (`main.go` for compile mode, `main_test.go` for test mode), a manifest entry in `info.toml` in learning order, a `// TODO:` at the spot to fix.
