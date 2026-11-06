# Design note: exercises that span several files and packages

Status: accepted. The decision is to build step 1 and step 2, use them for new exercises in existing topics and a package-layout topic, and leave the capstone for later (nothing here closes that door). Step 3 waits until someone asks for it.

## Why

Every exercise today is one broken file. That cannot teach three things learners meet in real projects:

- **Several files in one package**: where code is split, how a type in one file is used in another.
- **Several packages**: `internal/`, exported against unexported names, import cycles, a `main` that wires packages together.
- **Modules**: `go.mod`, `require`, `replace`, versions, `go.work`.

It also blocks a **capstone**: a small program (say a word counter or a key-value store) that uses what the course taught, with a test suite for the whole program.

## What the code assumes today

Checked against the code at v1.4.0:

| Place | Assumption |
|---|---|
| `info.toml` | `path` is one file: `exercises/<topic>/<name>/main.go` or `main_test.go` |
| `exercises/runner.go` | builds the file: `go build ./<file>` or `go test -c ./<file>`; test mode runs the binary in the folder of the file |
| `exercises/exercise.go` | the marker `// I AM NOT DONE` is read from, and removed from, that one file |
| `exercises/topic.go` | `SolutionPath()` maps the file to `solutions/<topic>/<name>/<file>` |
| `solutions_test.go` | reads one solution file and runs it in place of the exercise file |
| `workspace.Reset` | restores one file |
| `cmd/solution.go`, `tui` | show one solution file |
| `workspace.Init`, `embed.go` | copy the whole tree, so they already work for any number of files |
| `watcher` | watches every `.go` file below the folder, so it already works |

The key observation: **every exercise already lives in a folder of its own** (`exercises/<topic>/<name>/`), and that folder holds one file. So building the *folder* instead of the *file* gives the same result for all 243 exercises and allows more files for free.

## Proposal, in three steps

Each step is a small PR that can be released alone. Step 1 changes no behavior.

### Step 1: an exercise is a folder; `path` stays the main file

- The runner builds the package in the folder of `path` (`go build ./<dir>`, `go test -c ./<dir>`). For the existing exercises this is the same build.
- `path` keeps its meaning as the **main file**: it holds the `// I AM NOT DONE` marker, it is the file shown in the TUI and the one `n` takes the marker out of. Other files in the folder are plain code with their own `// TODO:` comments.
- No new manifest fields. A multi-file exercise is one whose folder has more than one `.go` file.
- `reset` restores every file of the exercise folder from the embedded copy and removes the Go files the learner added to the folder, so the exercise starts over (other files, such as notes, stay).
- Solutions: `solutions/<topic>/<name>/` holds the files that **replace or add to** the exercise files of the same name; it need not have the main file at all (an exercise solved by adding a new file has only that file). `TestSolutions` copies the exercise folder to a temporary folder, copies the solution files over it, and runs that. That also removes the need for solution files to be complete copies.
- `golings solution` and the `s` key show all solution files, each with its name, not only the main one.
- `TestNoOrphanedSolutions` and the manifest tests check "every solution file has an exercise file or is an addition inside its folder" instead of "same path".
- The Go version checks in the manifest tests read every file of the folder.
- CONTRIBUTING.md and CLAUDE.md change "one file per exercise" to "one folder per exercise, usually one file".

Cost: small. Risk: low, because the existing 243 exercises are the test.

### Step 2: sub-packages in the exercise folder

This is what `internal/`, import cycles and the capstone need. The open problem is the **import path**: code in `exercises/capstone/c1/store/` must be imported by something, and the path depends on the module name, which differs between the repository (`github.com/cleonte/golings`) and a learner's folder (`golings-exercises`, written by `init`).

Options:

1. **Nested `go.mod` per exercise** (`module c1`, imports `c1/store`). Works the same in the repository, in a learner's folder and in an editor. Cost: `go:embed` skips folders that contain a `go.mod`, so the file would have to be stored under another name (`go.mod.txt`) and renamed by `init` and `reset`; and the repository root module would not see these folders (good for lint, bad for editing them in the repository).
2. **One module, fixed name.** `init` already writes `module golings-exercises`. Exercises import `golings-exercises/exercises/capstone/c1/store`. In the repository the path does not resolve, so the exercises cannot be built in place; `TestSolutions` already copies to a temporary folder in step 1, so it would copy into a temporary module named `golings-exercises`. Cost: an editor in the repository shows red imports. They already show errors in unsolved files.
3. **Rename the learner module** to the repository's (`github.com/cleonte/golings`) so the paths are identical. Rejected: a learner's folder would claim our import path, and `go mod tidy` or an IDE could try to download it.

Recommendation: **option 2** for the capstone and sub-package exercises, because it needs no embed workaround and no new files in the exercise folder. The runner needs to know which package to build: the folder of `path` stays the one built (`main` for compile mode, the test package for test mode), and it imports its siblings by the `golings-exercises/...` path. Tests of the sub-packages themselves can come later (`go test -c` builds one package at a time, so the runner would need a `packages = [...]` list).

### Step 3: module exercises

Exercises about `go.mod` itself (`go mod init`, `require`, `replace`, a local dependency, `go.work`) need their own module, so option 1 applies, only here. Proposal: the exercise folder holds `go.mod.txt` (and `go.sum.txt` if needed), and the manifest field `module = true` tells the runner to run `go build`/`go test -c` **inside** the folder (`cmd.Dir`) and `init`/`reset` to write the files under their real names. The learner's task is then often to fix `go.mod`, so the main file for these exercises is `go.mod`, which means the marker in it is `// I AM NOT DONE`, a valid comment there too. The marker regex already allows `//`.

This is the part with the most open questions; it should wait until steps 1 and 2 are in use.

## The capstone

Once step 2 exists, a capstone is a series of exercises on **one small program**, each a step in its life, so a learner never faces a blank page:

1. `capstone1`: a package with a type and its methods (fix the compile errors).
2. `capstone2`: add the parser; a second file in the same package.
3. `capstone3`: an interface and a fake for the store, in a sub-package.
4. `capstone4`: wire it together in `main` with `flag` and exit codes.
5. `capstone5`: make it concurrent without a race.

Each step starts from the solved previous step, so the exercise folder carries the earlier code (finished) and the new broken part. Candidate programs: a word counter for a text stream, an in-memory key-value store with a text protocol over `net.Pipe`. The second reuses the `net`, `cli` and `fakes` topics.

## What stays the same

- `golings run`, `verify`, `list`, `hint`, `doctor`, `watch` work by exercise name.
- The manifest keeps its fields; step 1 adds none, step 3 adds one optional field (`module`).
- Every exercise keeps a verified solution, and `TestSolutions` keeps proving that the unsolved version fails and the solution passes.
- Windows and race-detector behavior are unchanged.

## Open questions

1. Is the capstone worth the maintenance? Decided: not now; see Order of work.
2. Option 2 against option 1 for import paths (the recommendation above).
3. Should `golings solution` for several files print all of them, or ask which one?
4. Should a multi-file exercise be done when the main file loses its marker, or only when every file has no `// TODO:` left? Proposal: the marker in the main file, as now; it is the learner's own statement.

## Order of work

Decided: no capstone for now.

1. Step 1 (an exercise is a folder), with a CHANGELOG line and no visible effect for learners.
2. Two-file exercises in existing topics (for example `packages`) to prove the mechanism.
3. Step 2 (sub-packages, one module with a fixed name) and a package-layout topic of independent exercises.
4. Later, if it is wanted: a one-exercise pilot of the capstone, then the other steps. Step 3 (module exercises) after that.
