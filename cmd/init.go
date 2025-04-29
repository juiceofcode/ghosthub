package cmd

import (
	"fmt"
	"ghosthub-cli/internal/sshconfig"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [profile]",
	Short: "Initialize a new SSH profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profile := args[0]
		if err := sshconfig.GenerateKeyPair(profile); err != nil {
			return err
		}
		if err := sshconfig.UpdateSSHConfig(profile); err != nil {
			return err
		}
		fmt.Printf("Profile '%s' initialized successfully!\n", profile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
