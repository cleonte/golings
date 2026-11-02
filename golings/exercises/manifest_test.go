package exercises_test

import (
	"go/version"
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

	It("points every exercise at an existing file", func() {
		for _, ex := range all {
			Expect(filepath.Join("../..", ex.Path)).To(BeARegularFile(), "exercise %s", ex.Name)
		}
	})

	It("ships every exercise unsolved", func() {
		for _, ex := range all {
			data, err := os.ReadFile(filepath.Join("../..", ex.Path))
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(MatchRegexp(`(?m)^\s*///?\s*I\s+AM\s+NOT\s+DONE`), "exercise %s", ex.Name)
		}
	})

	It("keeps test-mode exercises next to a test file", func() {
		for _, ex := range all {
			if ex.Mode != "test" {
				continue
			}
			tests, err := filepath.Glob(filepath.Join("../..", filepath.Dir(ex.Path), "*_test.go"))
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
			for _, file := range []string{ex.Path, ex.SolutionPath()} {
				data, err := os.ReadFile(filepath.Join("../..", file))
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
