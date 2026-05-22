package templates

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"

	"ecli/internal/config"
	"ecli/internal/paths"
	"path/filepath"
)

func EnsureCloned(cfg config.Config) error {

	if cfg.Debug && cfg.LocalTemplatesPath != "" {
		fmt.Println("🧪 DEBUG MODE: using local templates")
		return nil
	}

	gitDir := filepath.Join(paths.TemplatesDir(), ".git")

	updated := false

	if _, err := os.Stat(gitDir); os.IsNotExist(err) {

		fmt.Println("📦 Cloning templates repo...")

		if err := os.RemoveAll(paths.TemplatesDir()); err != nil {
			return err
		}

		if err := os.MkdirAll(paths.CacheDir(), 0755); err != nil {
			return err
		}

		_, err := git.PlainClone(
			paths.TemplatesDir(),
			false,
			&git.CloneOptions{
				URL: cfg.TemplatesRepo,
			},
		)

		if err != nil {
			return fmt.Errorf("failed to clone templates repo: %w", err)
		}

		updated = true
	}

	if cfg.AutoUpdateTemplates && ShouldUpdate(cfg.UpdateIntervalHours) {

		fmt.Println("🔄 Updating templates...")

		if err := UpdateRepo(); err != nil {
			return err
		}

		updated = true
	}

	if updated {
		return MarkUpdated()
	}

	return nil
}
