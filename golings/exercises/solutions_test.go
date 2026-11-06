package exercises_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/cleonte/golings/golings/doctor"
	"github.com/cleonte/golings/golings/exercises"
)

// allSolutionsRequired makes an exercise without a solution fail the build.
// Every exercise has one now, so a new exercise has to come with its solution.
const allSolutionsRequired = true

const repoRoot = "../.."

// TestSolutions proves each exercise can be solved: the unsolved version has
// to fail and the solution has to pass, through the same runner learners use,
// so with the race detector exactly when the manifest says race = true.
func TestSolutions(t *testing.T) {
	if testing.Short() {
		t.Skip("runs every exercise with go run / go test; skipped in -short mode")
	}

	all, err := exercises.List(filepath.Join(repoRoot, "info.toml"))
	if err != nil {
		t.Fatal(err)
	}

	for _, ex := range all {
		solutionPath := ex.SolutionPath()
		if solutionPath == "" {
			t.Errorf("%s: path %q is not laid out as exercises/<topic>/<name>/...", ex.Name, ex.Path)
			continue
		}
		files, err := ex.SolutionFiles(os.DirFS(repoRoot))
		if err != nil || len(files) == 0 {
			if allSolutionsRequired {
				t.Errorf("%s: missing solution in %s", ex.Name, ex.SolutionDir())
			}
			continue
		}

		t.Run(ex.Name, func(t *testing.T) {
			t.Parallel()

			// An exercise this machine can't run: a Go older than its go_version
			// (the Go 1.27 ones on 1.26), or no race detector (cgo and a C compiler).
			if reason := ex.SkipReason(); reason != "" {
				t.Skipf("%s %s here", ex.Name, reason)
			}

			for _, file := range files {
				source, err := os.ReadFile(filepath.Join(repoRoot, filepath.FromSlash(file)))
				if err != nil {
					t.Fatal(err)
				}
				if strings.Contains(string(source), "I AM NOT DONE") {
					t.Errorf("%s still has the 'I AM NOT DONE' marker", file)
				}
			}

			unsolved := ex
			unsolved.Path = filepath.Join(repoRoot, filepath.FromSlash(ex.Path))
			if result, err := unsolved.Run(); err == nil {
				t.Errorf("the unsolved exercise passes, it should fail:\n%s", result.Out)
			}

			solved := inAFolderOfItsOwn(t, ex, files)
			if result, err := solved.Run(); err != nil {
				t.Errorf("the solution fails: %v\n%s%s", err, result.Err, result.Out)
			}
		})
	}
}

// inAFolderOfItsOwn copies the exercise folder to a new module and puts the
// files of the solution over it, as a learner's folder would look after
// finishing the exercise, and returns the exercise to run there. The files of a
// solution replace the files of the exercise with the same name, or add to them.
func inAFolderOfItsOwn(t *testing.T, ex exercises.Exercise, solutionFiles []string) exercises.Exercise {
	t.Helper()
	root := t.TempDir()

	goMod := "module " + doctor.ExercisesModule + "\n\ngo " + doctor.MinGoVersion + "\n"
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte(goMod), 0o644); err != nil {
		t.Fatal(err)
	}

	copyTree := func(from, to string) {
		err := filepath.WalkDir(from, func(name string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(from, name)
			if err != nil {
				return err
			}
			target := filepath.Join(to, rel)
			if d.IsDir() {
				return os.MkdirAll(target, 0o755)
			}
			data, err := os.ReadFile(name)
			if err != nil {
				return err
			}
			return os.WriteFile(target, data, 0o644)
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	copyTree(filepath.Join(repoRoot, ex.Dir()), filepath.Join(root, ex.Dir()))
	for _, file := range solutionFiles {
		rel := strings.TrimPrefix(file, ex.SolutionDir()+"/")
		target := filepath.Join(root, ex.Dir(), filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(filepath.Join(repoRoot, filepath.FromSlash(file)))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, data, 0o644); err != nil {
			t.Fatal(err)
		}
	}

	solved := ex
	solved.Root = root
	return solved
}

// TestNoOrphanedSolutions catches solutions left behind after an exercise was
// renamed or removed, and stray files in the solutions tree.
func TestNoOrphanedSolutions(t *testing.T) {
	all, err := exercises.List(filepath.Join(repoRoot, "info.toml"))
	if err != nil {
		t.Fatal(err)
	}
	// A solution folder may hold several files, so a file belongs to an
	// exercise when it is in the solution folder of one.
	var folders []string
	for _, ex := range all {
		folders = append(folders, ex.SolutionDir()+"/")
	}

	root := filepath.Join(repoRoot, "solutions")
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(repoRoot, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if !slices.ContainsFunc(folders, func(folder string) bool { return strings.HasPrefix(rel, folder) }) {
			t.Errorf("%s does not belong to any exercise in info.toml", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
