package workspace

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/cleonte/golings/golings/exercises"
	lw "github.com/cleonte/lingscore/workspace"
)

// UpdateResult says what Update did (or, for a dry run, would do).
type UpdateResult struct {
	// Added are the exercises of the binary that the workspace did not have.
	Added []exercises.Exercise
	// Readmes is how many README.md files were replaced by a newer text.
	Readmes int
	// Refreshed are exercises whose file was replaced by a newer version because
	// the learner had not touched the old one.
	Refreshed []string
	// Newer are exercises whose file has a newer version that was NOT applied,
	// because the learner has changed the old one: `reset` brings it, and loses
	// their work.
	Newer []string
}

func readInfo(src fs.FS, root string) (embedded []exercises.Exercise, embeddedText string, mine []exercises.Exercise, mineText string, err error) {
	raw, err := fs.ReadFile(src, "info.toml")
	if err != nil {
		return nil, "", nil, "", err
	}
	embedded, err = exercises.ParseInfo(raw)
	if err != nil {
		return nil, "", nil, "", err
	}
	own, err := os.ReadFile(filepath.Join(root, "info.toml"))
	if err != nil {
		return nil, "", nil, "", err
	}
	mine, err = exercises.ParseInfo(own)
	if err != nil {
		return nil, "", nil, "", err
	}
	return embedded, string(raw), mine, string(own), nil
}

// Missing returns the names of the exercises that the binary brings and the
// workspace in root does not list yet, in the order of the binary. It never
// fails loudly: a workspace that cannot be read has nothing to compare.
func Missing(src fs.FS, root string) []string {
	embedded, _, mine, _, err := readInfo(src, root)
	if err != nil {
		return nil
	}
	have := map[string]bool{}
	for _, e := range mine {
		have[e.Name] = true
	}
	var missing []string
	for _, e := range embedded {
		if !have[e.Name] {
			missing = append(missing, e.Name)
		}
	}
	return missing
}

// Update brings the workspace in root up to date with the exercises in src: it
// adds the exercises the workspace does not have, each in the place of the
// order of the course where it belongs, and replaces the README of a topic by its
// newer text. It never writes over a file of an exercise the workspace already
// has, since that may hold the learner's work, and it keeps a copy of info.toml
// as info.toml.bak. With dryRun it only reports.
func Update(src fs.FS, root string, dryRun bool) (UpdateResult, error) {
	var res UpdateResult
	embedded, embeddedText, mine, mineText, err := readInfo(src, root)
	if err != nil {
		return res, err
	}
	have := map[string]bool{}
	for _, e := range mine {
		have[e.Name] = true
	}
	for _, e := range embedded {
		if !have[e.Name] {
			res.Added = append(res.Added, e)
		}
	}

	// Exercises that are already there: a newer file replaces the old one where
	// the learner cannot have changed it. Only exercises that are a single file
	// are handled; a folder of several files is left to `reset`.
	originals := exercises.Originals(root)
	record := map[string]string{}
	for _, e := range embedded {
		if !have[e.Name] {
			continue
		}
		stored := e.StoredName(e.Path)
		if entries, err := fs.ReadDir(src, path.Dir(stored)); err != nil || len(entries) != 1 {
			continue
		}
		original, err := fs.ReadFile(src, stored)
		if err != nil {
			continue
		}
		target := filepath.Join(root, filepath.FromSlash(e.Path))
		current, err := os.ReadFile(target)
		if err != nil {
			continue
		}
		recorded, known := originals[e.Name]
		switch lw.Decide(current, original, recorded, known) {
		case lw.Record:
			record[e.Name] = lw.Hash(original) // untouched and current: a later release can refresh it
		case lw.Refresh:
			res.Refreshed = append(res.Refreshed, e.Name)
			if !dryRun {
				if err := lw.WriteFile(target, original, true); err != nil {
					return res, err
				}
				record[e.Name] = lw.Hash(original)
			}
		case lw.Newer:
			res.Newer = append(res.Newer, e.Name)
		}
	}

	// The README of a topic is text nobody edits: bring the newer one. A topic
	// the workspace has none of comes with its first added exercise.
	newTopic := map[string]bool{}
	for _, e := range res.Added {
		newTopic[path.Dir(path.Dir(e.Path))] = true
	}
	readmes, err := fs.Glob(src, "exercises/*/README.md")
	if err != nil {
		return res, err
	}
	for _, name := range readmes {
		data, err := fs.ReadFile(src, name)
		if err != nil {
			return res, err
		}
		target := filepath.Join(root, filepath.FromSlash(name))
		old, err := os.ReadFile(target)
		switch {
		case err == nil && string(old) == string(data):
			continue
		case err == nil:
			res.Readmes++
		case newTopic[path.Dir(name)]:
			// Not a replacement: it is part of the new topic.
		default:
			continue
		}
		if !dryRun {
			if err := lw.WriteFile(target, data, true); err != nil {
				return res, err
			}
		}
	}

	if dryRun {
		return res, nil
	}
	for _, e := range res.Added {
		if err := copyExercise(src, root, e, record); err != nil {
			return res, err
		}
	}
	if len(res.Added) > 0 {
		merged := lw.MergeInfo(embeddedText, mineText, exerciseNames(embedded), exerciseNames(res.Added), nil)
		if err := os.WriteFile(filepath.Join(root, "info.toml.bak"), []byte(mineText), 0o644); err != nil {
			return res, err
		}
		if err := os.WriteFile(filepath.Join(root, "info.toml"), []byte(merged), 0o644); err != nil {
			return res, err
		}
	}
	if len(record) > 0 || len(res.Added) > 0 {
		if err := exercises.RecordOriginals(root, record, exerciseNames(res.Added)); err != nil {
			return res, err
		}
	}
	return res, nil
}

// copyExercise writes the folder of one exercise into the workspace, with the
// names a learner sees (a module exercise is stored with a .txt added to its
// files), and never over a file that is already there.
func copyExercise(src fs.FS, root string, e exercises.Exercise, record map[string]string) error {
	return fs.WalkDir(src, path.Dir(e.Path), func(name string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := fs.ReadFile(src, name)
		if err != nil {
			return err
		}
		if name == e.StoredName(e.Path) {
			record[e.Name] = lw.Hash(data)
		}
		return lw.WriteFile(filepath.Join(root, filepath.FromSlash(e.LearnerName(name))), data, false)
	})
}

// exerciseNames are the names of the exercises, in order.
func exerciseNames(list []exercises.Exercise) []string {
	out := make([]string, len(list))
	for i, e := range list {
		out[i] = e.Name
	}
	return out
}

// Describe is a short phrase for a list of exercises: "config5, config6 and 24 more".
func Describe(names []string) string { return lw.Describe(names) }
