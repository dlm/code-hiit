package main

import (
	"strings"
	"testing"
)

func TestSelectWorkoutRecoveryQuote(t *testing.T) {
	quote := SelectWorkoutRecoveryQuote()

	if quote.Content == "" {
		t.Error("Recovery quote content should not be empty")
	}

	if quote.Author == "" {
		t.Error("Recovery quote should have an author")
	}

	if quote.Category == "" {
		t.Error("Recovery quote should have a category")
	}

	// Verify category is valid
	validCategories := map[string]bool{
		"motivational": true,
		"educational":  true,
		"meditative":   true,
	}

	if !validCategories[quote.Category] {
		t.Errorf("Invalid category '%s', expected motivational, educational, or meditative", quote.Category)
	}
}

func TestGetRecoverySnippet(t *testing.T) {
	quote := RecoveryQuote{
		Content:  "Test quote content here.",
		Author:   "Test Author",
		Category: "motivational",
	}

	snippet := GetRecoverySnippet(quote)

	if snippet.Content != quote.Content {
		t.Errorf("Snippet content should match quote content, got '%s'", snippet.Content)
	}

	if snippet.Language != "Recovery" {
		t.Errorf("Expected language 'Recovery', got '%s'", snippet.Language)
	}
}

func TestGetAllRecoveryQuotes(t *testing.T) {
	quotes := GetAllRecoveryQuotes()

	if len(quotes) < len(motivationalQuotes)+len(educationalQuotes)+len(meditativeMantras) {
		t.Errorf("Expected to load all built-in quotes, got %d", len(quotes))
	}

	// Ensure at least one quote per category
	categories := map[string]bool{}
	for _, quote := range quotes {
		categories[quote.Category] = true
	}
	for _, category := range []string{"motivational", "educational", "meditative"} {
		if !categories[category] {
			t.Errorf("Expected at least one quote in category %s", category)
		}
	}
}

func TestQuoteWordCounts(t *testing.T) {
	quotes := GetAllRecoveryQuotes()

	for _, quote := range quotes {
		words := strings.Fields(quote.Content)
		wordCount := len(words)

		// Target is 7-12 words for ~10s recovery typing
		if wordCount < 7 || wordCount > 12 {
			t.Errorf("Quote '%s' has %d words, expected 7-12 words", quote.Content, wordCount)
		}
	}
}

func TestRecoveryQuoteVariety(t *testing.T) {
	// Generate multiple quotes and ensure we get variety
	seen := make(map[string]bool)

	for i := 0; i < 20; i++ {
		quote := SelectWorkoutRecoveryQuote()
		seen[quote.Content] = true
	}

	// Should have at least a few different quotes
	if len(seen) < 5 {
		t.Errorf("Expected variety in recovery quotes, only got %d unique quotes out of 20 selections", len(seen))
	}
}

func TestMotivationalQuotesCount(t *testing.T) {
	if len(motivationalQuotes) == 0 {
		t.Errorf("Expected at least one motivational quote")
	}
}

func TestEducationalQuotesCount(t *testing.T) {
	if len(educationalQuotes) == 0 {
		t.Errorf("Expected at least one educational quote")
	}
}

func TestMeditativeMantrasCount(t *testing.T) {
	if len(meditativeMantras) == 0 {
		t.Errorf("Expected at least one meditative mantra")
	}
}

func TestAllQuotesHaveRequiredFields(t *testing.T) {
	quotes := GetAllRecoveryQuotes()

	for i, quote := range quotes {
		if quote.Content == "" {
			t.Errorf("Quote %d has empty content", i)
		}
		if quote.Author == "" {
			t.Errorf("Quote %d (%s) has empty author", i, quote.Content)
		}
		if quote.Category == "" {
			t.Errorf("Quote %d (%s) has empty category", i, quote.Content)
		}
	}
}
