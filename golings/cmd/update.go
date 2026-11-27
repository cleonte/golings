package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/cleonte/golings"
	"github.com/cleonte/golings/golings/workspace"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// UpdateCmd brings an existing exercises folder up to date with the exercises of
// this version of golings.
func UpdateCmd(src fs.FS, root string) *cobra.Command {
	var dryRun bool
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Add the exercises of this version that your folder does not have yet",
		Long: `A new version of golings can bring new exercises, and your exercises folder, made
by an older version, does not know about them. update adds them to the folder and to
info.toml, each one where it belongs in the order of the course.

It never writes over a file of an exercise you already have, because that may be
your work. It keeps a copy of info.toml as info.toml.bak. The README of a topic is
replaced by its newer text. To get the newer version of an exercise you have,
use 'golings reset <name>', which loses your changes to it.

Use --dry-run to see what it would do.`,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := workspace.Update(src, root, dryRun)
			if err != nil {
				color.Red("Cannot update: %v (run it from the folder with info.toml)", err)
				return err
			}
			printUpdate(res, dryRun)
			return nil
		},
	}
	updateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "only say what would be added or replaced")
	return updateCmd
}

func printUpdate(res workspace.UpdateResult, dryRun bool) {
	if len(res.Added) == 0 && res.Readmes == 0 {
		color.Green("Your exercises are up to date with this version of golings.")
		return
	}
	if len(res.Added) > 0 {
		verb := "Added"
		if dryRun {
			verb = "Would add"
		}
		topic := ""
		color.Green("%s %d new exercises:", verb, len(res.Added))
		for _, ex := range res.Added {
			if ex.Topic() != topic {
				topic = ex.Topic()
				color.Cyan("\n  %s", topic)
			}
			_, _ = fmt.Fprintf(color.Output, "    %-24s %s\n", ex.Name, ex.Mode)
		}
		_, _ = fmt.Fprintln(color.Output)
		if !dryRun {
			color.White("They are in the order of the course. Your earlier info.toml is info.toml.bak.")
		}
	}
	if res.Readmes > 0 {
		verb := "Replaced"
		if dryRun {
			verb = "Would replace"
		}
		_, _ = fmt.Fprintf(color.Output, "%s %d README files with their newer text.\n", verb, res.Readmes)
	}
}

// newExercises is how many exercises this version of golings has that the
// folder of the info file does not list yet.
func newExercises(infoFile string) int {
	return len(workspace.Missing(golings.FS, filepath.Dir(infoFile)))
}

// noticeNewExercises prints one line, in yellow, when the folder is behind this
// version of golings. It is shown by run, list and watch.
func noticeNewExercises(infoFile string) {
	if n := newExercises(infoFile); n > 0 {
		color.Yellow("This version of golings has %d exercises that your folder does not have yet: run `golings update`.", n)
	}
}
