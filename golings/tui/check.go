package tui

import (
	"context"
	"errors"
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/lingscore/tui/checkall"
)

// checkWorkers is how many exercises are built and run at the same time while
// checking all of them.
const checkWorkers = 4

// next moves on from an exercise that works: it takes out the "I AM NOT DONE"
// marker, which is what the learner would do by hand, and runs the next one.
func (m *Model) next() tea.Cmd {
	switch {
	case !m.hasExercise():
		return nil
	case m.status == statusFinished && exercises.IsSkipped(m.runErr):
		// Only an exercise picked from the list can be a skipped one. It is not
		// done, so its marker stays; go back to the first pending exercise.
		name := m.exercise.Name
		m.pinned = ""
		cmd := m.startRun()
		m.notice = name + " is skipped here, so it is not done."
		return cmd
	case m.status != statusFinished || m.runErr != nil:
		m.notice = "Get " + m.exercise.Name + " to work first, then n moves on to the next exercise."
		return nil
	}

	if err := m.exercise.MarkDone(); err != nil && !errors.Is(err, exercises.ErrNotMarked) {
		m.notice = "Could not mark " + m.exercise.Name + " as done: " + err.Error()
		return nil
	}

	name := m.exercise.Name
	m.pinned = ""
	cmd := m.startRun()
	m.notice = name + " is done."
	return cmd
}

// startCheck runs every exercise, a few at a time, without changing any file.
func (m *Model) startCheck() tea.Cmd {
	if m.check.Active() {
		m.notice = "Already checking the exercises."
		return nil
	}

	all, err := exercises.List(m.cfg.InfoFile)
	if err != nil || len(all) == 0 {
		m.notice = "There are no exercises to check."
		return nil
	}

	m.stopRun()
	if m.status == statusRunning {
		m.status = statusStarting
	}
	m.mode, m.scroll = modeCheck, 0

	run := m.cfg.Run
	return m.check.Start(all, checkWorkers, func(ctx context.Context, e exercises.Exercise) checkall.Result {
		return checkResult(e, run)(ctx)
	})
}

// checkResult runs one exercise and says what the view shows about it: how it
// failed, or why it was skipped.
func checkResult(e exercises.Exercise, run RunFunc) func(context.Context) checkall.Result {
	return func(ctx context.Context) checkall.Result {
		result, err := run(ctx, e)
		switch {
		case exercises.IsSkipped(err):
			return checkall.Result{Status: checkall.Skipped, Detail: skipReason(e, err)}
		case err != nil:
			return checkall.Result{Status: checkall.Failed, Detail: result.Failure().Kind.String()}
		}
		return checkall.Result{Status: checkall.Passed}
	}
}

func (m *Model) checkLines() []string {
	lines := []string{titleStyle.Render(m.check.Title()), ""}
	for _, item := range m.check.Items {
		lines = append(lines, checkLine(item))
	}
	return lines
}

// skipReason says why an exercise was skipped, from the exercise itself or, if
// that can't tell (the local Go changed since), from the error of its run.
func skipReason(e exercises.Exercise, err error) string {
	if reason := e.SkipReason(); reason != "" {
		return reason
	}
	if errors.Is(err, exercises.ErrRaceUnsupported) {
		return "needs the race detector"
	}
	return "needs a newer Go"
}

func checkLine(item checkall.Item[exercises.Exercise]) string {
	name := item.Value.Name
	switch item.Status {
	case checkall.Waiting:
		return dimStyle.Render("· " + name)
	case checkall.Running:
		return dimStyle.Render("… " + name)
	case checkall.Skipped:
		return warnStyle.Render(fmt.Sprintf("• %s  %s", name, item.Detail))
	case checkall.Passed:
		if item.Value.State() == exercises.Pending {
			return okStyle.Render("✔ "+name) + dimStyle.Render("  works, the I AM NOT DONE marker is still there")
		}
		return okStyle.Render("✔ " + name)
	}

	// A failure: say how it failed, and shout if it was finished before.
	if item.Value.State() == exercises.Done {
		return warnStyle.Render(fmt.Sprintf("✗ %s  was finished, but now %s", name, item.Detail))
	}
	return errStyle.Render("✗ "+name) + dimStyle.Render("  "+item.Detail)
}
