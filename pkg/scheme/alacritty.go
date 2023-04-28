package scheme

import "fmt"

func init() {
	formats["alacritty"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`colors:
  normal:
    black:   '%v'
    red:     '%v'
    green:   '%v'
    yellow:  '%v'
    blue:    '%v'
    magenta: '%v'
    cyan:    '%v'
    white:   '%v'
  
  bright:
    black:   '%v'
    red:     '%v'
    green:   '%v'
    yellow:  '%v'
    blue:    '%v'
    magenta: '%v'
    cyan:    '%v'
    white:   '%v'

  primary:
    background: '%v'
    foreground: '%v'
  
  cursor:
    cursor: '%v'`,
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

			t.Background,
			t.Foreground,

			t.Cursor), nil
	}
}
