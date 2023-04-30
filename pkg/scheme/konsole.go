package scheme

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	formats["konsole"] = func(t Scheme) (s string, err error) {
		color := func(name, color string) string {
			var c colorful.Color
			c, err = colorful.Hex(color)
			if err != nil {
				return ""
			}

			r, g, b := c.RGB255()
			return fmt.Sprintf(`[%v]
Color=%v,%v,%v

[%vFaint]
Color=%v,%v,%v

[%vIntense]
Color=%v,%v,%v

`,
				name, r, g, b,
				name, r, g, b,
				name, r, g, b)
		}

		for index, c := range t.ColorsNormal() {
			s += color(fmt.Sprintf("Color%v", index), c)
		}
		s += fmt.Sprintf(`[General]
Blur=false
ColorRandomization=false
Description=%v
Opacity=1
Wallpaper=
`, t.Name)
		return
	}
}
