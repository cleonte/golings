package exercises_test

import (
	"go/version"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// The tests in this file run against the real manifest and exercises of the
// repository, so mistakes in info.toml are caught before a learner finds them.
var _ = Describe("Repository manifest", func() {
	const (
		infoFile     = "../../info.toml"
		exercisesDir = "../../exercises"
	)

	var all []exercises.Exercise

	BeforeEach(func() {
		var err error
		all, err = exercises.List(infoFile)
		Expect(err).NotTo(HaveOccurred())
		Expect(all).NotTo(BeEmpty())
	})

	It("has unique exercise names", func() {
		seen := map[string]bool{}
		for _, ex := range all {
			Expect(seen).NotTo(HaveKey(ex.Name), "duplicate exercise name %q", ex.Name)
			seen[ex.Name] = true
		}
	})

	It("uses a known mode for every exercise", func() {
		for _, ex := range all {
			Expect([]string{"compile", "test"}).To(ContainElement(ex.Mode), "exercise %s", ex.Name)
		}
	})

	It("has a hint for every exercise", func() {
		for _, ex := range all {
			Expect(ex.Hint).NotTo(BeEmpty(), "exercise %s", ex.Name)
		}
	})

	It("has all three hint levels for every exercise", func() {
		for _, ex := range all {
			Expect(ex.Nudge).NotTo(BeEmpty(), "exercise %s has no nudge: say what the failure means", ex.Name)
			Expect(ex.Stronger).NotTo(BeEmpty(), "exercise %s has no stronger hint: say where to change", ex.Name)
		}
	})

	It("has a nudge before a stronger hint, and smaller hints that differ from the full one", func() {
		for _, ex := range all {
			if ex.Stronger != "" {
				Expect(ex.Nudge).NotTo(BeEmpty(), "exercise %s has a stronger hint but no nudge: the levels must not skip", ex.Name)
			}
			levels := ex.HintLevels()
			for i := range levels {
				for j := i + 1; j < len(levels); j++ {
					Expect(levels[i]).NotTo(Equal(levels[j]), "exercise %s repeats a hint", ex.Name)
				}
			}
		}
	})

	It("points every exercise at an existing file", func() {
		for _, ex := range all {
			Expect(filepath.Join("../..", ex.StoredName(ex.Path))).To(BeARegularFile(), "exercise %s", ex.Name)
		}
	})

	It("ships every exercise unsolved", func() {
		for _, ex := range all {
			data, err := os.ReadFile(filepath.Join("../..", ex.StoredName(ex.Path)))
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(MatchRegexp(`(?m)^\s*///?\s*I\s+AM\s+NOT\s+DONE`), "exercise %s", ex.Name)
		}
	})

	It("stores the files of a module exercise with a .txt, and no other exercise has a go.mod", func() {
		// go:embed leaves out a folder that holds a go.mod, and a go.mod or Go
		// files would make the folder a part of (or a module inside) the module
		// of this repository, so they only exist under a .txt name.
		modules := map[string]bool{}
		for _, ex := range all {
			modules[filepath.ToSlash(ex.Dir())] = ex.Module
		}
		for dir, module := range modules {
			err := filepath.WalkDir(filepath.Join("../..", filepath.FromSlash(dir)), func(name string, d fs.DirEntry, err error) error {
				if err != nil || d.IsDir() {
					return err
				}
				switch base := d.Name(); {
				case module:
					Expect(base).To(HaveSuffix(".txt"), "%s is in a module exercise, so it is stored with a .txt", name)
				default:
					Expect([]string{"go.mod", "go.sum", "go.work"}).NotTo(ContainElement(base), "%s: only a module exercise has module files", name)
				}
				return nil
			})
			Expect(err).NotTo(HaveOccurred())
		}
	})

	It("keeps test-mode exercises next to a test file", func() {
		for _, ex := range all {
			if ex.Mode != "test" {
				continue
			}
			tests, err := filepath.Glob(filepath.Join("../..", filepath.Dir(ex.Path), ex.StoredName("*_test.go")))
			Expect(err).NotTo(HaveOccurred())
			Expect(tests).NotTo(BeEmpty(), "exercise %s", ex.Name)
		}
	})

	It("sets race only on test-mode exercises", func() {
		for _, ex := range all {
			if ex.Race {
				Expect(ex.Mode).To(Equal("test"), "exercise %s: -race only applies to tests", ex.Name)
			}
		}
	})

	It("marks the exercises that only fail with the race detector", func() {
		var racy []string
		for _, ex := range all {
			if ex.Race {
				racy = append(racy, ex.Name)
			}
		}
		Expect(racy).To(Equal([]string{"concurrent2", "concurrent10"}),
			"set race = true only on an exercise whose unsolved version does not fail without -race (see CONTRIBUTING.md)")
	})

	It("tags the exercises that need a newer Go with its go_version", func() {
		// What only a newer Go has, and the version that brought it. A learner
		// on an older Go would otherwise get a confusing build error, not a skip.
		needs := []struct {
			pattern *regexp.Regexp
			version string
		}{
			{regexp.MustCompile(`(?m)^//go:build go1\.27`), "1.27"},
			{regexp.MustCompile(`"uuid"`), "1.27"},
			{regexp.MustCompile(`"encoding/json/v2"`), "1.27"},
			{regexp.MustCompile(`\bstrings\.CutLast\b|\bbytes\.CutLast\b`), "1.27"},
			{regexp.MustCompile(`\bsynctest\.Sleep\b`), "1.27"},
		}

		for _, ex := range all {
			// Every Go file of the exercise folder and of the solution folder.
			var sources []string
			for _, dir := range []string{ex.Dir(), ex.SolutionDir()} {
				found, err := filepath.Glob(filepath.Join("../..", dir, "*.go"))
				Expect(err).NotTo(HaveOccurred())
				sources = append(sources, found...)
			}
			for _, file := range sources {
				data, err := os.ReadFile(file)
				Expect(err).NotTo(HaveOccurred())
				for _, need := range needs {
					if need.pattern.Match(data) {
						Expect(ex.GoVersion).NotTo(BeEmpty(), "%s uses %s, which needs Go %s: set go_version", file, need.pattern, need.version)
						Expect(version.Compare("go"+ex.GoVersion, "go"+need.version)).To(BeNumerically(">=", 0), "%s uses %s, which needs Go %s, but go_version is %s", file, need.pattern, need.version, ex.GoVersion)
					}
				}
			}
		}
	})

	It("puts the //go:build line of its go_version at the top of every file of the exercise and its solution", func() {
		// The folder is built as a package, so the compiler refuses what is newer
		// than the module unless the file says which Go it is for.
		for _, ex := range all {
			if ex.GoVersion == "" {
				continue
			}
			for _, dir := range []string{ex.Dir(), ex.SolutionDir()} {
				files, err := filepath.Glob(filepath.Join("../..", dir, "*.go"))
				Expect(err).NotTo(HaveOccurred())
				for _, file := range files {
					data, err := os.ReadFile(file)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(HavePrefix("//go:build go"+ex.GoVersion+"\n"), "%s needs //go:build go%s as its first line", file, ex.GoVersion)
				}
			}
		}
	})

	It("tags every exercise of the go127 topic with go_version 1.27", func() {
		for _, ex := range all {
			if ex.Topic() == "go127" {
				Expect(ex.GoVersion).To(Equal("1.27"), "exercise %s", ex.Name)
			}
		}
	})

	It("lists every exercise directory in the manifest", func() {
		listed := []string{}
		for _, ex := range all {
			listed = append(listed, filepath.Dir(ex.Path))
		}

		var onDisk []string
		topics, err := os.ReadDir(exercisesDir)
		Expect(err).NotTo(HaveOccurred())
		for _, topic := range topics {
			if !topic.IsDir() {
				continue
			}
			dirs, err := os.ReadDir(filepath.Join(exercisesDir, topic.Name()))
			Expect(err).NotTo(HaveOccurred())
			for _, d := range dirs {
				if d.IsDir() {
					onDisk = append(onDisk, filepath.Join("exercises", topic.Name(), d.Name()))
				}
			}
		}

		for _, dir := range onDisk {
			Expect(slices.Contains(listed, dir)).To(BeTrue(), "%s is not in info.toml", dir)
		}
	})
})
