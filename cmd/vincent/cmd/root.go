package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/vincent/pkg/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vincent",
	Short: "terminal color theme chooser",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Execute()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd)
}
