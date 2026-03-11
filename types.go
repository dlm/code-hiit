package main

import (
	"encoding/json"
	"time"
)

type Mode int

const (
	EasyCode Mode = iota
	MediumCode
	HardCode
	NumbersPractice
	SymbolsPractice
	HexNumbers
	BracketsPractice
	RegexPatterns
	Custom
)

type CodeSnippet struct {
	Content  string `json:"content"`
	Language string `json:"language"`
	Mode     Mode   `json:"mode"`
}

type TypingStats struct {
	TotalChars     int
	CorrectChars   int
	IncorrectChars int
	TotalSymbols   int
	CorrectSymbols int
	TotalNumbers   int
	CorrectNumbers int
	StartTime      time.Time
	EndTime        time.Time
	Errors         []TypingError
}

type TypingError struct {
	Position int
	Expected rune
	Actual   rune
}

func (ts *TypingStats) WPM() float64 {
	if ts.StartTime.IsZero() || ts.EndTime.IsZero() {
		return 0
	}
	minutes := ts.EndTime.Sub(ts.StartTime).Minutes()
	if minutes == 0 {
		return 0
	}
	words := float64(ts.CorrectChars) / 5.0
	return words / minutes
}

func (ts *TypingStats) Accuracy() float64 {
	if ts.TotalChars == 0 {
		return 0
	}
	return float64(ts.CorrectChars) / float64(ts.TotalChars) * 100
}

func (ts *TypingStats) SymbolAccuracy() float64 {
	if ts.TotalSymbols == 0 {
		return 0
	}
	return float64(ts.CorrectSymbols) / float64(ts.TotalSymbols) * 100
}

func (ts *TypingStats) NumberAccuracy() float64 {
	if ts.TotalNumbers == 0 {
		return 0
	}
	return float64(ts.CorrectNumbers) / float64(ts.TotalNumbers) * 100
}

type SessionHistory struct {
	Sessions []Session `json:"sessions"`
}

type Session struct {
	Date     time.Time   `json:"date"`
	Mode     Mode        `json:"mode"`
	Language string      `json:"language"`
	Stats    TypingStats `json:"stats"`
}

// UnmarshalJSON accepts both the new "mode" field and legacy "difficulty" data.
func (s *Session) UnmarshalJSON(data []byte) error {
	type sessionAlias struct {
		Date       time.Time   `json:"date"`
		Mode       *Mode       `json:"mode"`
		Difficulty *Mode       `json:"difficulty"`
		Language   string      `json:"language"`
		Stats      TypingStats `json:"stats"`
	}

	var alias sessionAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	s.Date = alias.Date
	s.Language = alias.Language
	s.Stats = alias.Stats

	switch {
	case alias.Mode != nil:
		s.Mode = *alias.Mode
	case alias.Difficulty != nil:
		s.Mode = *alias.Difficulty
	default:
		s.Mode = EasyCode
	}

	return nil
}

func isSymbol(r rune) bool {
	symbols := "!@#$%^&*()_+-=[]{}|;':\",./<>?`~\\"
	for _, s := range symbols {
		if r == s {
			return true
		}
	}
	return false
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
