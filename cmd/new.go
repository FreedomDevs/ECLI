package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"ecli/internal/config"
	"ecli/internal/templates"
	"ecli/internal/ui"
	"sort"
)

func runGenerate(cfg *config.Config, reg templates.Registry, projectName string, selectedName string, nameToKey map[string]string) {

	selectedKey := nameToKey[selectedName]

	fmt.Println("\n🚀 Creating:", projectName)
	fmt.Println("📦 Template:", reg[selectedKey].Name)

	if err := templates.CopyTemplate(*cfg, reg[selectedKey].Path, projectName); err != nil {
		fmt.Println("❌ generation failed:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Done!")
}

var newCmd = &cobra.Command{
	Use: "new [project-name]",
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := config.LoadOrCreate()
		if err != nil {
			fmt.Println("config error:", err)
			os.Exit(1)
		}

		if err := templates.EnsureCloned(*cfg); err != nil {
			fmt.Println("❌ failed to prepare templates:", err)
			os.Exit(1)
		}

		reg, err := templates.LoadRegistry(*cfg)
		if err != nil {
			fmt.Println("❌ failed to load registry:", err)
			os.Exit(1)
		}

		items := []string{}
		nameToKey := map[string]string{}

		for k, v := range reg {
			items = append(items, v.Name)
			nameToKey[v.Name] = k
		}

		sort.Strings(items)

		var projectName string
		var selectedName string

		if len(args) == 0 {

			w := ui.NewWizard(items)

			p := tea.NewProgram(w)
			m, _ := p.Run()

			model := m.(*ui.Wizard)

			projectName = model.ProjectName
			selectedName = model.Choice

			runGenerate(cfg, reg, projectName, selectedName, nameToKey)
			return
		}

		projectName = args[0]

		p := tea.NewProgram(ui.NewSelector(items))
		m, _ := p.Run()

		model := m.(*ui.Model)

		selectedName = model.Choice

		runGenerate(cfg, reg, projectName, selectedName, nameToKey)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
