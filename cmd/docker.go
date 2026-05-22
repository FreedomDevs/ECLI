package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use: "dockernetinit",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command(ShellPath, ScriptPath)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Printf("Ошибка выполнения: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
