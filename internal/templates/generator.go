package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ecli/internal/config"
)

func CopyTemplate(cfg config.Config, templatePath, projectName string) error {

	base := BasePath(cfg)

	src := filepath.Join(base, templatePath, "template")
	dst := filepath.Join(".", projectName)

	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("directory already exists: %s", projectName)
	}

	fmt.Println("📦 Generating project...")

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		content := strings.ReplaceAll(
			string(data),
			"{{PROJECT_NAME}}",
			projectName,
		)

		return os.WriteFile(
			target,
			[]byte(content),
			info.Mode(),
		)
	})
}

