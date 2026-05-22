package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ShellPath  = "sh"
	ScriptPath = "./create_svc_network.sh"
)

var rootCmd = &cobra.Command{
	Use: "ecli",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
