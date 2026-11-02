package exercises_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRaceEnvSupported(t *testing.T) {
	yes, no := func() bool { return true }, func() bool { return false }
	tests := []struct {
		name string
		env  exercises.RaceEnv
		want bool
	}{
		{"linux with cgo and a compiler", exercises.RaceEnv{GOOS: "linux", CGOEnabled: yes, HasCCompiler: yes}, true},
		{"linux without cgo", exercises.RaceEnv{GOOS: "linux", CGOEnabled: no, HasCCompiler: yes}, false},
		{"linux without a compiler", exercises.RaceEnv{GOOS: "linux", CGOEnabled: yes, HasCCompiler: no}, false},
		{"windows without a compiler", exercises.RaceEnv{GOOS: "windows", CGOEnabled: yes, HasCCompiler: no}, false},
		{"macOS needs no cgo", exercises.RaceEnv{GOOS: "darwin", CGOEnabled: no, HasCCompiler: no}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.env.Supported(); got != tt.want {
				t.Errorf("Supported() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _ = Describe("Exercises that need the race detector", func() {
	var infoFile string

	// raceSupport fakes what the machine can do, for this test only.
	raceSupport := func(supported bool) {
		previous := exercises.RaceSupported
		exercises.RaceSupported = func() bool { return supported }
		DeferCleanup(func() { exercises.RaceSupported = previous })
	}

	BeforeEach(func() {
		dir := GinkgoT().TempDir()
		var info string
		for _, spec := range []struct {
			name string
			race bool
		}{{"racy", true}, {"plain", false}} {
			file := filepath.Join(dir, spec.name+"_test.go")
			Expect(os.WriteFile(file, []byte("// I AM NOT DONE\npackage main_test\n"), 0o644)).To(Succeed())
			info += fmt.Sprintf("[[exercises]]\nname = %q\npath = %q\nmode = \"test\"\nrace = %t\nhint = \"h\"\n\n", spec.name, file, spec.race)
		}
		infoFile = filepath.Join(dir, "info.toml")
		Expect(os.WriteFile(infoFile, []byte(info), 0o644)).To(Succeed())
	})

	It("reads the race flag from the manifest", func() {
		list, err := exercises.List(infoFile)

		Expect(err).NotTo(HaveOccurred())
		Expect(list[0].Race).To(BeTrue())
		Expect(list[1].Race).To(BeFalse())
	})

	Context("where the race detector is missing", func() {
		BeforeEach(func() { raceSupport(false) })

		It("is not run, and the error says what to do", func() {
			list, _ := exercises.List(infoFile)

			_, err := list[0].Run()

			Expect(err).To(MatchError(exercises.ErrRaceUnsupported))
			Expect(exercises.IsSkipped(err)).To(BeTrue())
			Expect(err.Error()).To(And(ContainSubstring("C compiler"), ContainSubstring("CGO_ENABLED=1"), ContainSubstring("golings doctor")))
		})
		It("is skipped, with the reason", func() {
			skipped, err := exercises.Skipped(infoFile)

			Expect(err).NotTo(HaveOccurred())
			Expect(skipped).To(HaveLen(1))
			Expect(skipped[0].Name).To(Equal("racy"))
			Expect(skipped[0].SkipReason()).To(Equal("needs the race detector"))
		})
		It("does not block the next exercise", func() {
			next, err := exercises.NextPending(infoFile)

			Expect(err).NotTo(HaveOccurred())
			Expect(next.Name).To(Equal("plain"))
		})
		It("is left out of the progress, and not counted as done", func() {
			_, done, total, err := exercises.Progress(infoFile)

			Expect(err).NotTo(HaveOccurred())
			Expect(done).To(Equal(0))
			Expect(total).To(Equal(1))
		})
		It("does not stop an exercise without the flag", func() {
			list, _ := exercises.List(infoFile)

			Expect(list[1].SkipReason()).To(BeEmpty())
		})
	})

	Context("where the race detector works", func() {
		BeforeEach(func() { raceSupport(true) })

		It("offers the exercise in order", func() {
			next, err := exercises.NextPending(infoFile)

			Expect(err).NotTo(HaveOccurred())
			Expect(next.Name).To(Equal("racy"))
			skipped, _ := exercises.Skipped(infoFile)
			Expect(skipped).To(BeEmpty())
		})
	})
})
