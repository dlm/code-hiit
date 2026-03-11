package main

import (
	"os"
	"path/filepath"
)

// getConfigDir returns the code-hiit config directory following XDG spec.
// Defaults to ~/.config/code-hiit
func getConfigDir() (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configHome = filepath.Join(home, ".config")
	}

	configDir := filepath.Join(configHome, "code-hiit")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// getConfigFile returns the full path to a config file
func getConfigFile(filename string) (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, filename), nil
}

// migrateOldFile moves a file from old location to new if it exists
func migrateOldFile(oldPath, newPath string) error {
	// Check if old file exists
	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		return nil // Nothing to migrate
	}

	// Check if new file already exists
	if _, err := os.Stat(newPath); err == nil {
		return nil // New file exists, don't overwrite
	}

	// Move the file
	return os.Rename(oldPath, newPath)
}
