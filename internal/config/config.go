package config

import (
	"encoding/json"
	"os"

	"ecli/internal/paths"
)

type Config struct {
	TemplatesRepo      string `json:"templates_repo"`
	Debug              bool   `json:"debug"`
	LocalTemplatesPath string `json:"local_templates_path"`
}

func Exists() bool {
	_, err := os.Stat(paths.ConfigFile())
	return err == nil
}

func Default() Config {
	return Config{
		TemplatesRepo:      "https://github.com/FreedomDevs/templates",
		Debug:              false,
		LocalTemplatesPath: "",
	}
}

func Load() (*Config, error) {
	b, err := os.ReadFile(paths.ConfigFile())
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(b, &cfg)
	return &cfg, err
}

func Save(cfg Config) error {
	os.MkdirAll(paths.ConfigDir(), 0755)

	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(paths.ConfigFile(), data, 0644)
}

func LoadOrCreate() (*Config, error) {
	if Exists() {
		return Load()
	}

	cfg := Default()

	if err := Save(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
