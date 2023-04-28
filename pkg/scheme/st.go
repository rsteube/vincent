package scheme

import "fmt"

func init() {
	outputs["st"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`/* Terminal colors (16 first used in escape sequence) */
static const char *colorname[] = {
	/* 8 normal colors */
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",

	/* 8 bright colors */
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",
	"%v",

[256] = "%v", /* default foreground colour */
[257] = "%v", /* default background colour */
[258] = "%v", /*575268*/

};


/*
 * foreground, background, cursor, reverse cursor
 */
unsigned int defaultfg = 256;
unsigned int defaultbg = 257;
unsigned int defaultcs = 258;
static unsigned int defaultrcs = 258;`,
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
		), nil
	}
}
