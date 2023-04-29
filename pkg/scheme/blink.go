package scheme

import (
	"fmt"
	"strings"
)

func init() {
	formats["blink"] = func(t Scheme) (string, error) {
		colors := t.Colors()
		for index, c := range t.Colors() {
			colors[index] = fmt.Sprintf(`"%v"`, c)
		}

		return fmt.Sprintf(`t.prefs_.set('color-palette-overrides', [%v]);
t.prefs_.set('background-color', "%v");
t.prefs_.set('foreground-color', "%v");
t.prefs_.set('cursor-color', "%v");`,
			strings.Join(colors, ", "),
			t.Background,
			t.Foreground,
			t.Cursor), nil
	}
}
