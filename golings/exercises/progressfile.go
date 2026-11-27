package exercises

import "github.com/cleonte/lingscore/progress"

// ProgressFile is where golings keeps what it cannot keep in an exercise file,
// next to info.toml. It is kept by lingscore's progress package. For now it
// holds the hash of every exercise file as golings wrote it (init, update,
// reset), which is how `update` tells a file the learner has not touched, safe
// to replace with a newer version, from one with their work in it, and the
// exercises `update` added. Whether an exercise is done is still decided by its
// "I AM NOT DONE" marker alone.
const ProgressFile = ".golings.json"

func store(dir string) progress.Store { return progress.Store{Dir: dir, File: ProgressFile} }

// Originals returns the recorded hashes of the exercise files written into the
// workspace in dir.
func Originals(dir string) map[string]string { return store(dir).Originals() }

// RecordOriginals remembers the hashes of exercise files that golings has just
// written into the workspace in dir, and which exercises `update` has just added.
func RecordOriginals(dir string, hashes map[string]string, added []string) error {
	return store(dir).RecordOriginals(hashes, added)
}

// IsNew reports whether `update` added the exercise to the workspace in dir.
func IsNew(dir, name string) bool { return store(dir).IsNew(name) }
