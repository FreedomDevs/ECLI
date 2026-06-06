package templates

import (
	"ecli/internal/paths"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"path/filepath"
	"time"
)

func lastUpdateFile() string {
	return filepath.Join(paths.CacheDir(), ".last_update")
}

func ShouldUpdate(intervalHours int) bool {

	file := lastUpdateFile()

	info, err := os.Stat(file)
	if err != nil {
		return true
	}

	maxAge := time.Duration(intervalHours) * time.Hour

	return time.Since(info.ModTime()) > maxAge
}

func MarkUpdated() error {

	now := []byte(time.Now().Format(time.RFC3339))

	return os.WriteFile(
		lastUpdateFile(),
		now,
		0644,
	)
}

func UpdateRepo() error {

	repo, err := git.PlainOpen(paths.TemplatesDir())
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Force:      true,
	})

	if err == git.NoErrAlreadyUpToDate {
		return MarkUpdated()
	}

	if err != nil {
		return fmt.Errorf("pull failed: %w", err)
	}

	return MarkUpdated()
}
