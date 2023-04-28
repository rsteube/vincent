package scheme

import (
	"fmt"
	"strings"
)

func init() {
	formats["terminator"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`[[%v]]
    palette = "%v"
    background_color = "%v"
    foreground_color = "%v"
    cursor_color = "%v"`,
			t.Name,
			strings.Join([]string{
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
				t.BrightWhite}, ":"),

			t.Background,
			t.Foreground,

			t.Cursor), nil
	}
}
