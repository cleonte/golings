package cmd

import (
	"os"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func ListCmd(infoFile string) *cobra.Command {
	var asJSON bool

	listCmd := &cobra.Command{
		Use:           "list",
		Short:         "List all exercises",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if asJSON {
				list, err := listJSON(infoFile)
				if err != nil {
					return jsonError(cmd.OutOrStdout(), err)
				}
				return writeJSON(cmd.OutOrStdout(), list)
			}
			exs, err := exercises.List(infoFile)
			if err != nil {
				color.Red(err.Error())
				return err
			}
			ui.PrintList(os.Stdout, exs)
			printTopicProgress(infoFile)
			return nil
		},
	}
	listCmd.Flags().BoolVar(&asJSON, "json", false, jsonFlagHelp)

	return listCmd
}
