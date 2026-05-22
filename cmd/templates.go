package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"ecli/internal/config"
	"ecli/internal/paths"
	"ecli/internal/templates"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Manage templates cache",
}

var templatesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update templates cache from remote repo",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("🔄 Updating templates...")

		if err := templates.UpdateRepo(); err != nil {
			fmt.Println("❌ update failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Templates updated")
	},
}

var templatesResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete templates cache and re-clone",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("🧹 Resetting templates cache...")

		if err := os.RemoveAll(paths.TemplatesDir()); err != nil {
			fmt.Println("❌ failed to remove cache:", err)
			os.Exit(1)
		}

		cfg, err := config.LoadOrCreate()
		if err != nil {
			fmt.Println("config error:", err)
			os.Exit(1)
		}

		if err := templates.EnsureCloned(*cfg); err != nil {
			fmt.Println("❌ failed to re-clone templates:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Templates reset completed")
	},
}

var templatesPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show templates cache path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(paths.TemplatesDir())
	},
}

var templatesCloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Force clone templates repository into cache",
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := config.LoadOrCreate()
		if err != nil {
			fmt.Println("config error:", err)
			os.Exit(1)
		}

		if err := os.RemoveAll(paths.TemplatesDir()); err != nil {
			fmt.Println("❌ failed to clean cache:", err)
			os.Exit(1)
		}

		if err := templates.EnsureCloned(*cfg); err != nil {
			fmt.Println("❌ clone failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Templates cloned successfully")
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)

	templatesCmd.AddCommand(templatesUpdateCmd)
	templatesCmd.AddCommand(templatesResetCmd)
	templatesCmd.AddCommand(templatesPathCmd)
	templatesCmd.AddCommand(templatesCloneCmd)
}
