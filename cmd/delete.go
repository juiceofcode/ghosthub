package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juiceofcode/ghosthub/internal/profile"
	"github.com/juiceofcode/ghosthub/internal/sshconfig"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [profile]",
	Short: "Delete a Git profile",
	Long: `Delete a Git profile and its associated SSH keys.
This command will:
1. Remove the profile from Git configuration
2. Remove the SSH keys
3. Update the SSH config`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := args[0]

		profiles, err := profile.LoadProfiles()
		if err != nil {
			return fmt.Errorf("failed to load Git profiles: %w", err)
		}

		if _, exists := profiles[profileName]; !exists {
			return fmt.Errorf("profile '%s' not found in Git configuration", profileName)
		}
		
		if err := sshconfig.RemoveFromConfig(profileName); err != nil {
			return fmt.Errorf("failed to remove from SSH config: %w", err)
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}

		ghosthubDir := filepath.Join(home, ".ssh", "ghosthub")
		keyTypes := []string{"ed25519", "rsa-4096", "rsa-2048", "ecdsa-521", "ecdsa-384", "ecdsa-256"}
		for _, keyType := range keyTypes {
			privKeyPath := filepath.Join(ghosthubDir, fmt.Sprintf("%s_id_%s", profileName, keyType))
			pubKeyPath := filepath.Join(ghosthubDir, fmt.Sprintf("%s_id_%s.pub", profileName, keyType))
			os.Remove(privKeyPath)
			os.Remove(pubKeyPath)
		}

		if err := profile.RemoveProfile(profileName); err != nil {
			return fmt.Errorf("failed to remove Git profile: %w", err)
		}

		fmt.Printf("✅ Profile '%s' deleted successfully!\n", profileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
} 