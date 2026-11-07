# Machine-readable output (`--json`)

`golings list --json` and `golings run --json` print one JSON document to standard output, for editors, scripts and other tools. The text output is for people and may change; this one is kept stable.

Run them from the folder that holds `info.toml`, like every other command. Nothing else is written to standard output: no colors, no spinner, no progress bar.

## Versioning

Every document has `"version": 1`. The version changes only when a field is renamed, removed or changes its meaning. A new field does not change it, so a reader must ignore the fields it does not know.

## Exit status

- `list --json` exits 0 when it printed the list.
- `run --json` exits 0 whenever it could tell what happened, including when the exercise fails: the answer is in `status`. This is different from `golings run`, which exits 1 for a failing exercise.
- Both exit 1 when they cannot answer: the info file is missing, the exercise does not exist, nothing is pending for `next`, or the exercise could not be run. They then print `{"version": 1, "error": "..."}` and no other field.

## `golings list --json`

```json
{
  "version": 1,
  "done": 12,
  "total": 262,
  "topics": [{ "name": "variables", "done": 6, "total": 6 }],
  "exercises": [
    {
      "name": "variables1",
      "topic": "variables",
      "path": "exercises/variables/variables1/main.go",
      "mode": "compile",
      "state": "done"
    }
  ]
}
```

| Field | Meaning |
|---|---|
| `done`, `total` | Exercises finished and in all. Exercises this machine can't run are in neither. |
| `topics` | Progress per topic, in the order of the info file, with the same counting. |
| `exercises[].path` | The main file of the exercise, relative to the folder of `info.toml`, with forward slashes. |
| `exercises[].mode` | `compile` or `test`. |
| `exercises[].state` | `pending` (the file still has the `// I AM NOT DONE` marker) or `done`. |
| `exercises[].skipped` | Only there when this machine can't run the exercise: why, for example `needs Go 1.27` or `needs the race detector`. |

`topics` and `exercises` are `[]` when empty, never `null`.

## `golings run --json <name | next>`

```json
{
  "version": 1,
  "exercise": { "name": "variables1", "topic": "variables", "path": "exercises/variables/variables1/main.go", "mode": "compile", "state": "pending" },
  "status": "failed",
  "failure": {
    "kind": "compile",
    "diagnostics": [
      { "file": "exercises/variables/variables1/main.go", "line": 11, "column": 6, "message": "syntax error: unexpected =, expected name" }
    ]
  }
}
```

`exercise` is described as in the list. `status` is one of:

| Status | Meaning |
|---|---|
| `passed` | It built and ran, and the marker is gone. `output` is what it printed. |
| `pending` | It passed, but still has the `// I AM NOT DONE` marker. `output` is what it printed. |
| `failed` | It did not build, its tests failed, or it crashed. See `failure`. |
| `timeout` | It was stopped for running too long. See `failure`. |
| `skipped` | This machine can't run it. `message` says why. |

`failure` has:

| Field | Meaning |
|---|---|
| `kind` | `compile`, `test`, `runtime` or `timeout`. |
| `diagnostics` | For `compile`: the compiler messages (`file`, `line`, `column` when the compiler gave one, `message`). |
| `output` | The rest worth showing, without toolchain noise. Long panics and data race reports are cut down to the lines in the exercise. |
| `full` | All of the output, only there when `output` was cut down. |
| `timeout_seconds` | For `timeout`: how long the exercise was given. |

Running an exercise executes the learner's code, as `golings run` does.
