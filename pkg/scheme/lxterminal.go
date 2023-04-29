package scheme

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	rgb := func(s string) string {
		c, err := colorful.Hex(s)
		if err != nil {
			return err.Error()
		}
		r, g, b := c.RGB255()
		return fmt.Sprintf("rgb(%v,%v,%v)", r, g, b)
	}
	formats["lxterminal"] = func(t Scheme) (string, error) {
		return fmt.Sprintf(`color_preset=%v
palette_color_0=%v
palette_color_1=%v
palette_color_2=%v
palette_color_3=%v
palette_color_4=%v
palette_color_5=%v
palette_color_6=%v
palette_color_7=%v
palette_color_8=%v
palette_color_9=%v
palette_color_10=%v
palette_color_11=%v
palette_color_12=%v
palette_color_13=%v
palette_color_14=%v
palette_color_15=%v
bgcolor=%v
fgcolor=%v`,
			t.Name,
			rgb(t.Black),
			rgb(t.Red),
			rgb(t.Green),
			rgb(t.Yellow),
			rgb(t.Blue),
			rgb(t.Magenta),
			rgb(t.Cyan),
			rgb(t.White),

			rgb(t.BrightBlack),
			rgb(t.BrightRed),
			rgb(t.BrightGreen),
			rgb(t.BrightYellow),
			rgb(t.BrightBlue),
			rgb(t.BrightMagenta),
			rgb(t.BrightCyan),
			rgb(t.BrightWhite),

			rgb(t.Background),
			rgb(t.Foreground)), nil
	}
}
