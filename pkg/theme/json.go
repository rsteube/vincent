package theme

import "encoding/json"

func init() {
	outputs["json"] = func(t Theme) string {
		if m, err := json.MarshalIndent(t, "", "  "); err == nil {
			return string(m)
		}
		return ""
	}
}
