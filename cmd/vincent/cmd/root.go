package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/vincent"
	"github.com/rsteube/vincent/pkg/scheme"
	"github.com/rsteube/vincent/pkg/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vincent [scheme] [format]",
	Short: "terminal color scheme chooser",
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
					fmt.Fprintln(cmd.OutOrStdout(), strings.Join(scheme.Formats(), "\n"))

				case "full":
					for _, scheme := range vincent.Schemes() {
						if err := write(cmd.OutOrStdout(), scheme, "render"); err != nil {
							return err
						}
					}

				case "name":
					fmt.Fprintln(cmd.OutOrStdout(), strings.Join(vincent.Schemes(), "\n"))
				}
			} else {
				return ui.Execute()
			}

		case 1:
			return write(cmd.OutOrStdout(), args[0], "render")

		case 2:
			return write(cmd.OutOrStdout(), args[0], args[1])
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

func write(w io.Writer, scheme, format string) error {
	t, err := vincent.Load(scheme)
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

	rootCmd.Flags().StringP("list", "l", "", "list schemes")
	rootCmd.Flag("list").NoOptDefVal = "full"

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"list": carapace.ActionValues("full", "name", "formats"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if hasPathPrefix(c.Value) {
				return carapace.ActionFiles(".yml", ".yaml")
			}
			return carapace.ActionValues(vincent.Schemes()...).Tag("schemes")
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(scheme.Formats()...).Tag("formats")
		}),
	)
}

func hasPathPrefix(s string) bool { // TODO provide this in carapace (maybe add to context?)
	return strings.HasPrefix(s, ".") || strings.HasPrefix(s, "~") || strings.HasPrefix(s, "/")
}
