package tui

import (
	"errors"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/gomega"
)

// raceOutput is what a test prints for a data race in the exercise file.
func raceOutput(path string) string {
	return strings.Join([]string{
		"=== RUN   TestCounter",
		"==================",
		"WARNING: DATA RACE",
		"Read at 0x00c0000122c8 by goroutine 12:",
		"  command-line-arguments_test.updateCounter.func1()",
		"      " + path + ":28 +0x68",
		"",
		"Previous write at 0x00c0000122c8 by goroutine 8:",
		"  command-line-arguments_test.updateCounter.func1()",
		"      " + path + ":28 +0x78",
		"",
		"Goroutine 12 (running) created at:",
		"  command-line-arguments_test.updateCounter()",
		"      " + path + ":25 +0x6c",
		"==================",
		"    testing.go:1867: race detected during execution of test",
		"--- FAIL: TestCounter (0.00s)",
	}, "\n")
}

// racing starts a model whose exercise fails with a data race report.
func racing(t *testing.T) *Model {
	t.Helper()
	noColor(t)
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) {
		return exercises.Result{Exercise: ex, Out: raceOutput(ex.Path)}, errors.New("exit status 1")
	}}
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, "!a/a1"), Run: fake.run})
	runCmds(t, m, cmd)
	return m
}

func TestAShortenedFailureSaysHowToSeeAllOfIt(t *testing.T) {
	g := NewWithT(t)
	m := racing(t)

	view := plain(m.View().Content)

	g.Expect(view).To(ContainSubstring("Data race: two goroutines use the same variable without synchronization"))
	g.Expect(view).To(ContainSubstring("main.go:28  in updateCounter.func1"))
	g.Expect(view).NotTo(ContainSubstring("WARNING: DATA RACE"))
	g.Expect(view).NotTo(ContainSubstring("0x00c0"))
	g.Expect(view).To(ContainSubstring("Some of the output was shortened: press f to see all of it."))
	g.Expect(view).To(ContainSubstring("c check all   f full   q quit"), "the footer shows the key (the scroll hint makes way when the line is too long)")
}

func TestFShowsTheWholeOutputAndGoesBack(t *testing.T) {
	g := NewWithT(t)
	m := racing(t)

	press(m, "f")
	view := plain(m.View().Content)

	g.Expect(m.mode).To(Equal(modeFull))
	g.Expect(view).To(ContainSubstring("WARNING: DATA RACE"))
	g.Expect(view).To(ContainSubstring("0x00c0000122c8"))
	g.Expect(view).To(ContainSubstring("Goroutine 12 (running) created at:"))
	g.Expect(view).NotTo(ContainSubstring("Data race: two goroutines"))
	g.Expect(view).To(ContainSubstring("esc back   f shortened"))

	press(m, "f")
	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(plain(m.View().Content)).To(ContainSubstring("Data race: two goroutines"))

	press(m, "f", "esc")
	g.Expect(m.mode).To(Equal(modeResult), "esc goes back too")
}

func TestFScrolls(t *testing.T) {
	g := NewWithT(t)
	m := racing(t)
	m.Update(tea.WindowSizeMsg{Width: 60, Height: 12})

	press(m, "f")
	g.Expect(plain(m.View().Content)).NotTo(ContainSubstring("Goroutine 12 (running) created at:"))
	for range 30 {
		press(m, "down")
	}

	g.Expect(plain(m.View().Content)).To(ContainSubstring("Goroutine 12 (running) created at:"))
}

func TestFWhenNothingWasShortened(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	fake := &fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) { return failingToCompile(ex) }}
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, "!a/a1"), Run: fake.run})
	runCmds(t, m, cmd)

	g.Expect(plain(m.View().Content)).NotTo(ContainSubstring("f full"))
	press(m, "f")

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(plain(m.View().Content)).To(ContainSubstring("Nothing was shortened: this is all of the output."))
}

func TestFWhenTheExerciseWorks(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	m, cmd := start(t, Config{InfoFile: newWorkspace(t, "!a/a1"), Run: (&fakeRuns{result: func(ex exercises.Exercise) (exercises.Result, error) { return passing(ex) }}).run})
	runCmds(t, m, cmd)

	press(m, "f")

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(m.notice).To(BeEmpty())
}

func TestANewRunLeavesTheFullOutput(t *testing.T) {
	g := NewWithT(t)
	m := racing(t)
	press(m, "f")

	_, cmd := m.Update(changesMsg{files: []string{"a1"}})
	runCmds(t, m, cmd)

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(plain(m.View().Content)).To(ContainSubstring("Data race: two goroutines"))
}

func TestFWhileTheRunIsStillGoing(t *testing.T) {
	g := NewWithT(t)
	noColor(t)
	m, _ := start(t, Config{InfoFile: newWorkspace(t, "!a/a1"), Run: (&fakeRuns{block: true}).run})

	press(m, "f")

	g.Expect(m.mode).To(Equal(modeResult))
	g.Expect(m.fullAvailable).To(BeFalse())
}
