package main

import (
	"math/rand"
	"strings"
)

// WarmupPattern represents a typing warmup exercise
type WarmupPattern struct {
	Pattern     string
	Description string
	Category    string // "bigram", "trigram", "code"
}

// Common English bigrams for finger warmup
var bigramPatterns = []WarmupPattern{
	{Pattern: "th", Description: "Common bigram", Category: "bigram"},
	{Pattern: "he", Description: "Common bigram", Category: "bigram"},
	{Pattern: "in", Description: "Common bigram", Category: "bigram"},
	{Pattern: "er", Description: "Common bigram", Category: "bigram"},
	{Pattern: "an", Description: "Common bigram", Category: "bigram"},
	{Pattern: "re", Description: "Common bigram", Category: "bigram"},
	{Pattern: "on", Description: "Common bigram", Category: "bigram"},
	{Pattern: "at", Description: "Common bigram", Category: "bigram"},
	{Pattern: "st", Description: "Common bigram", Category: "bigram"},
	{Pattern: "en", Description: "Common bigram", Category: "bigram"},
	{Pattern: "nd", Description: "Common bigram", Category: "bigram"},
	{Pattern: "ti", Description: "Common bigram", Category: "bigram"},
	{Pattern: "es", Description: "Common bigram", Category: "bigram"},
	{Pattern: "or", Description: "Common bigram", Category: "bigram"},
	{Pattern: "te", Description: "Common bigram", Category: "bigram"},
	{Pattern: "of", Description: "Common bigram", Category: "bigram"},
	{Pattern: "ed", Description: "Common bigram", Category: "bigram"},
	{Pattern: "is", Description: "Common bigram", Category: "bigram"},
	{Pattern: "it", Description: "Common bigram", Category: "bigram"},
	{Pattern: "al", Description: "Common bigram", Category: "bigram"},
	{Pattern: "ar", Description: "Common bigram", Category: "bigram"},
}

// Common English trigrams
var trigramPatterns = []WarmupPattern{
	{Pattern: "the", Description: "Most common trigram", Category: "trigram"},
	{Pattern: "and", Description: "Common trigram", Category: "trigram"},
	{Pattern: "ing", Description: "Common trigram", Category: "trigram"},
	{Pattern: "ion", Description: "Common trigram", Category: "trigram"},
	{Pattern: "tio", Description: "Common trigram", Category: "trigram"},
	{Pattern: "ent", Description: "Common trigram", Category: "trigram"},
	{Pattern: "for", Description: "Common trigram", Category: "trigram"},
	{Pattern: "you", Description: "Common trigram", Category: "trigram"},
	{Pattern: "all", Description: "Common trigram", Category: "trigram"},
	{Pattern: "not", Description: "Common trigram", Category: "trigram"},
	{Pattern: "her", Description: "Common trigram", Category: "trigram"},
	{Pattern: "was", Description: "Common trigram", Category: "trigram"},
	{Pattern: "one", Description: "Common trigram", Category: "trigram"},
	{Pattern: "our", Description: "Common trigram", Category: "trigram"},
	{Pattern: "out", Description: "Common trigram", Category: "trigram"},
}

// Code-specific patterns for programming warmup
var codePatterns = []WarmupPattern{
	{Pattern: "->", Description: "Arrow operator", Category: "code"},
	{Pattern: "=>", Description: "Fat arrow", Category: "code"},
	{Pattern: "::", Description: "Scope resolution", Category: "code"},
	{Pattern: "==", Description: "Equality", Category: "code"},
	{Pattern: "!=", Description: "Not equal", Category: "code"},
	{Pattern: ">=", Description: "Greater or equal", Category: "code"},
	{Pattern: "<=", Description: "Less or equal", Category: "code"},
	{Pattern: "&&", Description: "Logical AND", Category: "code"},
	{Pattern: "||", Description: "Logical OR", Category: "code"},
	{Pattern: "++", Description: "Increment", Category: "code"},
	{Pattern: "--", Description: "Decrement", Category: "code"},
	{Pattern: "<<", Description: "Left shift", Category: "code"},
	{Pattern: ">>", Description: "Right shift", Category: "code"},
	{Pattern: "...", Description: "Spread/rest", Category: "code"},
	{Pattern: "<-", Description: "Channel operator", Category: "code"},
	{Pattern: "//", Description: "Comment", Category: "code"},
	{Pattern: "/*", Description: "Block comment start", Category: "code"},
	{Pattern: "*/", Description: "Block comment end", Category: "code"},
	{Pattern: "()", Description: "Parentheses", Category: "code"},
	{Pattern: "{}", Description: "Braces", Category: "code"},
	{Pattern: "[]", Description: "Brackets", Category: "code"},
	{Pattern: "<>", Description: "Angle brackets", Category: "code"},
}

// GetWarmupSnippet generates a warmup exercise snippet
func GetWarmupSnippet() CodeSnippet {
	// Mix of bigrams, trigrams, and code patterns
	patterns := selectWarmupPatterns()
	content := buildWarmupContent(patterns)

	return CodeSnippet{
		Content:  content,
		Language: "Warmup",
		Mode:     EasyCode, // Placeholder mode for warmup
	}
}

// selectWarmupPatterns chooses a mix of patterns for warmup
func selectWarmupPatterns() []WarmupPattern {
	selected := []WarmupPattern{}

	// Select 2-3 bigrams
	bigramCount := 2 + rand.Intn(2)
	for i := 0; i < bigramCount; i++ {
		selected = append(selected, bigramPatterns[rand.Intn(len(bigramPatterns))])
	}

	// Select 1-2 trigrams
	trigramCount := 1 + rand.Intn(2)
	for i := 0; i < trigramCount; i++ {
		selected = append(selected, trigramPatterns[rand.Intn(len(trigramPatterns))])
	}

	// Select 1-2 code patterns
	codeCount := 1 + rand.Intn(2)
	for i := 0; i < codeCount; i++ {
		selected = append(selected, codePatterns[rand.Intn(len(codePatterns))])
	}

	return selected
}

// buildWarmupContent creates repeated pattern content
func buildWarmupContent(patterns []WarmupPattern) string {
	var lines []string

	for _, pattern := range patterns {
		// Repeat each pattern 3-5 times with spaces
		repeatCount := 3 + rand.Intn(3)
		repeated := make([]string, repeatCount)
		for i := 0; i < repeatCount; i++ {
			repeated[i] = pattern.Pattern
		}
		line := strings.Join(repeated, " ")
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

// GetTargetedWarmup generates warmup focusing on specific pattern type
func GetTargetedWarmup(category string) CodeSnippet {
	var patterns []WarmupPattern

	switch category {
	case "bigram":
		// Select 4-6 random bigrams
		count := 4 + rand.Intn(3)
		for i := 0; i < count; i++ {
			patterns = append(patterns, bigramPatterns[rand.Intn(len(bigramPatterns))])
		}
	case "trigram":
		// Select 3-5 random trigrams
		count := 3 + rand.Intn(3)
		for i := 0; i < count; i++ {
			patterns = append(patterns, trigramPatterns[rand.Intn(len(trigramPatterns))])
		}
	case "code":
		// Select 4-6 random code patterns
		count := 4 + rand.Intn(3)
		for i := 0; i < count; i++ {
			patterns = append(patterns, codePatterns[rand.Intn(len(codePatterns))])
		}
	default:
		// Default to mixed
		return GetWarmupSnippet()
	}

	content := buildWarmupContent(patterns)

	return CodeSnippet{
		Content:  content,
		Language: "Warmup - " + category,
		Mode:     EasyCode,
	}
}
