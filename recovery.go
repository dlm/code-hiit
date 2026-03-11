package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
)

const (
	recoveryFileName    = "recovery.json"
	oldRecoveryFileName = ".typing-recovery.json"
)

// Motivational quotes (40% of content) - inspirational and uplifting
var motivationalQuotes = []RecoveryQuote{
	{Content: "Make it work, make it right, make it fast.", Author: "Kent Beck", Category: "motivational"},
	{Content: "The best way to predict the future is to invent it.", Author: "Alan Kay", Category: "motivational"},
	{Content: "First, solve the problem. Then, write the code.", Author: "John Johnson", Category: "motivational"},
	{Content: "Code is like humor. When you have to explain it, it's bad.", Author: "Cory House", Category: "motivational"},
	{Content: "Make each program do one thing well.", Author: "Doug McIlroy", Category: "motivational"},
	{Content: "Walk on water and develop software from a specification.", Author: "Edward Berard", Category: "motivational"},
	{Content: "Software is a great combination of artistry and engineering.", Author: "Bill Gates", Category: "motivational"},
	{Content: "The function of good software is to make the complex appear simple.", Author: "Grady Booch", Category: "motivational"},
	{Content: "Programs must be written for people to read.", Author: "Abelson and Sussman", Category: "motivational"},
	{Content: "Good code is its own best documentation.", Author: "Steve McConnell", Category: "motivational"},
	{Content: "Controlling complexity is the essence of computer programming.", Author: "Brian Kernighan", Category: "motivational"},
	{Content: "In programming the hard part isn't solving problems.", Author: "Chris Pine", Category: "motivational"},
	{Content: "It's not a bug. It's an undocumented feature.", Author: "Anonymous", Category: "motivational"},
	{Content: "The best programs are written so that computing machines can perform them.", Author: "Donald Knuth", Category: "motivational"},
	{Content: "Programming isn't about what you know, it's about what you figure out.", Author: "Chris Pine", Category: "motivational"},
	{Content: "Sometimes it pays to stay in bed on Monday.", Author: "Christopher Thompson", Category: "motivational"},
	{Content: "Perfection is achieved not when there is nothing more to add.", Author: "Antoine de Saint-Exupéry", Category: "motivational"},
	{Content: "Testing leads to failure, and failure leads to understanding.", Author: "Burt Rutan", Category: "motivational"},
	{Content: "Before software can be reusable it first has to be usable.", Author: "Ralph Johnson", Category: "motivational"},
}

// Educational/technical quotes (30% of content) - wisdom and best practices
var educationalQuotes = []RecoveryQuote{
	{Content: "Talk is cheap. Show me the code.", Author: "Linus Torvalds", Category: "educational"},
	{Content: "Premature optimization is the root of all evil.", Author: "Donald Knuth", Category: "educational"},
	{Content: "Any fool can write code that a computer can understand.", Author: "Martin Fowler", Category: "educational"},
	{Content: "Truth can only be found in one place: the code.", Author: "Robert C. Martin", Category: "educational"},
	{Content: "Always code as if the person who ends up maintaining your code.", Author: "Martin Golding", Category: "educational"},
	{Content: "There are only two hard things in computer science.", Author: "Phil Karlton", Category: "educational"},
	{Content: "Duplication is far cheaper than the wrong abstraction.", Author: "Sandi Metz", Category: "educational"},
	{Content: "Clarity and brevity sometimes are at odds.", Author: "Kernighan and Plauger", Category: "educational"},
	{Content: "The cheapest, fastest, and most reliable components are those that aren't there.", Author: "Gordon Bell", Category: "educational"},
	{Content: "The most effective debugging tool is still careful thought.", Author: "Brian Kernighan", Category: "educational"},
	{Content: "Make it correct, make it clear, make it concise, make it fast.", Author: "Wes Dyer", Category: "educational"},
	{Content: "Debugging is twice as hard as writing the code.", Author: "Brian Kernighan", Category: "educational"},
}

// Meditative mantras (30% of content) - calming and focusing
var meditativeMantras = []RecoveryQuote{
	{Content: "The journey of a thousand lines begins with a single character.", Author: "Anonymous", Category: "meditative"},
	{Content: "Precision is a practice, not a destination.", Author: "Anonymous", Category: "meditative"},
}

// Custom recovery quotes file structure
type RecoveryQuoteFile struct {
	Quotes []RecoveryQuote `json:"quotes"`
}

func getRecoveryQuotesPath() (string, error) {
	newPath, err := getConfigFile(recoveryFileName)
	if err != nil {
		return "", err
	}

	// Try to migrate from old location
	home, err := os.UserHomeDir()
	if err == nil {
		oldPath := filepath.Join(home, oldRecoveryFileName)
		migrateOldFile(oldPath, newPath)
	}

	return newPath, nil
}

// LoadCustomRecoveryQuotes loads user-defined recovery quotes
func LoadCustomRecoveryQuotes() ([]RecoveryQuote, error) {
	path, err := getRecoveryQuotesPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var file RecoveryQuoteFile
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, err
	}

	return file.Quotes, nil
}

// GetAllRecoveryQuotes returns all quotes (built-in + custom)
func GetAllRecoveryQuotes() []RecoveryQuote {
	all := make([]RecoveryQuote, 0)
	all = append(all, motivationalQuotes...)
	all = append(all, educationalQuotes...)
	all = append(all, meditativeMantras...)

	// Add custom quotes if available
	custom, err := LoadCustomRecoveryQuotes()
	if err == nil && len(custom) > 0 {
		all = append(all, custom...)
	}

	return all
}

// SelectWorkoutRecoveryQuote selects one quote for the entire workout
// This should be called once at workout start
func SelectWorkoutRecoveryQuote() RecoveryQuote {
	quotes := GetAllRecoveryQuotes()
	if len(quotes) == 0 {
		return RecoveryQuote{
			Content:  "Rest and recover.",
			Author:   "Anonymous",
			Category: "meditative",
		}
	}
	return quotes[rand.Intn(len(quotes))]
}

// GetRecoverySnippet converts a RecoveryQuote to a CodeSnippet for typing
// The user types only the Content, Author is displayed separately in UI
func GetRecoverySnippet(quote RecoveryQuote) CodeSnippet {
	return CodeSnippet{
		Content:  quote.Content,
		Language: "Recovery",
		Mode:     EasyCode, // Placeholder mode
	}
}
