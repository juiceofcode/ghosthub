package cmd

import (
	"fmt"

	"github.com/juiceofcode/ghosthub/internal/profile"
	"github.com/juiceofcode/ghosthub/internal/sshconfig"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Git profiles",
	Long: `List all Git profiles and their configurations.
This command will show:
1. Profile name
2. Git user name
3. Git user email
4. SSH key path`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := profile.LoadProfiles()
		if err != nil {
			return fmt.Errorf("failed to load Git profiles: %w", err)
		}

		sshProfiles, err := sshconfig.ListProfiles()
		if err != nil {
			return fmt.Errorf("failed to load SSH profiles: %w", err)
		}

		if len(profiles) == 0 && len(sshProfiles) == 0 {
			fmt.Println("No profiles found.")
			return nil
		}

		fmt.Println("\nGit Profiles:")
		fmt.Println("-------------")
		for name, prof := range profiles {
			fmt.Printf("Profile: %s\n", name)
			fmt.Printf("  Name: %s\n", prof.Name)
			fmt.Printf("  Email: %s\n", prof.Email)
			fmt.Printf("  SSH Key: %s\n", prof.SSHKeyPath)
			fmt.Println()
		}

		fmt.Println("SSH Profiles:")
		fmt.Println("------------")
		for _, name := range sshProfiles {
			fmt.Printf("- %s\n", name)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}