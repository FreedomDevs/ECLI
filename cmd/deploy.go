package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:           "deploy [service_name]",
	Short:         "Deploy service in .",
	Args:          cobra.ExactArgs(1), // Гарантирует, что передан ровно один аргумент
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		svcName := args[0]
		tag := fmt.Sprintf("%s:latest", svcName)

		fmt.Printf("=== building image: %s ===\n", tag)

		buildCmd := exec.Command("docker", "build", "--network=host", ".", "-t", tag)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			return fmt.Errorf("failed to build docker image: %w", err)
		}

		fmt.Printf("=== deploying %s to remote server ===\n", tag)

		pipelineStr := fmt.Sprintf(
			"docker save %s | zstd -3 - | pv -b | ssh mikinol-serv \"zstd -d - | docker load && sudo /usr/local/bin/deploy/deploy\"",
			tag,
		)

		deployCmd := exec.Command("bash", "-c", pipelineStr)
		deployCmd.Stdout = os.Stdout
		deployCmd.Stderr = os.Stderr

		if err := deployCmd.Run(); err != nil {
			return fmt.Errorf("failed to deploy image: %w", err)
		}

		fmt.Println("=== deployment finished successfully ===")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
