# ROADMAP — modernizing golings

Where things stand (latest release: see the GitHub releases page): 215 exercises in 43 topics, each with a verified solution; a full-screen `watch`; `init`, `reset`, `solution`, `doctor`. Phases 1 to 5 and 7 are done, and the fork's independence is decided (Phase 6). What is left is Phase 8 (tool features) and the topics in the Ideas list.

Baseline (upstream HEAD, for the record): last commit 2025-03-12, `go 1.19` in go.mod and CI, `actions/checkout@v3` / `setup-go@v3`, deps from 2022 (cobra 1.6, fsnotify 1.6, go-pretty 6.4, ginkgo 2.5, progressbar 3.12, go-toml 2.0.5, fatih/color 1.13). 43 exercises in 13 topics.

Go version policy: the project floor is **Go 1.26** (latest deps such as x/sys, x/term, x/text require it; raised from the original 1.24 plan). Exercises needing a newer Go than the floor are tagged with a minimum Go version (`go_version` in `info.toml`). CI tests a matrix: the floor (1.26), 1.27 and the latest stable, which keeps tracking new releases.

## Phase 1 — Housekeeping (done)
- [x] Fork, `upstream` remote, branch per phase.
- [x] `go.mod` at `go 1.26`, every dependency updated.
- [x] CI: `actions/checkout`, `actions/setup-go`, version matrix (1.26, 1.27, stable), cache on, `fixtures/error1` kept out of build and lint.
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
- [x] Condense the data race report like panics are condensed: the two accesses (kind, goroutine, line and function in the exercise file), the same pair of lines once with a count, anything unplaced left as printed. `f` in the full-screen view and `--verbose` on `run`, `verify` and `watch --plain` show the whole output.

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
- [x] Expand **generics** beyond 2 exercises (now 7): `cmp.Ordered` and `~`, a generic struct, `comparable`, generic type aliases, self-referential constraints. (The old "methods can't have type parameters" rule is not an exercise: generic methods arrived in Go 1.27, so it is wrong there. They are covered in Phase 4b, tagged for 1.27.)

## Phase 4b — Go 1.27
Everything here is tagged `go_version = "1.27"`, so it is skipped on 1.26 (the floor) and runs on 1.27. Language features also need a `//go:build go1.27` line in the exercise file, because the module language version is 1.26 (see CONTRIBUTING). The exercises live in the `go127` topic, except the one about `synctest`, which goes with the other concurrency exercises. Checked against the release notes and the real toolchain.
- [x] Gating: `TestSolutions` skips them on a Go that is too old (the tests pass on 1.26 and 1.27); a manifest test requires `go_version` for any exercise that uses a Go 1.27 feature or API; specs show them skipped on 1.26 and running on 1.27.
- [x] `strings.CutLast`: split a file name and a path on the last separator (go127_1).
- [x] `uuid`: generate and parse (go127_2).
- [x] Struct literal keys that are promoted fields (go127_3).
- [x] Generic methods: a method with its own type parameter, `rand.Rand.N` as the reference (go127_4).
- [x] The limit of generic methods: they can't satisfy interface methods (go127_5).
- [x] `encoding/json/v2`: `Marshal` and `Unmarshal`, and where v2 differs from v1 (go127_6).
- [x] `synctest.Sleep` (concurrent13).

## Phase 5 — Concurrency and pitfalls (done)
- [x] Concurrency from 3 to 10 exercises (18 in all, with `concurrent11` to `18` from Phase 4 and below): goroutines, `WaitGroup`, `Mutex`, unbuffered and buffered channels and deadlocks, `select` with a timeout, `context` cancellation, worker pool, fan-in, `sync.Once`, and the race detector.
- [x] More concurrency (concurrent14 to 18): `context.WithTimeout`, a limit on how many tasks run at once (a channel as a semaphore), an `errgroup`-style pattern with the standard library only, `RWMutex`, `sync/atomic` (`CompareAndSwap`).
- [x] "Pitfalls" track (pitfalls, 7): `append` into a shared array, `range` copies the element, shadowing with `:=`, writing to a nil map, map iteration order, closing a channel twice, copying a mutex. Not repeated: defer in loops (defer4), the nil-interface gotcha (interfaces5), closure capture (closures3, closures4). `append` forgetting its result is already a compiler error, so it is not an exercise.
- [x] Solutions tree: a verified solution for every exercise, used by CI to prove each exercise is solvable.

## Phase 6 — Quality and release
- [x] CI check that every unsolved exercise fails and every solution passes (`TestSolutions`, part of the normal test run; a new exercise without a solution fails the build).
- [x] Contributor docs: how to add an exercise, its solution and its `TODO` (CONTRIBUTING.md).
- [x] Releases with binaries on every `v*` tag.
- [x] Decide: open PRs back to upstream (small, per-topic) or keep the fork independent. Decided: the fork stays independent (see Decided).

## Phase 7 — More topics after 1.0 (done)
Same format as before: 5 to 6 exercises per topic, each with a hint, a `// TODO:` at the fix site and a verified solution, added to `info.toml` in learning order, one small PR per topic. Everything stays in memory or in `t.TempDir()`, so no network is needed.
- [x] `context` (6): cancel, timeout and deadline, `WithValue` and when not to use it, `WithCancelCause`, `AfterFunc`, `ctx` as the first argument.
- [x] `time` (6): durations, the reference layout of `Format` and `Parse`, `Truncate`, monotonic against wall clock, timers and tickers (the Go 1.23 change), `testing/synctest`.
- [x] `io` and `bufio` (6): `Reader` and `Writer`, `io.Copy`, `TeeReader`, `bufio.Scanner` and its token limit, `strings.NewReader`, `io.Pipe`.
- [x] `os` and `io/fs` (5, topic `fs`): `os.ReadFile`, `os.Root` (1.24), `fstest.MapFS`, `filepath.WalkDir`, `embed`.
- [x] `net/http` (6, topic `http`): a handler tested with `httptest`, pattern routing (`GET /items/{id}`, 1.22), middleware, client timeouts.
- [x] `cmp`, `sort` and `container` (5, topic `sorting`): `cmp.Compare`, `slices.SortFunc` with several keys, `container/heap` and `list`.
- [x] `types` (5): `iota` enums, a `String` method, `iota` expressions, alias against defined type, conversion between defined types.
- [x] `regexp` (5): anchors, groups, `ReplaceAllString`, `FindAllString`, `QuoteMeta`.
- [x] `template` (5): fields, `range`, `Funcs`, variables and `if` in a range, `html/template` against `text/template`.
- [x] `encoding` (5): URL-safe `base64`, the error of `hex`, reading and writing `csv`, and `binary` byte order.
- [x] `rand` (4): `IntN` in `math/rand/v2`, a generator from a seed, `rand.N` for durations, `crypto/rand.Text` for secrets.
- [x] `memory` (4): `unique.Make`, `sync.Pool`, `weak.Pointer`, `runtime.AddCleanup`.

## Phase 8 — Tool features (planned)
Ideas from reading a similar project, [madhank93/golings](https://github.com/madhank93/golings) (another fork of upstream; Apache-2.0). Ideas only: nothing is copied, and every exercise here is written and verified on its own.
- [x] Graded hints, the tool: besides `hint`, optional `nudge` (what the error means) and `stronger` (where to change it) in `info.toml`; `h` in `watch` shows them one at a time, then the full hint; `golings hint --level N`; `hint` in the plain watch steps through them. Exercises without them keep working as now.
- [x] Graded hints, the content (all 182 exercises): `nudge` and `stronger` for every exercise, written from the real failure of the unsolved version. A manifest test requires them.
- [ ] A key in `watch` to open the exercise in `$VISUAL` or `$EDITOR`.
- [ ] Exercises that span several files and packages (modules, `internal`, multi-package layouts). Needs `init`, `reset`, the embedded copy and `TestSolutions` to handle a folder, and an exercise with its own `go.mod`. Design first, as a short note, before any exercise.

## Phase 9 — Closing the gap with other projects
Topics a similar project ([madhank93/golings](https://github.com/madhank93/golings)) has and this one lacks. Same format as Phase 7, one PR per topic, released together.
- [x] `reflect` (5): kind against type, struct tags, setting through a pointer, `DeepEqual`, walking a value by kind.
- [ ] `unsafe` (pure Go, no C compiler): `Sizeof` and alignment, `unsafe.String`, `unsafe.Slice`.
- [ ] Testing with fakes and dependency injection.
- [ ] `net` over loopback, `net.Pipe` first.
- [ ] Profiling and benchmarks, then fuzzing.
- [ ] `cli`: `flag`, `os.Args`, exit codes.
- [ ] Multi-file modules and a capstone (needs the Phase 8 tool design first).

## Ideas, not scheduled
- More topics from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) (MIT licensed, used as a topic list only): `encoding/json` and `reflect` are done.
- Not planned: `cgo` (needs a C compiler, which the supported setups do not all have) and exercises that start other programs (they depend on what is installed).
- Fuzzing as a testing exercise.
- A web version, as the upstream README mentions.

## Decided
- The fork stays independent. No pull requests back to upstream: it has had no commit since March 2025, and the changes (Go 1.26 floor, a new manifest field, the TUI, 105 more exercises) are too large to send in small pieces. Upstream gets a friendly issue saying a maintained fork exists; the README credits Mauricio Abreu and keeps the Apache-2.0 license.
- Module path: `github.com/cleonte/golings`.
- Where solutions live: a top-level `solutions/` tree that mirrors `exercises/`, embedded in the binary separately so `golings init` never writes it, shown only by `golings solution`.
- Exercise format: one file per exercise (`main.go` for compile mode, `main_test.go` for test mode), a manifest entry in `info.toml` in learning order, a `// TODO:` at the spot to fix.
