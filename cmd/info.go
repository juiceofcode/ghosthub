package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show current Git profile information",
	Long: `Show current Git profile information including:
- Git user.name
- Git user.email
- Active SSH key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		gitConfig, err := exec.Command("git", "config", "--global", "--list").Output()
		if err != nil {
			return fmt.Errorf("failed to get Git config: %w", err)
		}

		config := make(map[string]string)
		for _, line := range strings.Split(string(gitConfig), "\n") {
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				config[parts[0]] = parts[1]
			}
		}
		
		sshKey, err := exec.Command("ssh-add", "-l").Output()
		if err != nil {
			return fmt.Errorf("failed to get SSH keys: %w", err)
		}

		fmt.Println("Current Git Profile:")
		fmt.Println("-------------------")
		fmt.Printf("Name:  %s\n", config["user.name"])
		fmt.Printf("Email: %s\n", config["user.email"])
		fmt.Println("\nActive SSH Keys:")
		fmt.Println("----------------")
		fmt.Println(strings.TrimSpace(string(sshKey)))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
} 