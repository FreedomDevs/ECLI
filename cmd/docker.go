package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use: "docker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🐳 HelloWorld from ECLI docker command")
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
