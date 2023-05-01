package scheme

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
	"gopkg.in/yaml.v3"
)

func init() {
	formats["tabby"] = func(t Scheme) (string, error) {
		tabby := struct {
			Name       string
			Foreground string
			Background string
			// Selection    string
			Cursor string
			// CursorAccent string
			Colors []string
		}{
			Name:       t.Name,
			Foreground: t.Foreground,
			Background: t.Background,
			Cursor:     t.Cursor,
			Colors:     make([]string, 0),
		}

		for _, color := range t.Colors() {
			c, err := colorful.Hex(color)
			if err != nil {
				return "", err
			}
			h, s, l := c.Hsl()
			tabby.Colors = append(tabby.Colors, fmt.Sprintf("HSL(%v, %v%%, %v%%)", int(h), int(s*100), int(l*100)))

		}

		m, err := yaml.Marshal(tabby)
		if err != nil {
			return "", err
		}
		return string(m), nil
	}
}
