package scheme

import "fmt"

func init() {
	formats["vhs"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`Set Theme { "name": "%v", "black": "%v", "red": "%v", "green": "%v", "yellow": "%v", "blue": "%v", "magenta": "%v", "cyan": "%v", "white": "%v", "brightBlack": "%v", "brightRed": "%v", "brightGreen": "%v", "brightYellow": "%v", "brightBlue": "%v", "brightMagenta": "%v", "brightCyan": "%v", "brightWhite": "%v", "background": "%v", "foreground": "%v", "cursor": "%v" }`,
			t.Name,

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
