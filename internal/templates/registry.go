package templates

import (
	"encoding/json"
	"os"
	"path/filepath"

	"ecli/internal/config"
)

type Template struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

type Registry map[string]Template

func LoadRegistry(cfg config.Config) (Registry, error) {

	base := BasePath(cfg)

	file := filepath.Join(base, "registry.json")

	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var r Registry
	err = json.Unmarshal(b, &r)
	return r, err
}
