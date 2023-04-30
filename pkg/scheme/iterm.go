package scheme

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	formats["iterm"] = func(t Scheme) (s string, err error) {
		color := func(name, color string) string {
			var c colorful.Color
			c, err = colorful.Hex(color)
			if err != nil {
				return ""
			}

			return fmt.Sprintf(`  <key>%v</key>
  <dict>
    <key>Color Space</key>
    <string>sRGB</string>
    <key>Red Component</key>
    <real>%v</real>
    <key>Green Component</key>
    <real>%v</real>
    <key>Blue Component</key>
    <real>%v</real>
    <key>Alpha Component</key>
    <real>1</real>
  </dict>
`,
				name, c.R, c.G, c.B)
		}

		s = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
`
		for index, c := range t.Colors() {
			s += color(fmt.Sprintf("Ansi %v Color", index), c)
		}
		s += color("Background Color", t.Background)
		s += color("Foreground Color", t.Foreground)
		s += color("Cursor Color", t.Cursor)

		// TODO Link, Bold, Cursor Text, Cursor Guide, Selection, Selection Text

		s += `</dict>
</plist>`

		return
	}
}
