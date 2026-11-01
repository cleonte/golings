package ui

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
)

const tabWidth = 4

// PrintFailure explains why an exercise failed. Compiler errors are shown with
// the offending source line and a caret under the column.
func PrintFailure(w io.Writer, ex exercises.Exercise, f exercises.Failure) {
	cyan, red, white := color.New(color.FgCyan), color.New(color.FgRed), color.New(color.FgWhite)

	switch f.Kind {
	case exercises.CompileFailure:
		cyan.Fprintf(w, "Failed to compile the exercise %s\n\n", ex.Path)
	case exercises.TestFailure:
		cyan.Fprintf(w, "Tests failed for the exercise %s\n\n", ex.Path)
	default:
		cyan.Fprintf(w, "The exercise %s failed while running\n\n", ex.Path)
	}

	for _, d := range f.Diagnostics {
		location := fmt.Sprintf("%s:%d", d.File, d.Line)
		if d.Col > 0 {
			location += fmt.Sprintf(":%d", d.Col)
		}
		red.Fprintf(w, "%s: %s\n", location, d.Message)
		printSnippet(w, d)
		fmt.Fprintln(w)
	}

	if f.Output != "" {
		if f.Kind != exercises.CompileFailure {
			white.Fprintln(w, "Output:")
			fmt.Fprintln(w)
		}
		red.Fprintln(w, f.Output)
		fmt.Fprintln(w)
	}
}

func printSnippet(w io.Writer, d exercises.Diagnostic) {
	data, err := os.ReadFile(d.File)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	if d.Line < 1 || d.Line > len(lines) {
		return
	}

	source := strings.TrimRight(lines[d.Line-1], "\r")
	gutter := fmt.Sprintf("%d", d.Line)
	pad := strings.Repeat(" ", len(gutter))

	fmt.Fprintf(w, "\n  %s | %s\n", gutter, expandTabs(source))
	if d.Col > 0 {
		prefix := source
		if d.Col-1 < len(prefix) {
			prefix = prefix[:d.Col-1]
		}
		// Columns count bytes, the caret has to line up with what is displayed.
		offset := len([]rune(expandTabs(prefix)))
		fmt.Fprintf(w, "  %s | %s^\n", pad, strings.Repeat(" ", offset))
	}
}

func expandTabs(s string) string {
	return strings.ReplaceAll(s, "\t", strings.Repeat(" ", tabWidth))
}
