package scheme

import "encoding/json"

func init() {
	formats["windows"] = func(t Scheme) (string, error) {
		type windows struct {
			Name                string `json:"name"`
			CursorColor         string `json:"cursorColor"`
			SelectionBackground string `json:"selectionBackground"`
			Background          string `json:"background"`
			Foreground          string `json:"foreground"`
			Black               string `json:"black"`
			Red                 string `json:"red"`
			Green               string `json:"green"`
			Yellow              string `json:"yellow"`
			Blue                string `json:"blue"`
			Purple              string `json:"purple"`
			Cyan                string `json:"cyan"`
			White               string `json:"white"`
			BrightBlack         string `json:"brightBlack"`
			BrightRed           string `json:"brightRed"`
			BrightGreen         string `json:"brightGreen"`
			BrightYellow        string `json:"brightYellow"`
			BrightBlue          string `json:"brightBlue"`
			BrightPurple        string `json:"brightPurple"`
			BrightCyan          string `json:"brightCyan"`
			BrightWhite         string `json:"brightWhite"`
		}

		w := &windows{
			Name:        t.Name,
			CursorColor: t.Cursor,
			// TODO SelectionBackground
			Background:   t.Background,
			Foreground:   t.Foreground,
			Black:        t.Black,
			Red:          t.Red,
			Green:        t.Green,
			Yellow:       t.Yellow,
			Blue:         t.Blue,
			Purple:       t.Magenta,
			Cyan:         t.Cyan,
			White:        t.White,
			BrightBlack:  t.BrightBlack,
			BrightRed:    t.BrightRed,
			BrightGreen:  t.BrightGreen,
			BrightYellow: t.BrightYellow,
			BrightBlue:   t.BrightBlue,
			BrightPurple: t.BrightMagenta,
			BrightCyan:   t.BrightCyan,
			BrightWhite:  t.BrightWhite,
		}

		m, err := json.MarshalIndent(w, "", "  ")
		if err != nil {
			return "", err
		}
		return string(m), nil
	}
}
