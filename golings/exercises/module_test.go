package exercises_test

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing/fstest"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	appSource   = "package main\n\nimport (\n\t\"fmt\"\n\n\t\"example.com/greet\"\n)\n\nfunc main() { fmt.Println(greet.Hello()) }\n"
	greetSource = "package greet\n\nfunc Hello() string { return \"hello from a module\" }\n"
)

// moduleExercise writes an exercise that is a module of its own, with a local
// dependency in a folder of it, and returns it with go.mod as the main file.
func moduleExercise(goMod string, more map[string]string) exercises.Exercise {
	files := map[string]string{
		"go.mod":         goMod,
		"main.go":        appSource,
		"greet/go.mod":   "module example.com/greet\n\ngo 1.26\n",
		"greet/greet.go": greetSource,
	}
	for name, source := range more {
		files[name] = source
	}
	ex := writeFolderExercise("compile", files, "go.mod")
	ex.Module = true
	return ex
}

var _ = Describe("Module exercises", func() {
	It("builds in the folder of the exercise, with a dependency that is a folder on disk", func() {
		ex := moduleExercise("module example.com/app\n\ngo 1.26\n\nrequire example.com/greet v0.0.0\n\nreplace example.com/greet => ./greet\n", nil)

		result, err := ex.Run()

		Expect(err).NotTo(HaveOccurred())
		Expect(result.Out).To(ContainSubstring("hello from a module"))
	})

	It("fails while the go.mod does not ask for the dependency", func() {
		ex := moduleExercise("module example.com/app\n\ngo 1.26\n", nil)

		result, err := ex.Run()

		Expect(err).To(HaveOccurred())
		Expect(result.BuildFailed).To(BeTrue())
		Expect(result.Err).To(ContainSubstring("example.com/greet"))
	})

	It("never goes to the network for a dependency", func() {
		ex := moduleExercise("module example.com/app\n\ngo 1.26\n\nrequire example.com/greet v1.2.3\n", nil)

		_, err := ex.Run()

		Expect(err).To(HaveOccurred())
	})

	It("does not download a newer Go for a go.mod that asks for one", func() {
		ex := moduleExercise("module example.com/app\n\ngo 1.99\n\nrequire example.com/greet v0.0.0\n\nreplace example.com/greet => ./greet\n", nil)

		result, err := ex.Run()

		Expect(err).To(HaveOccurred())
		Expect(result.Err).To(ContainSubstring("1.99"))
	})

	It("ignores a go.work of the learner's own", func() {
		ex := moduleExercise("module example.com/app\n\ngo 1.26\n\nrequire example.com/greet v0.0.0\n\nreplace example.com/greet => ./greet\n", nil)
		// a workspace of the learner's own, which would not build this module
		work := filepath.Join(GinkgoT().TempDir(), "go.work")
		Expect(os.WriteFile(work, []byte("go 1.26\n\nuse ./nowhere\n"), 0o644)).To(Succeed())
		GinkgoT().Setenv("GOWORK", work)

		_, err := ex.Run()

		Expect(err).NotTo(HaveOccurred())
	})

	It("uses the go.work of the exercise when it has one", func() {
		ex := moduleExercise("module example.com/app\n\ngo 1.26\n", map[string]string{
			"go.work": "go 1.26\n\nuse (\n\t.\n\t./greet\n)\n",
		})

		result, err := ex.Run()

		Expect(err).NotTo(HaveOccurred())
		Expect(result.Out).To(ContainSubstring("hello from a module"))
	})

	It("finds the go.work of the exercise when the folder is reached through a symbolic link", func() {
		if runtime.GOOS == "windows" {
			Skip("symbolic links need privileges on Windows")
		}
		ex := moduleExercise("module example.com/app\n\ngo 1.26\n", map[string]string{
			"go.work": "go 1.26\n\nuse (\n\t.\n\t./greet\n)\n",
		})
		real, err := filepath.Abs(filepath.Join(ex.Root, filepath.Dir(ex.Path)))
		Expect(err).NotTo(HaveOccurred())
		link := filepath.Join(GinkgoT().TempDir(), "link")
		Expect(os.Symlink(real, link)).To(Succeed())
		ex.Root, ex.Path = link, "go.mod"

		result, err := ex.Run()

		Expect(err).NotTo(HaveOccurred(), result.Err)
	})

	It("says to run golings init where the files are not in place", func() {
		ex := writeFolderExercise("compile", map[string]string{"go.mod.txt": "module example.com/app\n"}, "go.mod.txt")
		ex.Module = true
		ex.Path = filepath.Join(filepath.Dir(ex.Path), "go.mod")

		_, err := ex.Run()

		Expect(errors.Is(err, exercises.ErrNeedsInit)).To(BeTrue())
	})

	It("stores the files with a .txt, and names them without it for the learner", func() {
		ex := exercises.Exercise{Path: "exercises/modules/modules1/go.mod", Module: true}
		Expect(ex.StoredName("go.mod")).To(Equal("go.mod.txt"))
		Expect(ex.LearnerName("greet/go.mod.txt")).To(Equal("greet/go.mod"))

		plain := exercises.Exercise{Path: "exercises/a/a1/main.go"}
		Expect(plain.StoredName("main.go")).To(Equal("main.go"))
		Expect(plain.LearnerName("notes.txt")).To(Equal("notes.txt"))
	})

	It("lists the solution files with their stored names, the main one first", func() {
		ex := exercises.Exercise{Name: "m1", Path: "exercises/m/m1/go.mod", Module: true}
		src := fstest.MapFS{
			"solutions/m/m1/main.go.txt": {Data: []byte("x")},
			"solutions/m/m1/go.mod.txt":  {Data: []byte("x")},
		}

		files, err := ex.SolutionFiles(src)

		Expect(err).NotTo(HaveOccurred())
		Expect(files).To(Equal([]string{"solutions/m/m1/go.mod.txt", "solutions/m/m1/main.go.txt"}))
	})
})
