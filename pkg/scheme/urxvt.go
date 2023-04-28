package scheme

import "fmt"

func init() {
	formats["urxvt"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`URxvt.color0: %v
URxvt.color1: %v
URxvt.color2: %v
URxvt.color3: %v
URxvt.color4: %v
URxvt.color5: %v
URxvt.color6: %v
URxvt.color7: %v
URxvt.color8: %v
URxvt.color9: %v
URxvt.color10: %v
URxvt.color11: %v
URxvt.color12: %v
URxvt.color13: %v
URxvt.color14: %v
URxvt.color15: %v

URxvt.background: %v
URxvt.foreground: %v
URxvt.cursorColor: %v`,
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
