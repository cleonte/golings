package workspace

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
)

// UpdateResult says what Update did (or, for a dry run, would do).
type UpdateResult struct {
	// Added are the exercises of the binary that the workspace did not have.
	Added []exercises.Exercise
	// Readmes is how many README.md files were replaced by a newer text.
	Readmes int
}

var (
	infoBlock = regexp.MustCompile(`(?m)^\[\[exercises\]\]`)
	infoName  = regexp.MustCompile(`(?m)^name\s*=\s*"([^"]+)"`)
)

// splitInfo cuts the text of an info.toml into what comes before the first
// exercise and the text of each exercise, in order. It works on the text, not on
// the parsed manifest, so that nothing else in the file is rewritten.
func splitInfo(text string) (header string, names []string, blocks map[string]string) {
	blocks = map[string]string{}
	locs := infoBlock.FindAllStringIndex(text, -1)
	if len(locs) == 0 {
		return text, nil, blocks
	}
	header = text[:locs[0][0]]
	for i, loc := range locs {
		end := len(text)
		if i+1 < len(locs) {
			end = locs[i+1][0]
		}
		block := text[loc[0]:end]
		if m := infoName.FindStringSubmatch(block); m != nil {
			names = append(names, m[1])
			blocks[m[1]] = block
		}
	}
	return header, names, blocks
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
			if err := writeFile(target, data, true); err != nil {
				return res, err
			}
		}
	}

	if dryRun {
		return res, nil
	}
	for _, e := range res.Added {
		if err := copyExercise(src, root, e); err != nil {
			return res, err
		}
	}
	if len(res.Added) > 0 {
		merged := mergeInfo(embeddedText, mineText, embedded, res.Added)
		if err := os.WriteFile(filepath.Join(root, "info.toml.bak"), []byte(mineText), 0o644); err != nil {
			return res, err
		}
		if err := os.WriteFile(filepath.Join(root, "info.toml"), []byte(merged), 0o644); err != nil {
			return res, err
		}
	}
	return res, nil
}

// mergeInfo puts the entries of the added exercises into the info.toml of the
// workspace, each after the exercise that comes before it in the binary's order.
// Entries of the workspace that the binary does not know are left where they are.
func mergeInfo(embeddedText, workspaceText string, embedded, added []exercises.Exercise) string {
	header, names, blocks := splitInfo(workspaceText)
	_, _, embeddedBlocks := splitInfo(embeddedText)
	isAdded := map[string]bool{}
	for _, e := range added {
		isAdded[e.Name] = true
	}
	prev := ""
	for _, e := range embedded {
		if !isAdded[e.Name] {
			prev = e.Name
			continue
		}
		at := 0
		if i := slices.Index(names, prev); prev != "" && i >= 0 {
			at = i + 1
		}
		blocks[e.Name] = embeddedBlocks[e.Name]
		names = slices.Insert(names, at, e.Name)
		prev = e.Name
	}
	var b strings.Builder
	b.WriteString(header)
	for i, n := range names {
		block := blocks[n]
		if i < len(names)-1 && !strings.HasSuffix(block, "\n\n") {
			block = strings.TrimRight(block, "\n") + "\n\n"
		}
		b.WriteString(block)
	}
	return b.String()
}

// copyExercise writes the folder of one exercise into the workspace, with the
// names a learner sees (a module exercise is stored with a .txt added to its
// files), and never over a file that is already there.
func copyExercise(src fs.FS, root string, e exercises.Exercise) error {
	return fs.WalkDir(src, path.Dir(e.Path), func(name string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := fs.ReadFile(src, name)
		if err != nil {
			return err
		}
		return writeFile(filepath.Join(root, filepath.FromSlash(e.LearnerName(name))), data, false)
	})
}

// writeFile writes a file, creating its folder. Without overwrite it leaves an
// existing file alone.
func writeFile(target string, data []byte, overwrite bool) error {
	if !overwrite {
		if _, err := os.Stat(target); err == nil {
			return nil
		} else if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	return os.WriteFile(target, data, 0o644)
}

// Describe is a short phrase for a list of exercises: "config5, config6 and 24 more".
func Describe(names []string) string {
	if len(names) <= 3 {
		return strings.Join(names, ", ")
	}
	return fmt.Sprintf("%s and %d more", strings.Join(names[:3], ", "), len(names)-3)
}
