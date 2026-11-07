package exercises_test

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// The capstone is one program in steps, and each step starts from the finished
// previous one. This checks that: every file of the solved step N-1 (its folder
// with its solution on top) is in the folder of step N with the same content,
// unless step N changes it (it is in the solution of N) or it is the main file
// or a test, which are new in each step. The import path of a package of its
// own has the step in it, so the step numbers are not compared.
func TestCapstoneStepsStartFromTheFinishedPreviousStep(t *testing.T) {
	step := regexp.MustCompile(`capstone[0-9]+`)
	read := func(dir string) map[string]string {
		files := map[string]string{}
		root := filepath.Join(repoRoot, filepath.FromSlash(dir))
		err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			source, err := os.ReadFile(p)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(root, p)
			files[filepath.ToSlash(rel)] = string(source)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		return files
	}
	solved := func(n int) map[string]string {
		files := read(fmt.Sprintf("exercises/capstone/capstone%d", n))
		for name, source := range read(fmt.Sprintf("solutions/capstone/capstone%d", n)) {
			files[name] = source
		}
		return files
	}

	for n := 2; ; n++ {
		if _, err := os.Stat(filepath.Join(repoRoot, "exercises", "capstone", fmt.Sprintf("capstone%d", n))); err != nil {
			break
		}
		previous := solved(n - 1)
		current := read(fmt.Sprintf("exercises/capstone/capstone%d", n))
		changed := read(fmt.Sprintf("solutions/capstone/capstone%d", n))

		for name, want := range previous {
			if path.Base(name) == "main.go" || strings.HasSuffix(name, "_test.go") {
				continue
			}
			if _, ok := changed[name]; ok {
				continue
			}
			got, ok := current[name]
			if !ok {
				t.Errorf("capstone%d lacks %s, which the finished capstone%d has", n, name, n-1)
				continue
			}
			if step.ReplaceAllString(got, "capstoneN") != step.ReplaceAllString(want, "capstoneN") {
				t.Errorf("capstone%d: %s is not the finished one of capstone%d", n, name, n-1)
			}
		}
	}
}
