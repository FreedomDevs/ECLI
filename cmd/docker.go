package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

const (
	netName    = "svc-net"
	subnetIPv6 = "fd98:2dd6:8f48:1d99::/64"
	mtuValue   = "65535"
)

var ErrAborted = fmt.Errorf("aborted")

var (
	styleInfo = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	styleOk   = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	styleWarn = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	styleErr  = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	styleBold = lipgloss.NewStyle().Bold(true)
)

func info(msg string)   { fmt.Println(styleInfo.Render("[INFO] " + msg)) }
func ok(msg string)     { fmt.Println(styleOk.Render("[OK]   " + msg)) }
func warn(msg string)   { fmt.Println(styleWarn.Render("[WARN] " + msg)) }
func errMsg(msg string) { fmt.Println(styleErr.Render("[ERR]  " + msg)) }

func askYesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(styleInfo.Render(question + " [y/N]: "))
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(strings.ToLower(text))

	return text == "y" || text == "yes"
}

var dockerCmd = &cobra.Command{
	Use:           "dockernetinit",
	Short:         "Initialize docker IPv6 network",
	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {

		info("Checking Docker network: " + netName)

		exists, err := networkExists(netName)
		if err != nil {
			return err
		}

		if exists {
			warn("Network '" + netName + "' already exists")

			containers, err := getContainersInNetwork(netName)
			if err != nil {
				return err
			}

			if len(containers) > 0 {
				errMsg("Network is currently in use by containers:")

				for _, c := range containers {
					fmt.Println("  - " + styleBold.Render(c))
				}

				fmt.Println()

				if !askYesNo("Stop these containers automatically?") {
					errMsg("Aborted by user")
					return ErrAborted
				}

				for _, c := range containers {
					info("Stopping container: " + c)

					stopCmd := exec.Command("docker", "stop", c)
					stopCmd.Stdout = os.Stdout
					stopCmd.Stderr = os.Stderr

					if err := stopCmd.Run(); err != nil {
						errMsg("Failed to stop container: " + c)
						return err
					}

					ok("Stopped: " + c)
				}
			}

			info("Removing existing network...")

			if err := removeNetwork(netName); err != nil {
				errMsg("Failed to remove network")
				return err
			}

			ok("Old network removed")

		} else {
			ok("Network '" + netName + "' not found (good)")
		}

		info("Creating new IPv6-only Docker network...")

		if err := createNetwork(netName, subnetIPv6, mtuValue); err != nil {
			errMsg("Network creation failed")
			return err
		}

		ok("Network '" + netName + "' created successfully")
		fmt.Println(styleInfo.Render("[INFO] Subnet: " + subnetIPv6))
		fmt.Println(styleInfo.Render("[INFO] MTU: " + mtuValue))

		return nil
	},
}

// HELPERS

func networkExists(name string) (bool, error) {
	cmd := exec.Command("docker", "network", "inspect", name)
	err := cmd.Run()
	return err == nil, nil
}

func getContainersInNetwork(name string) ([]string, error) {
	cmd := exec.Command("docker", "network", "inspect", name, "-f", "{{json .Containers}}")
	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return nil, nil
	}

	var raw map[string]struct {
		Name string `json:"Name"`
	}

	if err := json.Unmarshal(out, &raw); err != nil {
		return nil, nil
	}

	var result []string
	for _, c := range raw {
		result = append(result, c.Name)
	}

	return result, nil
}

func removeNetwork(name string) error {
	cmd := exec.Command("docker", "network", "rm", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createNetwork(name, subnet, mtu string) error {
	cmd := exec.Command(
		"docker",
		"network",
		"create",
		"--ipv6",
		"--subnet="+subnet,
		"--opt", "com.docker.network.driver.mtu="+mtu,
		name,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}

