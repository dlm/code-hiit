package main

import (
	"strings"
	"testing"
)

func TestGetWarmupSnippet(t *testing.T) {
	snippet := GetWarmupSnippet()

	if snippet.Content == "" {
		t.Error("Warmup snippet content should not be empty")
	}

	if snippet.Language != "Warmup" {
		t.Errorf("Expected language 'Warmup', got '%s'", snippet.Language)
	}

	// Should contain multiple lines
	lines := strings.Split(snippet.Content, "\n")
	if len(lines) < 3 {
		t.Errorf("Expected at least 3 lines of warmup content, got %d", len(lines))
	}

	// Each line should have repeated patterns
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 3 {
			t.Errorf("Line %d should have at least 3 repeated patterns, got %d", i, len(parts))
		}
	}
}

func TestGetTargetedWarmupBigram(t *testing.T) {
	snippet := GetTargetedWarmup("bigram")

	if snippet.Content == "" {
		t.Error("Bigram warmup should not be empty")
	}

	if !strings.Contains(snippet.Language, "bigram") {
		t.Errorf("Expected language to contain 'bigram', got '%s'", snippet.Language)
	}

	// Check that content contains valid bigrams
	for _, pattern := range bigramPatterns {
		if strings.Contains(snippet.Content, pattern.Pattern) {
			return // Found at least one bigram
		}
	}
	t.Error("Bigram warmup should contain at least one bigram pattern")
}

func TestGetTargetedWarmupTrigram(t *testing.T) {
	snippet := GetTargetedWarmup("trigram")

	if snippet.Content == "" {
		t.Error("Trigram warmup should not be empty")
	}

	if !strings.Contains(snippet.Language, "trigram") {
		t.Errorf("Expected language to contain 'trigram', got '%s'", snippet.Language)
	}

	// Check that content contains valid trigrams
	for _, pattern := range trigramPatterns {
		if strings.Contains(snippet.Content, pattern.Pattern) {
			return // Found at least one trigram
		}
	}
	t.Error("Trigram warmup should contain at least one trigram pattern")
}

func TestGetTargetedWarmupCode(t *testing.T) {
	snippet := GetTargetedWarmup("code")

	if snippet.Content == "" {
		t.Error("Code warmup should not be empty")
	}

	if !strings.Contains(snippet.Language, "code") {
		t.Errorf("Expected language to contain 'code', got '%s'", snippet.Language)
	}

	// Check that content contains valid code patterns
	for _, pattern := range codePatterns {
		if strings.Contains(snippet.Content, pattern.Pattern) {
			return // Found at least one code pattern
		}
	}
	t.Error("Code warmup should contain at least one code pattern")
}

func TestBuildWarmupContent(t *testing.T) {
	patterns := []WarmupPattern{
		{Pattern: "th", Description: "test", Category: "bigram"},
		{Pattern: "->", Description: "test", Category: "code"},
	}

	content := buildWarmupContent(patterns)

	if content == "" {
		t.Error("Built content should not be empty")
	}

	// Should contain both patterns
	if !strings.Contains(content, "th") {
		t.Error("Content should contain 'th' pattern")
	}

	if !strings.Contains(content, "->") {
		t.Error("Content should contain '->' pattern")
	}

	// Should have multiple lines
	if !strings.Contains(content, "\n") {
		t.Error("Content should have multiple lines")
	}
}

func TestWarmupPatternVariety(t *testing.T) {
	// Generate multiple warmups and ensure we get variety
	seen := make(map[string]bool)

	for i := 0; i < 10; i++ {
		snippet := GetWarmupSnippet()
		seen[snippet.Content] = true
	}

	// Should have at least a few different variations
	if len(seen) < 3 {
		t.Errorf("Expected variety in warmup generation, only got %d unique snippets out of 10", len(seen))
	}
}
