package theme

import (
	"fmt"
)

func init() {
	outputs["kitty"] = func(t Theme) string {
		return fmt.Sprintf(`color0                %v
color1                %v
color2                %v
color3                %v
color4                %v
color5                %v
color6                %v
color7                %v
color8                %v
color9                %v
color10               %v
color11               %v
color12               %v
color13               %v
color14               %v
color15               %v

foreground            %v
background            %v
cursor                %v`,

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

			t.Foreground,
			t.Background,
			t.Cursor,
		)
	}
}
