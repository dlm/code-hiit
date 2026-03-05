package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	historyFileName    = "history.json"
	oldHistoryFileName = ".typing-history.json"
)

func getHistoryPath() (string, error) {
	newPath, err := getConfigFile(historyFileName)
	if err != nil {
		return "", err
	}

	// Try to migrate from old location
	home, err := os.UserHomeDir()
	if err == nil {
		oldPath := filepath.Join(home, oldHistoryFileName)
		migrateOldFile(oldPath, newPath)
	}

	return newPath, nil
}

func LoadHistory() (*SessionHistory, error) {
	path, err := getHistoryPath()
	if err != nil {
		return &SessionHistory{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &SessionHistory{Sessions: []Session{}}, nil
		}
		return nil, err
	}

	var history SessionHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

func SaveHistory(history *SessionHistory) error {
	path, err := getHistoryPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (h *SessionHistory) AddSession(session Session) {
	h.Sessions = append(h.Sessions, session)
}
