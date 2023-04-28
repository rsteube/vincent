package vincent

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/vincent/pkg/scheme"
	"gopkg.in/yaml.v3"
)

//go:embed third_party/github.com/Gogh-Co/Gogh/*
var fs embed.FS

func Schemes() []string {
	entries, err := fs.ReadDir("third_party/github.com/Gogh-Co/Gogh/themes")
	if err != nil {
		panic("failed to read schemes")
	}

	schemes := make([]string, 0)
	for _, entry := range entries {
		schemes = append(schemes, strings.TrimSuffix(entry.Name(), ".yml"))
	}
	return schemes
}

func Load(name string) (*scheme.Scheme, error) {
	f := readEmbedded
	if strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml") {
		f = os.ReadFile
	}

	content, err := f(name)
	if err != nil {
		return nil, err
	}

	var scheme scheme.Scheme
	if err := yaml.Unmarshal(content, &scheme); err != nil {
		return nil, err
	}
	return &scheme, nil
}

func readEmbedded(name string) ([]byte, error) {
	content, err := fs.ReadFile("third_party/github.com/Gogh-Co/Gogh/themes/" + name + ".yml")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("unknown scheme: %v", name)
		}
		return nil, err
	}
	return content, err
}
