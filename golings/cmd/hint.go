package cmd

import (
	"fmt"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func hintLevelsText(n int) string {
	if n == 1 {
		return "1 hint level"
	}
	return fmt.Sprintf("%d hint levels", n)
}

func HintCmd(infoFile string) *cobra.Command {
	var level int

	command := &cobra.Command{
		Use:           "hint <exercise name>",
		Short:         "Get a hint for an exercise",
		Long:          "Get a hint for an exercise. Some exercises have smaller hints too: --level 1 is the smallest, and the highest level is the full hint, which is what you get without the flag.",
		Args:          cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var exercise exercises.Exercise
			var err error
			if args[0] == "next" {
				exercise, err = exercises.NextPending(infoFile)
			} else {
				exercise, err = exercises.Find(args[0], infoFile)
			}
			if err != nil {
				color.Red(err.Error())
				return err
			}

			levels := exercise.HintLevels()
			if level == 0 {
				color.Yellow(exercise.Hint)
				if len(levels) > 1 {
					color.White("There are smaller hints: golings hint %s --level 1 (up to %d)", exercise.Name, len(levels)-1)
				}
				return nil
			}
			if level < 1 || level > len(levels) {
				err := fmt.Errorf("%s has %s, not level %d", exercise.Name, hintLevelsText(len(levels)), level)
				color.Red(err.Error())
				return err
			}
			color.Yellow(levels[level-1])
			return nil
		},
	}
	command.Flags().IntVar(&level, "level", 0, "which hint to show: 1 is the smallest (default: the full hint)")
	return command
}
