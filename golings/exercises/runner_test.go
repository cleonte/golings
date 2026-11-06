package exercises_test

import (
	"os"
	"path/filepath"

	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// tempDirHere makes a folder inside the working directory, removed after the
// test: the runner is given paths relative to it, and on Windows the temporary
// folder of the system can be on another drive, where no relative path exists.
func tempDirHere() string {
	dir, err := os.MkdirTemp(".", "tmp-")
	Expect(err).NotTo(HaveOccurred())
	DeferCleanup(os.RemoveAll, dir)
	abs, err := filepath.Abs(dir)
	Expect(err).NotTo(HaveOccurred())
	return abs
}

// writeExercise creates a throwaway exercise and returns it with a path
// relative to the working directory, which is how the runner is invoked.
func writeExercise(mode, name, source string) exercises.Exercise {
	dir := tempDirHere()
	file := filepath.Join(dir, name)
	Expect(os.WriteFile(file, []byte(source), 0o644)).To(Succeed())

	wd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	rel, err := filepath.Rel(wd, file)
	Expect(err).NotTo(HaveOccurred())

	return exercises.Exercise{Name: "tmp", Path: rel, Mode: mode}
}

// writeFolderExercise creates a throwaway exercise of several files, named by
// file name, and returns it with the first as its main file.
func writeFolderExercise(mode string, files map[string]string, main string) exercises.Exercise {
	dir := tempDirHere()
	for name, source := range files {
		Expect(os.WriteFile(filepath.Join(dir, name), []byte(source), 0o644)).To(Succeed())
	}
	wd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	rel, err := filepath.Rel(wd, filepath.Join(dir, main))
	Expect(err).NotTo(HaveOccurred())
	return exercises.Exercise{Name: "tmp", Path: rel, Mode: mode}
}

var _ = Describe("Running exercises", func() {
	Describe("an exercise of several files", func() {
		It("builds all the files of the folder as one package", func() {
			ex := writeFolderExercise("compile", map[string]string{
				"main.go":  "package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(greeting()) }\n",
				"greet.go": "package main\n\nfunc greeting() string { return \"from another file\" }\n",
			}, "main.go")

			result, err := ex.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Out).To(ContainSubstring("from another file"))
		})
		It("reports an error in a file that is not the main one", func() {
			ex := writeFolderExercise("compile", map[string]string{
				"main.go":  "package main\n\nfunc main() { greeting() }\n",
				"greet.go": "package main\n\nfunc greeting() { a := 1 }\n",
			}, "main.go")

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeTrue())
			Expect(result.Err).To(ContainSubstring("greet.go"))
		})
		It("runs the tests of a folder with the code in another file", func() {
			ex := writeFolderExercise("test", map[string]string{
				"x_test.go": "package x\n\nimport \"testing\"\n\nfunc TestIt(t *testing.T) {\n\tif double(2) != 4 {\n\t\tt.Fatal(\"wrong\")\n\t}\n}\n",
				"x.go":      "package x\n\nfunc double(n int) int { return n * 2 }\n",
			}, "x_test.go")

			_, err := ex.Run()

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Run", func() {
		It("captures the output of a compile-mode exercise that works", func() {
			ex := writeExercise("compile", "main.go", `package main

import "fmt"

func main() { fmt.Println("hello golings") }
`)

			result, err := ex.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Out).To(ContainSubstring("hello golings"))
			Expect(result.Exercise).To(Equal(ex))
		})
		It("reports the compiler error of a compile-mode exercise that is broken", func() {
			ex := writeExercise("compile", "main.go", `package main

func main() { a := 1 }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeTrue())
			Expect(result.Err).To(ContainSubstring("declared and not used"))
		})
		It("reports a test-mode exercise that does not compile as a build failure", func() {
			ex := writeExercise("test", "x_test.go", `package x

import "testing"

func TestBroken(t *testing.T) { undefinedThing() }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeTrue())
			Expect(result.Failure().Kind).To(Equal(exercises.CompileFailure))
		})
		It("reports a program that crashes while running as a runtime failure", func() {
			ex := writeExercise("compile", "main.go", `package main

func main() { panic("boom") }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeFalse())
			Expect(result.Failure().Kind).To(Equal(exercises.RuntimeFailure))
			Expect(result.Err).To(ContainSubstring("boom"))
		})
		It("passes a test-mode exercise whose tests pass", func() {
			ex := writeExercise("test", "x_test.go", `package x

import "testing"

func TestOK(t *testing.T) {}
`)

			result, err := ex.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Out).To(ContainSubstring("PASS"))
		})
		It("fails a test-mode exercise whose tests fail", func() {
			ex := writeExercise("test", "x_test.go", `package x

import "testing"

func TestBroken(t *testing.T) { t.Fatal("not yet") }
`)

			result, err := ex.Run()

			Expect(err).To(HaveOccurred())
			Expect(result.BuildFailed).To(BeFalse(), "it compiled, its tests failed")
			Expect(result.Out).To(ContainSubstring("not yet"))
		})
	})
})
