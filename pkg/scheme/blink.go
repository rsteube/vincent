package scheme

import (
	"fmt"
	"strings"
)

func init() {
	outputs["blink"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`t.prefs_.set('color-palette-overrides', [%v]);
t.prefs_.set('background-color', "%v");
t.prefs_.set('foreground-color', "%v");
t.prefs_.set('cursor-color', "%v");`,
			strings.Join([]string{
				`"` + t.Black + `"`,
				`"` + t.Red + `"`,
				`"` + t.Green + `"`,
				`"` + t.Yellow + `"`,
				`"` + t.Blue + `"`,
				`"` + t.Magenta + `"`,
				`"` + t.Cyan + `"`,
				`"` + t.White + `"`,

				`"` + t.BrightBlack + `"`,
				`"` + t.BrightRed + `"`,
				`"` + t.BrightGreen + `"`,
				`"` + t.BrightYellow + `"`,
				`"` + t.BrightBlue + `"`,
				`"` + t.BrightMagenta + `"`,
				`"` + t.BrightCyan + `"`,
				`"` + t.BrightWhite + `"`}, ", "),

			t.Background,
			t.Foreground,

			t.Cursor), nil
	}
}
