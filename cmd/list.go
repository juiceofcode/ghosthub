package cmd

import (
	"fmt"
	"ghosthub-cli/internal/sshconfig"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists configured SSH profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := sshconfig.ListProfiles()
		if err != nil {
			return err
		}
		for _, p := range profiles {
			fmt.Println(p)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}