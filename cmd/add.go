package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juiceofcode/ghosthub/internal/profile"
	"github.com/juiceofcode/ghosthub/internal/sshconfig"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [profile]",
	Short: "Add a new Git profile",
	Long: `Add a new Git profile with SSH key and Git configuration.
	
Available key types:
- ed25519 (default)
- rsa-4096
- rsa-2048
- ecdsa-521
- ecdsa-384
- ecdsa-256`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := args[0]
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		keyType, _ := cmd.Flags().GetString("keygen")

		validKeyTypes := map[string]bool{
			"ed25519":   true,
			"rsa-4096":  true,
			"rsa-2048":  true,
			"ecdsa-521": true,
			"ecdsa-384": true,
			"ecdsa-256": true,
		}
		if !validKeyTypes[keyType] {
			return fmt.Errorf("invalid key type. Available options: ed25519, rsa-4096, rsa-2048, ecdsa-521, ecdsa-384, ecdsa-256")
		}

		if err := sshconfig.GenerateKeyPair(profileName, keyType); err != nil {
			return fmt.Errorf("failed to generate SSH key: %w", err)
		}

		if err := sshconfig.UpdateSSHConfig(profileName); err != nil {
			return fmt.Errorf("failed to update SSH config: %w", err)
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}

		ghosthubDir := filepath.Join(home, ".ssh", "ghosthub")
		gitProfile := profile.Profile{
			Name:       name,
			Email:      email,
			SSHKeyPath: filepath.Join(ghosthubDir, fmt.Sprintf("%s_id_%s", profileName, keyType)),
		}

		if err := profile.AddProfile(profileName, gitProfile); err != nil {
			return fmt.Errorf("failed to create Git profile: %w", err)
		}

		pubKeyPath := filepath.Join(ghosthubDir, fmt.Sprintf("%s_id_%s.pub", profileName, keyType))
		pubKey, err := os.ReadFile(pubKeyPath)
		if err != nil {
			return fmt.Errorf("failed to read public key: %w", err)
		}

		fmt.Printf("\n✅ Profile '%s' created successfully!\n", profileName)
		fmt.Printf("\nYour public key is:\n\n%s\n", string(pubKey))
		fmt.Println("\nPlease add this key to your GitHub/GitLab/Bitbucket account.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("name", "", "Git user name")
	addCmd.Flags().String("email", "", "Git user email")
	addCmd.Flags().String("keygen", "ed25519", "SSH key type to generate")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
} 