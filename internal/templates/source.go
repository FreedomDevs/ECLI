package templates

import (
	"ecli/internal/config"
	"ecli/internal/paths"
)

func BasePath(cfg config.Config) string {
	if cfg.Debug && cfg.LocalTemplatesPath != "" {
		return paths.Expand(cfg.LocalTemplatesPath)
	}
	return paths.TemplatesDir()
}
