package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/vincent"
	"github.com/rsteube/vincent/pkg/theme"
	"github.com/rsteube/vincent/pkg/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vincent [theme] [format]",
	Short: "terminal color theme chooser",
	Example: `  vincent completion:
    bash:       source <(vincent _carapace bash)
    elvish:     eval (vincent _carapace elvish | slurp)
    fish:       vincent _carapace fish | source
    nushell:    vincent _carapace nushell
    oil:        source <(vincent _carapace oil)
    powershell: vincent _carapace powershell | Out-String | Invoke-Expression
    tcsh:       eval ` + "`" + `vincent _carapace tcsh` + "`" + `
    xonsh:      exec($(vincent _carapace xonsh))
    zsh:        source <(vincent _carapace zsh)
`,
	Args: cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			if f := cmd.Flag("list"); f.Changed {
				switch f.Value.String() {
				case "formats":
					fmt.Fprintln(cmd.OutOrStdout(), strings.Join(theme.Formats(), "\n"))

				case "full":
					for _, theme := range vincent.Themes() {
						if err := print(cmd.OutOrStdout(), theme, "render"); err != nil {
							return err
						}
					}

				case "name":
					fmt.Fprintln(cmd.OutOrStdout(), strings.Join(vincent.Themes(), "\n"))
				}
			} else {
				ui.Execute()
			}

		case 1:
			return print(cmd.OutOrStdout(), args[0], "render")

		case 2:
			return print(cmd.OutOrStdout(), args[0], args[1])
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

	f, err := t.Format(format)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, f)
	return nil
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.Flags().StringP("list", "l", "", "list themes")
	rootCmd.Flag("list").NoOptDefVal = "full"

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"list": carapace.ActionValues("full", "name", "formats"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(vincent.Themes()...)
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(theme.Formats()...)
		}),
	)
}
