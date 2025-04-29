package cmd

import (
	"fmt"

	"github.com/juiceofcode/ghosthub/internal/gitconfig"
	"github.com/juiceofcode/ghosthub/internal/profile"

	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manages Git profiles",
	Long:  `Commands to manage your Git profiles, including adding, removing, listing and switching between profiles.`,
}

var profileAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Adds a new profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		email, _ := cmd.Flags().GetString("email")
		gitName, _ := cmd.Flags().GetString("name")
		sshKeyPath, _ := cmd.Flags().GetString("ssh-key")

		if gitName == "" {
			gitName = name
		}

		prof := profile.Profile{
			Name:       gitName,
			Email:      email,
			SSHKeyPath: sshKeyPath,
		}

		if err := profile.AddProfile(name, prof); err != nil {
			return fmt.Errorf("error adding profile: %w", err)
		}

		fmt.Printf("✅ Profile '%s' added successfully!\n", name)
		return nil
	},
}

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := profile.ListProfiles()
		if err != nil {
			return fmt.Errorf("error listing profiles: %w", err)
		}

		if len(profiles) == 0 {
			fmt.Println("No profiles found.")
			return nil
		}

		fmt.Println("Available profiles:")
		for _, name := range profiles {
			prof, err := profile.GetProfile(name)
			if err != nil {
				continue
			}
			fmt.Printf("- %s:\n  Name: %s\n  Email: %s\n  SSH Key: %s\n\n", 
				name, prof.Name, prof.Email, prof.SSHKeyPath)
		}
		return nil
	},
}

var profileSwitchCmd = &cobra.Command{
	Use:   "switch [name]",
	Short: "Switches to a specific profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return gitconfig.SwitchProfile(args[0])
	},
}

var profileRemoveCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Removes a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := profile.RemoveProfile(name); err != nil {
			return fmt.Errorf("error removing profile: %w", err)
		}
		fmt.Printf("✅ Profile '%s' removed successfully!\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(profileAddCmd, profileListCmd, profileSwitchCmd, profileRemoveCmd)

	profileAddCmd.Flags().String("email", "", "Git profile email")
	profileAddCmd.Flags().String("name", "", "Git profile name (optional, uses profile name by default)")
	profileAddCmd.Flags().String("ssh-key", "", "Path to SSH key")
	
	profileAddCmd.MarkFlagRequired("email")
	profileAddCmd.MarkFlagRequired("ssh-key")
}