package sshconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const sshDir = ".ssh/ghosthub"

func GenerateKeyPair(profile string) error {
	home := os.Getenv("HOME")
	if home == "" {
		home = "/root"
	}
	pairDir := filepath.Join(home, sshDir)
	if err := os.MkdirAll(pairDir, 0700); err != nil {
		return err
	}
	priv := filepath.Join(pairDir, fmt.Sprintf("%s_id_ed25519", profile))
	pub := priv + ".pub"
	if _, err := os.Stat(priv); err == nil {
		return errors.New("key already exists")
	}
	err := exec.Command("ssh-keygen", "-t", "ed25519", "-f", priv, "-N", "", "-C", profile).Run()
	if err != nil {
		return err
	}
	fmt.Printf("Keys generated: %s, %s\n", priv, pub)
	return nil
}

func UpdateSSHConfig(profile string) error {
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".ssh", "config")
	block := fmt.Sprintf(`
Host github-%[1]s
  HostName github.com
  User git
  IdentityFile %s/%[1]s_id_ed25519
  IdentitiesOnly yes
`, profile, filepath.Join(home, sshDir))
	f, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(block); err != nil {
		return err
	}
	return nil
}

func ListProfiles() ([]string, error) {
	home := os.Getenv("HOME")
	dirs, err := ioutil.ReadDir(filepath.Join(home, sshDir))
	if err != nil {
		return nil, err
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
	home := os.Getenv("HOME")
	priv := filepath.Join(home, sshDir, fmt.Sprintf("%s_id_ed25519", profile))
	if _, err := os.Stat(priv); err != nil {
		return errors.New("profile not found")
	}
	err := exec.Command("ssh-add", priv).Run()
	if err != nil {
		return err
	}
	return nil
}