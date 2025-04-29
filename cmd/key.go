package cmd

import (
	"fmt"
	"ghosthub-cli/internal/sshconfig"

	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage SSH keys",
}

var createKeyCmd = &cobra.Command{
	Use:   "create [profile]",
	Short: "Generate SSH key pair for profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profile := args[0]
		if err := sshconfig.GenerateKeyPair(profile, "ed25519"); err != nil {
			return err
		}
		fmt.Printf("SSH keys generated for '%s'.\n", profile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)
	keyCmd.AddCommand(createKeyCmd)
}
