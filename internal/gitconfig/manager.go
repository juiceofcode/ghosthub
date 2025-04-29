package gitconfig

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/juiceofcode/ghosthub-cli/internal/profile"
)

func SwitchProfile(profileName string) error {
	prof, err := profile.GetProfile(profileName)
	if err != nil {
		return fmt.Errorf("error loading profile: %w", err)
	}

	if err := exec.Command("git", "config", "--global", "user.name", prof.Name).Run(); err != nil {
		return fmt.Errorf("error configuring user.name: %w", err)
	}
	if err := exec.Command("git", "config", "--global", "user.email", prof.Email).Run(); err != nil {
		return fmt.Errorf("error configuring user.email: %w", err)
	}

	sendNotification(profileName)

	fmt.Printf("\n✅ Profile '%s' successfully activated!\n- Git: %s <%s>\n- SSH Key: %s\n\n", 
		profileName, prof.Name, prof.Email, prof.SSHKeyPath)
	fmt.Println("⚠️  Remember to add the SSH key manually with:")
	fmt.Printf("   ssh-add %s\n", prof.SSHKeyPath)
	return nil
}

func sendNotification(profileName string) {
	title := "ghosthub-cli"
	message := fmt.Sprintf("Profile '%s' successfully activated!", profileName)
	switch runtime.GOOS {
	case "linux":
		exec.Command("notify-send", title, message).Run()
	case "darwin":
		exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)).Run()
	case "windows":
		exec.Command("powershell", "-Command", fmt.Sprintf(`New-BurntToastNotification -Text '%s', '%s'`, title, message)).Run()
	}
}
