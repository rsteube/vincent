package scheme

import (
	"fmt"
	"strings"
)

func init() {
	formats["xfce"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`[Scheme]
Name=%v
ColorPalette=%v
ColorForeground=%v
ColorBackground=%v
#TabActivityColor=#FAB387
ColorSelectionUseDefault=TRUE
ColorCursorUseDefault=FALSE
#ColorCursorForeground=#11111B
ColorCursor=%v`,
			t.Name,
			strings.Join(t.Colors(), ";"),

			t.Background,
			t.Foreground,

			t.Cursor), nil
	}
}
