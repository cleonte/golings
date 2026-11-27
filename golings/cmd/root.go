// Package cmd has the commands of golings: the cobra commands and what they print.
package cmd

import (
	"os"

	"github.com/cleonte/golings"
	"github.com/spf13/cobra"
)

// NewRootCmd is the golings command with all of its subcommands.
func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "golings",
		Short:         "Learn go through interactive exercises",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}

	rootCmd.AddCommand(DoctorCmd("info.toml"))
	rootCmd.AddCommand(HintCmd("info.toml"))
	rootCmd.AddCommand(InitCmd(golings.FS))
	rootCmd.AddCommand(ResetCmd(golings.FS, "."))
	rootCmd.AddCommand(SolutionCmd("info.toml", golings.Solutions))
	rootCmd.AddCommand(ListCmd("info.toml"))
	rootCmd.AddCommand(RunCmd("info.toml"))
	rootCmd.AddCommand(VerifyCmd("info.toml"))
	rootCmd.AddCommand(UpdateCmd(golings.FS, "."))
	rootCmd.AddCommand(WatchCmd("info.toml"))

	return rootCmd
}

// Execute runs the golings command and exits with 1 when it fails.
func Execute(version string) {
	if err := NewRootCmd(version).Execute(); err != nil {
		os.Exit(1)
	}
}
