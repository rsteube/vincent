package scheme

import (
	"gopkg.in/yaml.v3"
)

func init() {
	formats["yaml"] = func(t Scheme) (string, error) {
		m, err := yaml.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(m), nil
	}
}
