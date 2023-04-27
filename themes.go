package vincent

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/vincent/pkg/theme"
	"gopkg.in/yaml.v3"
)

//go:embed third_party/github.com/Gogh-Co/Gogh/*
var fs embed.FS

func Themes() []string {
	entries, err := fs.ReadDir("third_party/github.com/Gogh-Co/Gogh/themes")
	if err != nil {
		panic("failed to read themes")
	}

	themes := make([]string, 0)
	for _, entry := range entries {
		themes = append(themes, strings.TrimSuffix(entry.Name(), ".yml"))
	}
	return themes
}

func Load(name string) (*theme.Theme, error) {
	content, err := fs.ReadFile("third_party/github.com/Gogh-Co/Gogh/themes/" + name + ".yml")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("unknown theme: %v", name)
		}
		return nil, err
	}

	var theme theme.Theme
	if err := yaml.Unmarshal(content, &theme); err != nil {
		return nil, err
	}
	return &theme, nil
}
