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
			strings.Join(t.Colors(), ":"),

			t.Background,
			t.Foreground,

			t.Cursor), nil
	}
}
