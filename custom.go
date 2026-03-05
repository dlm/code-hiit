package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const customSnippetsFileName = ".typing-snippets.json"

type CustomSnippetFile struct {
	Snippets []CustomSnippetEntry `json:"snippets"`
}

type CustomSnippetEntry struct {
	Content  string `json:"content"`
	Language string `json:"language,omitempty"`
}

func getCustomSnippetsPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, customSnippetsFileName), nil
}

func LoadCustomSnippets() ([]CodeSnippet, error) {
	path, err := getCustomSnippetsPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []CodeSnippet{}, nil
		}
		return nil, err
	}

	var file CustomSnippetFile
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, err
	}

	snippets := make([]CodeSnippet, len(file.Snippets))
	for i, entry := range file.Snippets {
		language := entry.Language
		if language == "" {
			language = "Custom"
		}
		snippets[i] = CodeSnippet{
			Content:    entry.Content,
			Language:   language,
			Difficulty: Custom,
		}
	}

	return snippets, nil
}
