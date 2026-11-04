# Contributing to golings

Thank you for the interest on contributing to `golings`

Below you can find some useful information if you want to contribute to the project, be it opening a new issue or adding code.

## Adding an exercise

Two steps are required to add a new exercise to the project: adding the exercise metadata to the `info.toml` and creating the respective exercise in the `exercises` folder.

Add the metadata for your exercise in the correct order. Exercises are run in the order they are defined in the `info.toml` file. If you are unsure about the order, ask help in the pull request.

Here is an example of an exercise being added:

```toml
[[exercises]]
name = "variables1"
path = "exercises/variables/variables1/main.go"
mode = "compile"
hint = "hints are cool"
```

The exercise mode is very important. It tells `golings` how to run the exercise. If you are adding an exercise that expects the user to only make it compilable, use `compile` mode. If it has a test suite and you need the user to actually have the tests passing, use `test`.

If the exercise needs a newer Go than the project floor, add an optional `go_version` (for example `go_version = "1.27"`). `golings` then skips the exercise with a warning when the installed Go is older. Exercises for Go 1.27 live in the `go127` topic (or next to the topic they belong to) with `go_version = "1.27"`. The module language version is the floor (1.26), so an exercise that uses a *language* feature of a newer Go (generic methods, say) also needs a `//go:build go1.27` line at the top of its file, before the header comment, which raises the language version of that file. A manifest test fails when an exercise uses something known to need a newer Go (that build line, or the `uuid`, `encoding/json/v2`, `strings.CutLast` and `synctest.Sleep` APIs) without a `go_version` that is high enough: add new such APIs to that test. `TestSolutions` skips the exercises the Go running it is too old for, so the tests pass on 1.26 and 1.27.

Set `race = true` (test mode only) on an exercise that is only wrong because of a data race, so its unsolved version passes, or fails only now and then, without the race detector. Such an exercise is built with `-race`, which needs cgo and a C compiler on Linux and Windows; where that is missing, `golings` skips it with a message instead of failing. Check by building the unsolved test with and without `-race` and running it a few times: leave the flag off if it fails reliably without `-race`, because every exercise with the flag is one more that learners without a C compiler can't do. `TestSolutions` builds each exercise with `-race` exactly when it has the flag, and a manifest test lists the exercises that have it.

Say what the learner has to do in a `// TODO:` comment at the spot that needs fixing, one sentence, without giving the answer away (the solution does not keep the comment). The header comment of the file names the exercise and says whether to make it compile or make the tests pass.

Every exercise also needs a solution in `solutions/<topic>/<name>/`, using the same file name as the exercise (`main.go` or `main_test.go`) and without the `I AM NOT DONE` marker. `make test` runs each exercise through the real runner and checks that the unsolved version fails and the solution passes, and that no solution is left without an exercise. Solutions are embedded in the binary and only shown by `golings solution`; `golings init` never writes them.

Exercises and `info.toml` are embedded in the binary (see `embed.go`), which is what `golings init` and `golings reset` use, so a new exercise needs no extra registration. Files whose names start with `.` or `_` are not embedded.

## Running the test suite

```sh
cd golings
go test -coverprofile=coverage.out -v $(go list ./... | grep -v fixtures/error1)
```

If you have `make` installed it is easy as running `make test`. It runs every exercise through the real runner, so it takes a minute or two; `go test -short` skips that.

`make smoke` installs golings into a temporary folder and runs the first commands a learner types (`init`, `list`, `doctor`, `run` on the first exercise), checking the exit codes.

Learners have other machines than yours, and CI covers them (`.github/workflows/environments.yml`): Linux without the race detector (`CGO_ENABLED=0`, where the tests that look at that machine run, because `GOLINGS_EXPECT_NO_RACE=1` is set) and Windows, plus the smoke test on Linux and Windows. macOS runs the smoke test every week and on demand (`macos-smoke.yml`), because GitHub's macOS runners are often not available and a pull request must not wait for them. So tests must not assume Linux or macOS: no `/tmp`, no relative path from a folder on another drive (make the folder in the working directory), forward slashes in the paths of a manifest, and no file permissions.

## Issues and pull requests

There are specific templates that will guide you through opening issues or pull requests. Found a problem with an exercise? Open an issue with the Exercise problem template.

Thank you! ❤️

## Regenerating the demo GIFs

The GIFs in the README are recorded from scripts with [vhs](https://github.com/charmbracelet/vhs) (`brew install vhs`, which also installs `ttyd`). Install the golings of your checkout first, then run a tape:

```sh
go install ./golings
vhs demo-tui.tape   # misc/demo-tui.gif: the full-screen `golings watch`
vhs demo.tape       # misc/demo.gif: `golings run`
```

## Releasing

Pushing a tag `vX.Y.Z` runs the release workflow, which builds the archives with GoReleaser and publishes the release. The body of the release is the section of `CHANGELOG.md` for that version (`scripts/release-notes.sh vX.Y.Z` prints it, and fails when there is none, in which case GoReleaser writes the notes from the commits). So before tagging: move what is under "Unreleased" into a section `## [X.Y.Z] - date` and add its link at the bottom.
