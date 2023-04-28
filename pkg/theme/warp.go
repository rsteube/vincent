package theme

import (
	"fmt"
)

func init() {
	outputs["warp"] = func(t Theme) (string, error) {
		details := "lighter"
		if t.IsDark() {
			details = "darker"
		}

		return fmt.Sprintf(`background: '%v'
accent: '%v'
foreground: '%v'
details: %v
terminal_colors:
  normal:
    black: '%v'
    red: '%v'
    green: '%v'
    yellow: '%v'
    blue: '%v'
    magenta: '%v'
    cyan: '%v'
    white: '%v'
  bright:
    black: '%v'
    red: '%v'
    green: '%v'
    yellow: '%v'
    blue: '%v'
    magenta: '%v'
    cyan: '%v'
    white: '%v'`,
			t.Background,
			t.Cursor,
			t.Foreground,
			details,

			t.Black,
			t.Red,
			t.Green,
			t.Yellow,
			t.Blue,
			t.Magenta,
			t.Cyan,
			t.White,

			t.BrightBlack,
			t.BrightRed,
			t.BrightGreen,
			t.BrightYellow,
			t.BrightBlue,
			t.BrightMagenta,
			t.BrightCyan,
			t.BrightWhite,
		), nil
	}
}
