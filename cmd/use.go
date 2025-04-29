package cmd

import (
	"fmt"
	"ghosthub-cli/internal/gitconfig"
	"ghosthub-cli/internal/sshconfig"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [profile]",
	Short: "Activates an SSH profile for immediate use",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profile := args[0]
		if err := sshconfig.LoadProfile(profile); err != nil {
			return err
		}
		if err := gitconfig.SwitchProfile(profile); err != nil {
			return err
		}
		fmt.Printf("Profile '%s' activated!\n", profile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}