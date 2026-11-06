package watcher_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/cleonte/golings/golings/watcher"
	. "github.com/onsi/gomega"
)

// quiet is long enough that a slow machine does not split five quick writes into two reports.
const quiet = 250 * time.Millisecond

// start watches a fresh folder holding exercises/a/main.go.
func start(t *testing.T) (root, file string, changes <-chan []string, stop context.CancelFunc) {
	t.Helper()
	root = filepath.Join(t.TempDir(), "exercises")
	dir := filepath.Join(root, "a")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	file = filepath.Join(dir, "main.go")
	write(t, file, "package main\n")

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	changes, err := watcher.Watch(ctx, root, quiet)
	if err != nil {
		t.Fatal(err)
	}
	return root, file, changes, cancel
}

func write(t *testing.T, name, content string) {
	t.Helper()
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func next(t *testing.T, changes <-chan []string) []string {
	t.Helper()
	select {
	case batch, ok := <-changes:
		if !ok {
			t.Fatal("the channel was closed")
		}
		return batch
	case <-time.After(5 * time.Second):
		t.Fatal("no change reported")
		return nil
	}
}

func expectQuiet(g *WithT, changes <-chan []string) {
	g.Consistently(changes, 5*quiet).ShouldNot(Receive())
}

func TestReportsAWrite(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)

	write(t, file, "package main // edited\n")

	g.Expect(next(t, changes)).To(Equal([]string{file}))
}

// A burst of writes is one report only if the watcher is not paused for longer
// than the quiet period between two of its events. On a busy machine (other
// test packages are running) that can happen, and the burst comes as two
// reports. So the check gets three tries: a watcher that reports every write
// fails all of them.
func TestOneSaveIsOneReport(t *testing.T) {
	var problem string
	for range 3 {
		if problem = oneSaveProblem(t); problem == "" {
			return
		}
	}
	t.Error(problem)
}

// oneSaveProblem writes five times to a file and says what is wrong with the
// reports that follow, or returns an empty string.
func oneSaveProblem(t *testing.T) string {
	t.Helper()
	_, file, changes, stop := start(t)
	defer stop()

	for i := 0; i < 5; i++ {
		write(t, file, "package main // edit\n")
	}

	var first []string
	select {
	case first = <-changes:
	case <-time.After(5 * time.Second):
		return "no change reported"
	}
	if len(first) != 1 || first[0] != file {
		return fmt.Sprintf("the first report was %v, want [%s]", first, file)
	}

	select {
	case second := <-changes:
		return fmt.Sprintf("a second report followed the first: %v", second)
	case <-time.After(5 * quiet):
		return ""
	}
}

func TestSeesEditorsThatSaveByRenaming(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)

	// vim and JetBrains write a temporary file and rename it over the original.
	temp := file + ".tmp"
	write(t, temp, "package main // atomic\n")
	if err := os.Rename(temp, file); err != nil {
		t.Fatal(err)
	}

	g.Expect(next(t, changes)).To(Equal([]string{file}))
}

func TestWatchesFoldersCreatedLater(t *testing.T) {
	g := NewWithT(t)
	root, _, changes, _ := start(t)

	newDir := filepath.Join(root, "b", "b1")
	if err := os.MkdirAll(newDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Give the watcher a moment to register the folders before writing in them.
	time.Sleep(3 * quiet)
	for len(changes) > 0 {
		<-changes
	}
	file := filepath.Join(newDir, "main.go")
	write(t, file, "package main\n")

	g.Eventually(func() []string {
		select {
		case batch := <-changes:
			return batch
		default:
			return nil
		}
	}, 5*time.Second, 20*time.Millisecond).Should(ContainElement(file))
}

func TestIgnoresFilesThatAreNotGoSources(t *testing.T) {
	g := NewWithT(t)
	_, file, changes, _ := start(t)
	dir := filepath.Dir(file)

	write(t, filepath.Join(dir, "notes.txt"), "x")
	write(t, filepath.Join(dir, ".main.go.swp"), "x")
	write(t, filepath.Join(dir, "main.go~"), "x")
	write(t, filepath.Join(dir, "#main.go#"), "x")

	expectQuiet(g, changes)
}

func TestClosesTheChannelWhenCancelled(t *testing.T) {
	_, _, changes, stop := start(t)

	stop()

	select {
	case _, ok := <-changes:
		if ok {
			t.Error("got a change after cancelling")
		}
	case <-time.After(5 * time.Second):
		t.Error("the channel was not closed")
	}
}

func TestFailsForAFolderThatDoesNotExist(t *testing.T) {
	g := NewWithT(t)

	_, err := watcher.Watch(context.Background(), filepath.Join(t.TempDir(), "nope"), quiet)

	g.Expect(err).To(HaveOccurred())
}

// Events that arrive while the folders are being added must not stop Watch (the
// Windows backend of fsnotify used to deadlock when nobody was reading them yet).
func TestWatchStartsWhileFilesAreBeingWritten(t *testing.T) {
	root := t.TempDir()
	for i := range 150 {
		if err := os.MkdirAll(filepath.Join(root, "topic", fmt.Sprintf("exercise%d", i)), 0o755); err != nil {
			t.Fatal(err)
		}
	}

	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Go(func() {
		for i := 0; ; i++ {
			select {
			case <-stop:
				return
			default:
				_ = os.WriteFile(filepath.Join(root, "topic", fmt.Sprintf("exercise%d", i%150), "main.go"), []byte("package main\n"), 0o644)
			}
		}
	})
	defer func() {
		close(stop)
		wg.Wait()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan error, 1)
	go func() {
		_, err := watcher.Watch(ctx, root, quiet)
		started <- err
	}()

	select {
	case err := <-started:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(20 * time.Second):
		t.Fatal("Watch did not return while files were being written")
	}
}
