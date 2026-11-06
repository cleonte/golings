package cmd

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var errSolutionHidden = errors.New("solution not shown")

// SolutionCmd prints the solution of an exercise from src. Solutions are shown
// once the learner has finished the exercise, or on request with --reveal.
func SolutionCmd(infoFile string, src fs.FS) *cobra.Command {
	var reveal bool

	solutionCmd := &cobra.Command{
		Use:   "solution <exercise name>",
		Short: "Show the solution of an exercise",
		Long: `Prints the solution of an exercise, to compare with your own once you are done.

While the exercise is still pending (it still has its "I AM NOT DONE" marker)
the solution stays hidden unless you pass --reveal.`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			exercise, err := exercises.Find(args[0], infoFile)
			if errors.Is(err, exercises.ErrExerciseNotFound) {
				color.White("No exercise found for '%s'", args[0])
				return err
			}
			if err != nil {
				color.Red(err.Error())
				return err
			}

			if exercise.State() == exercises.Pending && !reveal {
				color.Yellow("%s is not done yet, so its solution stays hidden.", exercise.Name)
				color.White("Finish it and remove the 'I AM NOT DONE' marker, or use --reveal to look anyway.")
				return errSolutionHidden
			}

			files, err := exercise.SolutionFiles(src)
			if err != nil || len(files) == 0 {
				color.Red("No solution available for %s yet", exercise.Name)
				if err == nil {
					err = fs.ErrNotExist
				}
				return err
			}

			// A solution may have several files: the main one first, each with its name.
			for i, file := range files {
				source, err := fs.ReadFile(src, file)
				if err != nil {
					color.Red("No solution available for %s yet", exercise.Name)
					return err
				}
				if i == 0 {
					color.Cyan("Solution for %s (%s):\n\n", exercise.Name, file)
				} else {
					color.Cyan("\n%s:\n\n", file)
				}
				// Printed as is: the source is not a format string, it may well contain %d.
				_, _ = fmt.Fprint(color.Output, string(source))
			}
			return nil
		},
	}
	solutionCmd.Flags().BoolVar(&reveal, "reveal", false, "show the solution even if the exercise is not done")

	return solutionCmd
}
