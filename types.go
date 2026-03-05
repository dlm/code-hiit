package main

import "time"

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
	Numbers
	Symbols
	HexNumbers
	Brackets
	RegexPatterns
)

type CodeSnippet struct {
	Content    string
	Language   string
	Difficulty Difficulty
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
	Date       time.Time   `json:"date"`
	Difficulty Difficulty  `json:"difficulty"`
	Language   string      `json:"language"`
	Stats      TypingStats `json:"stats"`
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
