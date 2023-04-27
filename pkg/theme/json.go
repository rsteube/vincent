package theme

import "encoding/json"

func init() {
	outputs["json"] = func(t Theme) (string, error) {
		m, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return "", err
		}
		return string(m), nil
	}
}
