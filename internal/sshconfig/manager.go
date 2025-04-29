package sshconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const sshDir = ".ssh/ghosthub"

func getHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE") 
	}
	if home == "" {
		return "", errors.New("could not determine home directory")
	}
	return home, nil
}

func ensureSSHDir() (string, error) {
	home, err := getHomeDir()
	if err != nil {
		return "", err
	}

	sshDir := filepath.Join(home, ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create .ssh directory: %w", err)
	}

	ghosthubDir := filepath.Join(sshDir, "ghosthub")
	if err := os.MkdirAll(ghosthubDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create ghosthub directory: %w", err)
	}

	return ghosthubDir, nil
}

func setupWindowsSSHAgent() error {
	cmd := exec.Command("where", "ssh")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("OpenSSH not found. Please install it from: https://github.com/PowerShell/Win32-OpenSSH/releases")
	}

	cmd = exec.Command("sc", "query", "ssh-agent")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ssh-agent service not found. Please install OpenSSH properly")
	}

	cmd = exec.Command("sc", "config", "ssh-agent", "start=auto")
	if err := cmd.Run(); err != nil {
		fmt.Println("⚠️ Configurando ssh-agent automaticamente...")
		
		psCommand := "Set-Service -Name ssh-agent -StartupType Automatic; Start-Service ssh-agent"
		cmd = exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process powershell -Verb RunAs -ArgumentList '-Command %s'", psCommand))
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Não foi possível configurar o ssh-agent automaticamente.")
			fmt.Println("Por favor, execute estes comandos manualmente como Administrador:")
			fmt.Println("1. Abra o PowerShell como Administrador")
			fmt.Println("2. Execute: Set-Service -Name ssh-agent -StartupType Automatic")
			fmt.Println("3. Execute: Start-Service ssh-agent")
			return nil
		}
		
		fmt.Println("✅ ssh-agent configurado com sucesso!")
	}

	cmd = exec.Command("net", "start", "ssh-agent")
	if err := cmd.Run(); err != nil {
		if !strings.Contains(err.Error(), "service has already been started") {
			fmt.Println("⚠️ Iniciando ssh-agent automaticamente...")
			
			cmd = exec.Command("powershell", "-Command", "Start-Process powershell -Verb RunAs -ArgumentList '-Command Start-Service ssh-agent'")
			if err := cmd.Run(); err != nil {
				fmt.Println("❌ Não foi possível iniciar o ssh-agent automaticamente.")
				fmt.Println("Por favor, execute manualmente como Administrador:")
				fmt.Println("1. Abra o PowerShell como Administrador")
				fmt.Println("2. Execute: Start-Service ssh-agent")
				return nil
			}
			
			fmt.Println("✅ ssh-agent iniciado com sucesso!")
		}
	}

	return nil
}

func GenerateKeyPair(profile, keyType string) error {
	home, err := getHomeDir()
	if err != nil {
		return err
	}

	ghosthubDir := filepath.Join(home, ".ssh", "ghosthub")
	if err := os.MkdirAll(ghosthubDir, 0700); err != nil {
		return fmt.Errorf("failed to create ghosthub directory: %w", err)
	}

	keyPath := filepath.Join(ghosthubDir, fmt.Sprintf("%s_id_%s", profile, keyType))
	
	args := []string{"-t", keyType}
	if strings.HasPrefix(keyType, "rsa") {
		bits := strings.Split(keyType, "-")[1]
		args = append(args, "-b", bits)
	}
	args = append(args, "-f", keyPath, "-N", "", "-C", fmt.Sprintf("ghosthub-%s", profile))

	cmd := exec.Command("ssh-keygen", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	return nil
}

func UpdateSSHConfig(profile string) error {
	home, err := getHomeDir()
	if err != nil {
		return err
	}

	ghosthubDir, err := ensureSSHDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".ssh", "config")
	block := fmt.Sprintf(`
Host github-%[1]s
  HostName github.com
  User git
  IdentityFile %s/%[1]s_id_ed25519
  IdentitiesOnly yes
`, profile, ghosthubDir)

	f, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open SSH config: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(block); err != nil {
		return fmt.Errorf("failed to write SSH config: %w", err)
	}
	return nil
}

func ListProfiles() ([]string, error) {
	ghosthubDir, err := ensureSSHDir()
	if err != nil {
		return nil, err
	}

	dirs, err := ioutil.ReadDir(ghosthubDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("failed to read SSH directory: %w", err)
	}

	var profiles []string
	for _, fi := range dirs {
		if fi.IsDir() {
			continue
		}
		name := fi.Name()
		if filepath.Ext(name) == ".pub" {
			profiles = append(profiles, name[:len(name)-len("_id_ed25519.pub")])
		}
	}
	return profiles, nil
}

func LoadProfile(profile string) error {
	ghosthubDir, err := ensureSSHDir()
	if err != nil {
		return err
	}

	priv := filepath.Join(ghosthubDir, fmt.Sprintf("%s_id_ed25519", profile))
	if _, err := os.Stat(priv); err != nil {
		return fmt.Errorf("profile not found: %w", err)
	}

	if runtime.GOOS == "windows" {
		if err := setupWindowsSSHAgent(); err != nil {
			return fmt.Errorf("failed to setup ssh-agent: %w", err)
		}
	}

	cmd := exec.Command("ssh-add", priv)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("⚠️ Adicionando chave ao ssh-agent automaticamente...")
		
		cmd = exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process powershell -Verb RunAs -ArgumentList '-Command ssh-add %s'", priv))
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Não foi possível adicionar a chave automaticamente.")
			fmt.Println("Por favor, execute manualmente como Administrador:")
			fmt.Println("1. Abra o PowerShell como Administrador")
			fmt.Println("2. Execute: ssh-add", priv)
			return nil
		}
		
		fmt.Println("✅ Chave adicionada com sucesso!")
		return nil
	}

	fmt.Printf("✅ SSH key for profile '%s' added to ssh-agent\n", profile)
	return nil
}

func RemoveFromConfig(profile string) error {
	home, err := getHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".ssh", "config")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil 
		}
		return fmt.Errorf("failed to read SSH config: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	var newLines []string
	inBlock := false
	hostPattern := fmt.Sprintf("Host github-%s", profile)

	for _, line := range lines {
		if strings.TrimSpace(line) == hostPattern {
			inBlock = true
			continue
		}
		if inBlock && strings.TrimSpace(line) == "" {
			inBlock = false
			continue
		}
		if !inBlock {
			newLines = append(newLines, line)
		}
	}

	if err := os.WriteFile(configPath, []byte(strings.Join(newLines, "\n")), 0600); err != nil {
		return fmt.Errorf("failed to write SSH config: %w", err)
	}

	return nil
}