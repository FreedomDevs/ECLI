package paths

import (
	"os"
	"path/filepath"
	"strings"
)

func Home() string {
	h, _ := os.UserHomeDir()
	return h
}

func ConfigDir() string {
	if v := os.Getenv("XDG_CONFIG_HOME"); v != "" {
		return filepath.Join(v, "ecli")
	}
	return filepath.Join(Home(), ".config", "ecli")
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.json")
}

func CacheDir() string {
	if v := os.Getenv("XDG_CACHE_HOME"); v != "" {
		return filepath.Join(v, "ecli")
	}
	return filepath.Join(Home(), ".cache", "ecli")
}

func TemplatesDir() string {
	return filepath.Join(CacheDir(), "templates")
}

func Expand(path string) string {
	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}

