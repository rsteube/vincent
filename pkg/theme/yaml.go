package theme

import (
	"gopkg.in/yaml.v3"
)

func init() {
	outputs["yaml"] = func(t Theme) (string, error) {
		m, err := yaml.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(m), nil
	}
}
