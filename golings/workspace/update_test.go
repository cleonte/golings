package workspace_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/cleonte/golings"
	"github.com/cleonte/golings/golings/workspace"
	. "github.com/onsi/gomega"
)

func entry(name, topic string) string {
	return "[[exercises]]\nname = \"" + name + "\"\npath = \"exercises/" + topic + "/" + name + "/main.go\"\nmode = \"compile\"\nhint = \"\"\n\n"
}

// release is a source with the exercises of the given names (all in topic a,
// except the ones starting with b, which are in topic b), in that order.
func release(names ...string) fstest.MapFS {
	info := ""
	src := fstest.MapFS{}
	for _, n := range names {
		topic := "a"
		if strings.HasPrefix(n, "b") {
			topic = "b"
		}
		info += entry(n, topic)
		src["exercises/"+topic+"/"+n+"/main.go"] = &fstest.MapFile{Data: []byte(unsolved)}
		src["exercises/"+topic+"/README.md"] = &fstest.MapFile{Data: []byte("# " + topic)}
	}
	src["info.toml"] = &fstest.MapFile{Data: []byte(info)}
	return src
}

func read(g *WithT, root, name string) string {
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
	g.Expect(err).NotTo(HaveOccurred(), name)
	return string(data)
}

func TestMissingListsTheNewExercisesInOrder(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1", "a3"), root)).To(Succeed())

	g.Expect(workspace.Missing(release("a1", "a2", "a3", "a4"), root)).To(Equal([]string{"a2", "a4"}))
	g.Expect(workspace.Missing(release("a1", "a3"), root)).To(BeEmpty())
	g.Expect(workspace.Missing(release("a1"), filepath.Join(t.TempDir(), "nothing"))).To(BeEmpty(), "a folder without info.toml has nothing to compare")
}

func TestUpdateAddsNewExercisesWhereTheyBelongAndKeepsTheLearnersWork(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1", "a3"), root)).To(Succeed())
	work := "package main\n\nfunc main() { println(1) }\n"
	g.Expect(os.WriteFile(filepath.Join(root, "exercises", "a", "a1", "main.go"), []byte(work), 0o644)).To(Succeed())
	before := read(g, root, "info.toml")

	res, err := workspace.Update(release("a0", "a1", "a2", "a3", "a4"), root, false)

	g.Expect(err).NotTo(HaveOccurred())
	var added []string
	for _, e := range res.Added {
		added = append(added, e.Name)
	}
	g.Expect(added).To(Equal([]string{"a0", "a2", "a4"}))
	g.Expect(read(g, root, "exercises/a/a1/main.go")).To(Equal(work), "the work of the learner is not touched")
	for _, n := range added {
		g.Expect(read(g, root, "exercises/a/"+n+"/main.go")).To(Equal(unsolved))
	}
	info := read(g, root, "info.toml")
	g.Expect(info).To(Equal(string(release("a0", "a1", "a2", "a3", "a4")["info.toml"].Data)), "each entry is where it is in the course")
	g.Expect(read(g, root, "info.toml.bak")).To(Equal(before))
	g.Expect(workspace.Missing(release("a0", "a1", "a2", "a3", "a4"), root)).To(BeEmpty())
}

func TestUpdateKeepsEntriesOfTheFolderTheBinaryDoesNotKnow(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1", "old", "a3"), root)).To(Succeed())

	_, err := workspace.Update(release("a1", "a2", "a3"), root, false)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(read(g, root, "info.toml")).To(Equal(string(release("a1", "a2", "old", "a3")["info.toml"].Data)))
}

func TestUpdateDryRunChangesNothing(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1"), root)).To(Succeed())
	g.Expect(os.WriteFile(filepath.Join(root, "exercises", "a", "README.md"), []byte("old"), 0o644)).To(Succeed())

	res, err := workspace.Update(release("a1", "a2"), root, true)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Added).To(HaveLen(1))
	g.Expect(res.Readmes).To(Equal(1))
	g.Expect(read(g, root, "exercises/a/README.md")).To(Equal("old"))
	g.Expect(filepath.Join(root, "exercises", "a", "a2")).NotTo(BeADirectory())
	g.Expect(filepath.Join(root, "info.toml.bak")).NotTo(BeAnExistingFile())
}

func TestUpdateReplacesTheReadmeOfATopicAndBringsTheOneOfANewTopic(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1"), root)).To(Succeed())
	g.Expect(os.WriteFile(filepath.Join(root, "exercises", "a", "README.md"), []byte("old"), 0o644)).To(Succeed())

	res, err := workspace.Update(release("a1", "b1"), root, false)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Readmes).To(Equal(1), "the new topic is part of the added exercise, not a replacement")
	g.Expect(read(g, root, "exercises/a/README.md")).To(Equal("# a"))
	g.Expect(read(g, root, "exercises/b/README.md")).To(Equal("# b"))
}

func TestUpdateWritesAModuleExerciseWithoutTheSuffix(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1"), root)).To(Succeed())

	res, err := workspace.Update(moduleFS(), root, false)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Added).To(HaveLen(1))
	for _, name := range []string{"exercises/m/m1/go.mod", "exercises/m/m1/main.go", "exercises/m/m1/greet/go.mod"} {
		g.Expect(filepath.Join(root, filepath.FromSlash(name))).To(BeAnExistingFile(), name)
	}
	g.Expect(filepath.Join(root, "exercises", "m", "m1", "go.mod.txt")).NotTo(BeAnExistingFile())
}

func TestUpdateFailsOutsideAWorkspace(t *testing.T) {
	g := NewWithT(t)

	_, err := workspace.Update(release("a1"), t.TempDir(), false)

	g.Expect(err).To(HaveOccurred())
}

func TestUpdateWithTheRealExercisesIsANoOpOnAFreshFolder(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(golings.FS, root)).To(Succeed())

	res, err := workspace.Update(golings.FS, root, true)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Added).To(BeEmpty())
	g.Expect(res.Readmes).To(BeZero())
}

// releaseWith is a release whose exercise a1 has the given content.
func releaseWith(content string, names ...string) fstest.MapFS {
	src := release(names...)
	src["exercises/a/a1/main.go"] = &fstest.MapFile{Data: []byte(content)}
	return src
}

func TestInitRecordsTheOriginalsAndAnUntouchedFileIsRefreshed(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1"), root)).To(Succeed())
	_, err := os.Stat(filepath.Join(root, ".golings.json"))
	g.Expect(err).NotTo(HaveOccurred(), "init writes the progress file")

	newer := "package main\n\n// I AM NOT DONE\n// a better exercise\nfunc main() {}\n"
	dry, err := workspace.Update(releaseWith(newer, "a1"), root, true)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(dry.Refreshed).To(Equal([]string{"a1"}))
	g.Expect(read(g, root, "exercises/a/a1/main.go")).To(Equal(unsolved), "a dry run writes nothing")

	res, err := workspace.Update(releaseWith(newer, "a1"), root, false)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Refreshed).To(Equal([]string{"a1"}))
	g.Expect(res.Newer).To(BeEmpty())
	g.Expect(read(g, root, "exercises/a/a1/main.go")).To(Equal(newer))

	again, err := workspace.Update(releaseWith(newer, "a1"), root, false)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(again.Refreshed).To(BeEmpty(), "the new version is the recorded one now")
}

func TestUpdateKeepsAFileTheLearnerChangedAndSaysThereIsANewerOne(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1"), root)).To(Succeed())
	work := "package main\n\nfunc main() { println(1) }\n"
	g.Expect(os.WriteFile(filepath.Join(root, "exercises", "a", "a1", "main.go"), []byte(work), 0o644)).To(Succeed())

	res, err := workspace.Update(releaseWith("package main\n// newer\n", "a1"), root, false)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Newer).To(Equal([]string{"a1"}))
	g.Expect(res.Refreshed).To(BeEmpty())
	g.Expect(read(g, root, "exercises/a/a1/main.go")).To(Equal(work))
}

func TestAFolderMadeBeforeTheProgressFileIsOnlyRecordedNotRefreshed(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	g.Expect(workspace.Init(release("a1"), root)).To(Succeed())
	g.Expect(os.Remove(filepath.Join(root, ".golings.json"))).To(Succeed()) // as made by an older golings

	res, err := workspace.Update(release("a1"), root, false)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Refreshed).To(BeEmpty())
	_, err = os.Stat(filepath.Join(root, ".golings.json"))
	g.Expect(err).NotTo(HaveOccurred(), "an untouched, current file is recorded so that a later release can refresh it")

	newer := "package main\n// newer\n"
	res, err = workspace.Update(releaseWith(newer, "a1"), root, false)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Refreshed).To(Equal([]string{"a1"}))
}

func TestResetRecordsTheOriginalAgain(t *testing.T) {
	g := NewWithT(t)
	root := filepath.Join(t.TempDir(), "ws")
	src := release("a1")
	g.Expect(workspace.Init(src, root)).To(Succeed())
	g.Expect(os.WriteFile(filepath.Join(root, "exercises", "a", "a1", "main.go"), []byte("package main\n"), 0o644)).To(Succeed())
	_, err := workspace.Reset(src, root, "a1")
	g.Expect(err).NotTo(HaveOccurred())

	res, err := workspace.Update(releaseWith("package main\n// newer\n", "a1"), root, false)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Refreshed).To(Equal([]string{"a1"}), "a reset file is untouched again")
}
