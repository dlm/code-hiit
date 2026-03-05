package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	stateMenu state = iota
	stateTyping
	stateResults
)

var difficultyOptions = []Difficulty{Easy, Medium, Hard, Numbers, Symbols, HexNumbers, Brackets, RegexPatterns, Custom}

func difficultyName(d Difficulty) string {
	switch d {
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	case Numbers:
		return "Numbers"
	case HexNumbers:
		return "Hex Numbers"
	case Symbols:
		return "Symbols"
	case Brackets:
		return "Brackets"
	case RegexPatterns:
		return "Regex Patterns"
	case Custom:
		return "Custom"
	default:
		return "Unknown"
	}
}

func fuzzyMatch(query, target string) bool {
	if query == "" {
		return true
	}
	query = strings.ToLower(query)
	target = strings.ToLower(target)

	qi := 0
	for _, tc := range target {
		if qi < len(query) && rune(query[qi]) == tc {
			qi++
		}
	}
	return qi == len(query)
}

func (m *model) updateFilteredModes() {
	m.filteredModes = []Difficulty{}
	for _, diff := range difficultyOptions {
		if fuzzyMatch(m.searchInput, difficultyName(diff)) {
			m.filteredModes = append(m.filteredModes, diff)
		}
	}
	if m.menuCursor >= len(m.filteredModes) {
		m.menuCursor = len(m.filteredModes) - 1
	}
	if m.menuCursor < 0 {
		m.menuCursor = 0
	}
}

type model struct {
	state         state
	difficulty    Difficulty
	snippet       CodeSnippet
	snippetIndex  int
	typedText     string
	currentPos    int
	stats         TypingStats
	history       *SessionHistory
	menuCursor    int
	searchInput   string
	filteredModes []Difficulty
	err           error
}

func initialModel() model {
	history, err := LoadHistory()
	if err != nil {
		history = &SessionHistory{Sessions: []Session{}}
	}
	return model{
		state:         stateMenu,
		menuCursor:    0,
		history:       history,
		searchInput:   "",
		filteredModes: difficultyOptions,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case stateMenu:
			return m.updateMenu(msg)
		case stateTyping:
			return m.updateTyping(msg)
		case stateResults:
			return m.updateResults(msg)
		}
	}
	return m, nil
}

func (m model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "ctrl+p":
		if m.menuCursor > 0 {
			m.menuCursor--
		}
	case "down", "ctrl+n":
		if m.menuCursor < len(m.filteredModes)-1 {
			m.menuCursor++
		}
	case "enter":
		if len(m.filteredModes) > 0 {
			m.difficulty = m.filteredModes[m.menuCursor]
			m.snippet = GetRandomSnippet(m.difficulty)
			m.snippetIndex = 0
			m.typedText = ""
			m.currentPos = 0
			m.stats = TypingStats{
				StartTime: time.Now(),
				Errors:    []TypingError{},
			}
			m.state = stateTyping
		}
	case "backspace":
		if len(m.searchInput) > 0 {
			m.searchInput = m.searchInput[:len(m.searchInput)-1]
			m.updateFilteredModes()
		}
	case "esc":
		if m.searchInput != "" {
			m.searchInput = ""
			m.menuCursor = 0
			m.updateFilteredModes()
		} else {
			return m, tea.Quit
		}
	default:
		key := msg.String()
		if len(key) == 1 {
			m.searchInput += key
			m.updateFilteredModes()
		}
	}
	return m, nil
}

func (m model) updateTyping(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = stateMenu
		return m, nil
	case "ctrl+n", "ctrl+s":
		var newIndex int
		m.snippet, newIndex = GetNextSnippet(m.difficulty, m.snippetIndex)
		m.snippetIndex = newIndex
		m.typedText = ""
		m.currentPos = 0
		m.stats = TypingStats{
			StartTime: time.Now(),
			Errors:    []TypingError{},
		}
		return m, nil
	case "backspace":
		if len(m.typedText) > 0 {
			m.typedText = m.typedText[:len(m.typedText)-1]
			if m.currentPos > 0 {
				m.currentPos--
			}
		}
	case "enter":
		m = m.processChar('\n')
	case "tab":
		m = m.processChar('\t')
	default:
		key := msg.String()
		if len(key) == 1 {
			char := rune(key[0])
			m = m.processChar(char)
		}
	}
	return m, nil
}

func (m model) processChar(char rune) model {
	if m.currentPos >= len(m.snippet.Content) {
		return m
	}

	// Handle Tab -> 4 spaces conversion
	if char == '\t' {
		// Check if the next 4 characters are spaces
		if m.currentPos+4 <= len(m.snippet.Content) {
			nextFour := m.snippet.Content[m.currentPos : m.currentPos+4]
			if nextFour == "    " {
				// Treat tab as 4 spaces
				m.typedText += "    "
				for i := 0; i < 4; i++ {
					m.stats.TotalChars++
					m.stats.CorrectChars++
					m.currentPos++
				}
				if m.currentPos >= len(m.snippet.Content) {
					m.completeSession()
				}
				return m
			}
		}
	}

	m.typedText += string(char)
	expected := rune(m.snippet.Content[m.currentPos])

	m.stats.TotalChars++
	if char == expected {
		m.stats.CorrectChars++
	} else {
		m.stats.IncorrectChars++
		m.stats.Errors = append(m.stats.Errors, TypingError{
			Position: m.currentPos,
			Expected: expected,
			Actual:   char,
		})
	}

	if isSymbol(expected) {
		m.stats.TotalSymbols++
		if char == expected {
			m.stats.CorrectSymbols++
		}
	}

	if isNumber(expected) {
		m.stats.TotalNumbers++
		if char == expected {
			m.stats.CorrectNumbers++
		}
	}

	m.currentPos++

	if m.currentPos >= len(m.snippet.Content) {
		m.completeSession()
	}

	return m
}

func (m *model) completeSession() {
	m.stats.EndTime = time.Now()
	session := Session{
		Date:       time.Now(),
		Difficulty: m.difficulty,
		Language:   m.snippet.Language,
		Stats:      m.stats,
	}
	m.history.AddSession(session)
	SaveHistory(m.history)
	m.state = stateResults
}

func (m model) updateResults(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "enter", "r":
		m.snippet = GetRandomSnippet(m.difficulty)
		m.typedText = ""
		m.currentPos = 0
		m.stats = TypingStats{
			StartTime: time.Now(),
			Errors:    []TypingError{},
		}
		m.state = stateTyping
	case "m":
		m.state = stateMenu
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case stateMenu:
		return m.viewMenu()
	case stateTyping:
		return m.viewTyping()
	case stateResults:
		return m.viewResults()
	}
	return ""
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	correctStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	incorrectStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Background(lipgloss.Color("52"))

	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("240"))

	statsStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			MarginTop(1)
)

func (m model) viewMenu() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Code Typing Trainer"))
	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render("Select mode:"))
	b.WriteString("\n\n")

	// Show search input
	searchPrompt := "> "
	if m.searchInput == "" {
		b.WriteString(subtitleStyle.Render(searchPrompt + "_"))
	} else {
		b.WriteString(subtitleStyle.Render(searchPrompt + m.searchInput + "_"))
	}
	b.WriteString("\n\n")

	// Show filtered modes
	if len(m.filteredModes) == 0 {
		b.WriteString(subtitleStyle.Render("  No matches found"))
		b.WriteString("\n")
	} else {
		for i, diff := range m.filteredModes {
			cursor := " "
			if m.menuCursor == i {
				cursor = ">"
				b.WriteString(selectedStyle.Render(fmt.Sprintf(" %s %s", cursor, difficultyName(diff))))
			} else {
				b.WriteString(fmt.Sprintf(" %s %s", cursor, difficultyName(diff)))
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render("type to search • ↑/↓/ctrl-n/ctrl-p: navigate • enter: select • esc: clear/quit"))

	return b.String()
}

func (m model) viewTyping() string {
	var b strings.Builder

	header := fmt.Sprintf("%s - %s", difficultyName(m.difficulty), m.snippet.Language)
	b.WriteString(titleStyle.Render(header))
	b.WriteString("\n\n")

	for i, char := range m.snippet.Content {
		isCursor := i == m.currentPos
		isTyped := i < m.currentPos

		// Determine if this character was typed correctly
		var isCorrect bool
		if isTyped && i < len(m.typedText) {
			isCorrect = rune(m.typedText[i]) == char
		}

		// Render the character based on its type and state
		if char == '\n' {
			if isCursor {
				b.WriteString(cursorStyle.Render("↵"))
			} else if isTyped {
				if isCorrect {
					b.WriteString(correctStyle.Render("↵"))
				} else {
					b.WriteString(incorrectStyle.Render("↵"))
				}
			}
			b.WriteString("\n")
		} else if char == '\t' {
			if isCursor {
				b.WriteString(cursorStyle.Render("→   "))
			} else if isTyped {
				if isCorrect {
					b.WriteString(correctStyle.Render("→") + "   ")
				} else {
					b.WriteString(incorrectStyle.Render("→") + "   ")
				}
			} else {
				b.WriteString("    ")
			}
		} else {
			// Regular character
			if isCursor {
				b.WriteString(cursorStyle.Render(string(char)))
			} else if isTyped {
				if isCorrect {
					b.WriteString(correctStyle.Render(string(char)))
				} else {
					b.WriteString(incorrectStyle.Render(string(char)))
				}
			} else {
				b.WriteString(string(char))
			}
		}
	}

	b.WriteString("\n")
	progress := float64(m.currentPos) / float64(len(m.snippet.Content)) * 100
	b.WriteString(statsStyle.Render(fmt.Sprintf("\nProgress: %.0f%%", progress)))

	if m.stats.TotalChars > 0 {
		b.WriteString(statsStyle.Render(fmt.Sprintf(" | Accuracy: %.1f%%", m.stats.Accuracy())))
	}

	b.WriteString("\n\n")
	b.WriteString(subtitleStyle.Render("ctrl+n: next snippet • esc: menu • ctrl+c: quit"))

	return b.String()
}

func (m model) viewResults() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Results"))
	b.WriteString("\n\n")

	b.WriteString(fmt.Sprintf("WPM:              %.1f\n", m.stats.WPM()))
	b.WriteString(fmt.Sprintf("Accuracy:         %.1f%%\n", m.stats.Accuracy()))
	b.WriteString(fmt.Sprintf("Total Characters: %d\n", m.stats.TotalChars))
	b.WriteString(fmt.Sprintf("Errors:           %d\n", m.stats.IncorrectChars))
	b.WriteString("\n")

	if m.stats.TotalSymbols > 0 {
		b.WriteString(statsStyle.Render(fmt.Sprintf("Symbol Accuracy:  %.1f%% (%d/%d)\n",
			m.stats.SymbolAccuracy(), m.stats.CorrectSymbols, m.stats.TotalSymbols)))
	}

	if m.stats.TotalNumbers > 0 {
		b.WriteString(statsStyle.Render(fmt.Sprintf("Number Accuracy:  %.1f%% (%d/%d)\n",
			m.stats.NumberAccuracy(), m.stats.CorrectNumbers, m.stats.TotalNumbers)))
	}

	b.WriteString("\n")
	duration := m.stats.EndTime.Sub(m.stats.StartTime)
	b.WriteString(fmt.Sprintf("Time:             %.1fs\n", duration.Seconds()))

	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render("r/enter: retry • m: menu • q: quit"))

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
