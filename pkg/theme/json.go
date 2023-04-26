package theme

import "encoding/json"

func init() {
	outputs["json"] = func(t Theme) string {
		if m, err := json.Marshal(t); err == nil {
			return string(m)
		}
		return ""
	}
}
