package scheme

import "encoding/json"

func init() {
	outputs["json"] = func(t Scheme) (string, error) {
		m, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return "", err
		}
		return string(m), nil
	}
}
