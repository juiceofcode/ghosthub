package cmd

import (
	"fmt"

	"github.com/juiceofcode/ghosthub/internal/gitconfig"
	"github.com/juiceofcode/ghosthub/internal/sshconfig"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [profile]",
	Short: "Activate a Git profile",
	Long: `Activate a Git profile for immediate use.
This command will:
1. Set your Git user.name and user.email
2. Add the SSH key to ssh-agent
3. Show a system notification when done`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := args[0]

		if err := sshconfig.LoadProfile(profileName); err != nil {
			return fmt.Errorf("failed to load SSH profile: %w", err)
		}

		if err := gitconfig.SwitchProfile(profileName); err != nil {
			return fmt.Errorf("failed to switch Git profile: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}