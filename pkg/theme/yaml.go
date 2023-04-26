package theme

import (
	"gopkg.in/yaml.v3"
)

func init() {
	outputs["yaml"] = func(t Theme) string {
		if m, err := yaml.Marshal(t); err == nil {
			return string(m)
		}
		return ""
	}
}
