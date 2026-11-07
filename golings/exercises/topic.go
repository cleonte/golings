package exercises

import (
	"io/fs"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

// Topic is the folder an exercise lives in: for
// `exercises/variables/variables1/main.go` it is "variables". Exercises that
// aren't laid out as `exercises/<topic>/<name>/...` have no topic.
func (e Exercise) Topic() string {
	parts := strings.Split(filepath.ToSlash(e.Path), "/")
	for i, part := range parts {
		if part == "exercises" && len(parts) > i+3 {
			return parts[i+1]
		}
	}
	return ""
}

// SolutionPath is where the solution of the exercise lives, relative to the
// root of the solutions: `exercises/slices/slices1/main.go` becomes
// `solutions/slices/slices1/main.go`. It is empty for exercises that aren't
// laid out as `exercises/<topic>/<name>/...`.
func (e Exercise) SolutionPath() string {
	parts := strings.Split(filepath.ToSlash(e.Path), "/")
	for i, part := range parts {
		if part == "exercises" && len(parts) > i+3 {
			return "solutions/" + strings.Join(parts[i+1:], "/")
		}
	}
	return ""
}

// SolutionDir is the folder of the solution, laid out like the folder of the
// exercise: `solutions/slices/slices1`. It is empty when SolutionPath is.
func (e Exercise) SolutionDir() string {
	if p := e.SolutionPath(); p != "" {
		return path.Dir(p)
	}
	return ""
}

// SolutionFiles lists the files of the solution in src, as paths relative to
// the root of the solutions, with the main file first and the rest in order.
// A solution has the files that replace or add to the files of the exercise
// folder with the same names.
func (e Exercise) SolutionFiles(src fs.FS) ([]string, error) {
	dir, main := e.SolutionDir(), e.StoredName(e.SolutionPath())
	if dir == "" {
		return nil, fs.ErrNotExist
	}
	var files []string
	err := fs.WalkDir(src, dir, func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	slices.Sort(files)
	if i := slices.Index(files, main); i > 0 {
		files = slices.Insert(slices.Delete(files, i, i+1), 0, main)
	}
	return files, nil
}

type TopicProgress struct {
	Topic string
	Done  int
	Total int
}

// ProgressByTopic counts done and total exercises per topic, in the order
// topics first appear in the info file. Like Progress, it leaves out exercises
// the local Go can't run.
func ProgressByTopic(infoFile string) ([]TopicProgress, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return nil, err
	}

	var topics []TopicProgress
	index := map[string]int{}
	for _, exercise := range allExercises {
		if !runnable(exercise) {
			continue
		}
		topic := exercise.Topic()
		i, ok := index[topic]
		if !ok {
			i = len(topics)
			index[topic] = i
			topics = append(topics, TopicProgress{Topic: topic})
		}
		topics[i].Total++
		if exercise.State() == Done {
			topics[i].Done++
		}
	}
	return topics, nil
}
