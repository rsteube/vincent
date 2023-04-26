package cmd

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/vincent"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [prefix]",
	Short: "view themes",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, theme := range vincent.Themes() {
			t, err := vincent.Load(theme)
			if err != nil {
				return err
			}
			if len(args) == 0 || strings.HasPrefix(t.Name, args[0]) {
				fmt.Fprintln(cmd.OutOrStderr(), t.Name)
				fmt.Fprintln(cmd.OutOrStderr(), t.Render())
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	carapace.Gen(viewCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(vincent.Themes()...)
		}),
	)
}
