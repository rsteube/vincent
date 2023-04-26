package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/rsteube/carapace"
	"github.com/rsteube/vincent"
	"github.com/rsteube/vincent/pkg/ui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "vincent",
	Short: "terminal color theme chooser",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			if cmd.Flag("list").Changed {
				for _, theme := range vincent.Themes() {
					print(cmd.OutOrStdout(), theme, "render")
				}
			} else {
				ui.Execute()
			}

		case 1:
			print(cmd.OutOrStdout(), args[0], "render")

		case 2:
			print(cmd.OutOrStdout(), args[0], args[1])
		}

		return nil
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func print(w io.Writer, theme, format string) error {
	t, err := vincent.Load(theme)
	if err != nil {
		return err
	}

	switch format {
	case "render":
		fmt.Fprintln(w, t.Name)
		fmt.Fprintln(w, t.Render())
	case "json":
		m, err := json.Marshal(t)
		if err != nil {
			return err
		}
		fmt.Println(w, string(m))

	case "yaml":
		m, err := yaml.Marshal(t)
		if err != nil {
			return err
		}
		fmt.Println(w, string(m))
	}
	return nil
}

func init() {
	rootCmd.Flags().BoolP("list", "l", false, "list themes")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(vincent.Themes()...)
		}),
		carapace.ActionValues("json", "yaml"),
	)
}
