package templates

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"

	"ecli/internal/config"
	"ecli/internal/paths"
)

func EnsureCloned(cfg config.Config) error {

	if cfg.Debug && cfg.LocalTemplatesPath != "" {
		fmt.Println("🧪 DEBUG MODE: using local templates")
		return nil
	}

	if _, err := os.Stat(paths.TemplatesDir()); err == nil {
		return nil
	}

	fmt.Println("📦 Cloning templates repo...")

	os.MkdirAll(paths.CacheDir(), 0755)

	_, err := git.PlainClone(paths.TemplatesDir(), false, &git.CloneOptions{
		URL: cfg.TemplatesRepo,
	})

	if err != nil {
		return fmt.Errorf("failed to clone templates repo: %w", err)
	}

	return nil
}

