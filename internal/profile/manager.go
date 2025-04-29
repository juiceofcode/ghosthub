package profile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	profilesDir  = ".ghosthub"
	profilesFile = "profiles.json"
)

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}
	return filepath.Join(home, profilesDir), nil
}

func getProfilesPath() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, profilesFile), nil
}

func ensureConfigDir() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(configDir, 0755)
}

func LoadProfiles() (Profiles, error) {
	profilesPath, err := getProfilesPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(profilesPath)
	if os.IsNotExist(err) {
		return make(Profiles), nil
	}
	if err != nil {
		return nil, fmt.Errorf("error reading profiles.json: %w", err)
	}

	var profiles Profiles
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, fmt.Errorf("error decoding profiles.json: %w", err)
	}

	return profiles, nil
}

func SaveProfiles(profiles Profiles) error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	profilesPath, err := getProfilesPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding profiles.json: %w", err)
	}

	return os.WriteFile(profilesPath, data, 0644)
}

func AddProfile(name string, profile Profile) error {
	profiles, err := LoadProfiles()
	if err != nil {
		return err
	}

	profiles[name] = profile
	return SaveProfiles(profiles)
}

func RemoveProfile(name string) error {
	profiles, err := LoadProfiles()
	if err != nil {
		return err
	}

	if _, exists := profiles[name]; !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}

	delete(profiles, name)
	return SaveProfiles(profiles)
}

func GetProfile(name string) (Profile, error) {
	profiles, err := LoadProfiles()
	if err != nil {
		return Profile{}, err
	}

	profile, exists := profiles[name]
	if !exists {
		return Profile{}, fmt.Errorf("profile '%s' not found", name)
	}

	return profile, nil
}

func ListProfiles() ([]string, error) {
	profiles, err := LoadProfiles()
	if err != nil {
		return nil, err
	}

	var names []string
	for name := range profiles {
		names = append(names, name)
	}
	return names, nil
}