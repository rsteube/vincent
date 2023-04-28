package scheme

import (
	"fmt"
	"strings"
)

func init() {
	formats["foot"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`[colors]
alpha=1.0

regular0=%v
regular1=%v
regular2=%v
regular3=%v
regular4=%v
regular5=%v
regular6=%v
regular7=%v

bright0=%v
bright1=%v
bright2=%v
bright3=%v
bright4=%v
bright5=%v
bright6=%v
bright7=%v

foreground=%v
background=%v`,

			strings.TrimPrefix(t.Black, "#"),
			strings.TrimPrefix(t.Red, "#"),
			strings.TrimPrefix(t.Green, "#"),
			strings.TrimPrefix(t.Yellow, "#"),
			strings.TrimPrefix(t.Blue, "#"),
			strings.TrimPrefix(t.Magenta, "#"),
			strings.TrimPrefix(t.Cyan, "#"),
			strings.TrimPrefix(t.White, "#"),
			strings.TrimPrefix(t.BrightBlack, "#"),
			strings.TrimPrefix(t.BrightRed, "#"),
			strings.TrimPrefix(t.BrightGreen, "#"),
			strings.TrimPrefix(t.BrightYellow, "#"),
			strings.TrimPrefix(t.BrightBlue, "#"),
			strings.TrimPrefix(t.BrightMagenta, "#"),
			strings.TrimPrefix(t.BrightCyan, "#"),
			strings.TrimPrefix(t.BrightWhite, "#"),

			strings.TrimPrefix(t.Foreground, "#"),
			strings.TrimPrefix(t.Background, "#"),
		), nil
	}
}
