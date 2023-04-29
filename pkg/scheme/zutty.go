package scheme

import "fmt"

func init() {
	formats["zutty"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`Zutty.color0: %v
Zutty.color1: %v
Zutty.color2: %v
Zutty.color3: %v
Zutty.color4: %v
Zutty.color5: %v
Zutty.color6: %v
Zutty.color7: %v
Zutty.color8: %v
Zutty.color9: %v
Zutty.color10: %v
Zutty.color11: %v
Zutty.color12: %v
Zutty.color13: %v
Zutty.color14: %v
Zutty.color15: %v

Zutty.bg: %v
Zutty.fg: %v
Zutty.cr: %v`,
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
