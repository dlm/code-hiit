package main

import (
	"fmt"
	"math/rand"
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

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

var modeOptions = []Mode{EasyCode, MediumCode, HardCode, NumbersPractice, SymbolsPractice, HexNumbers, BracketsPractice, RegexPatterns, Custom}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func modeName(m Mode) string {
	switch m {
	case EasyCode:
		return "Easy Code"
	case MediumCode:
		return "Medium Code"
	case HardCode:
		return "Hard Code"
	case NumbersPractice:
		return "Numbers Practice"
	case HexNumbers:
		return "Hex Numbers"
	case SymbolsPractice:
		return "Symbols Practice"
	case BracketsPractice:
		return "Brackets Practice"
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
	m.filteredModes = []Mode{}
	for _, mode := range modeOptions {
		if fuzzyMatch(m.searchInput, modeName(mode)) {
			m.filteredModes = append(m.filteredModes, mode)
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
	mode          Mode
	snippet       CodeSnippet
	snippetIndex  int
	typedText     string
	currentPos    int
	stats         TypingStats
	history       *SessionHistory
	menuCursor    int
	searchInput   string
	filteredModes []Mode
	err           error
	// HIIT workout mode
	workoutState *WorkoutState
	isHIITMode   bool
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
		filteredModes: modeOptions,
	}
}

func (m model) Init() tea.Cmd {
	if m.isHIITMode && m.workoutState != nil {
		return tickCmd()
	}
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
	case tickMsg:
		if m.isHIITMode && m.workoutState != nil && !m.workoutState.Paused {
			if m.workoutState.RemainingTime() <= 0 {
				return m.transitionPhase()
			}
		}
		if m.isHIITMode {
			return m, tickCmd()
		}
	}
	return m, nil
}

func (m model) transitionPhase() (tea.Model, tea.Cmd) {
	ws := m.workoutState

	// Finalize current phase stats
	m.stats.EndTime = time.Now()

	// Save current phase stats
	currentPhaseStats := PhaseStats{
		Phase:     ws.CurrentPhase,
		Stats:     m.stats,
		Completed: true,
	}

	// Determine next phase
	var nextPhase WorkoutPhase
	var shouldSaveSet bool

	switch ws.CurrentPhase {
	case WarmupPhase:
		nextPhase = WorkPhase
	case WorkPhase:
		nextPhase = RecoveryPhase
	case RecoveryPhase:
		// Complete current set
		shouldSaveSet = true
		ws.CurrentSet++

		// Check if workout is complete
		if ws.CurrentSet >= ws.Workout.TotalSets {
			// Persist the final recovery stats before finishing
			setIndex := ws.CurrentSet - 1
			for len(ws.Workout.Sets) <= setIndex {
				ws.Workout.Sets = append(ws.Workout.Sets, NewSetStats(setIndex+1))
			}
			ws.Workout.Sets[setIndex].Recovery = currentPhaseStats

			m.completeHIITWorkout()
			return m, nil
		}

		// Start next set with warmup
		nextPhase = WarmupPhase
	}

	// Save the set if recovery just completed
	if shouldSaveSet {
		// Find or create the set stats for the current set
		setIndex := ws.CurrentSet - 1
		for len(ws.Workout.Sets) <= setIndex {
			ws.Workout.Sets = append(ws.Workout.Sets, NewSetStats(len(ws.Workout.Sets)+1))
		}

		// Add phase stats to appropriate field in set
		switch currentPhaseStats.Phase {
		case WarmupPhase:
			ws.Workout.Sets[setIndex].Warmup = currentPhaseStats
		case WorkPhase:
			ws.Workout.Sets[setIndex].Work = currentPhaseStats
		case RecoveryPhase:
			ws.Workout.Sets[setIndex].Recovery = currentPhaseStats
		}
	} else {
		// Save phase stats to current set
		setIndex := ws.CurrentSet
		for len(ws.Workout.Sets) <= setIndex {
			ws.Workout.Sets = append(ws.Workout.Sets, NewSetStats(len(ws.Workout.Sets)+1))
		}

		switch currentPhaseStats.Phase {
		case WarmupPhase:
			ws.Workout.Sets[setIndex].Warmup = currentPhaseStats
		case WorkPhase:
			ws.Workout.Sets[setIndex].Work = currentPhaseStats
		}
	}

	// Start next phase
	ws.StartPhase(nextPhase, time.Now())

	// Load appropriate snippet for next phase
	switch nextPhase {
	case WarmupPhase:
		ws.CurrentSnippet = GetWarmupSnippet()
	case WorkPhase:
		ws.CurrentSnippet = GetRandomSnippet(ws.Workout.FocusMode)
	case RecoveryPhase:
		ws.CurrentSnippet = GetRecoverySnippet(ws.RecoveryQuote)
	}

	// Reset typing state for new phase
	m.typedText = ""
	m.currentPos = 0
	m.snippet = ws.CurrentSnippet
	m.stats = TypingStats{
		StartTime: time.Now(),
		Errors:    []TypingError{},
	}

	return m, nil
}

func (m *model) completeHIITWorkout() {
	ws := m.workoutState
	ws.Workout.EndTime = time.Now()
	ws.Workout.Completed = true

	// Save to history
	m.history.HIITWorkouts = append(m.history.HIITWorkouts, *ws.Workout)
	SaveHistory(m.history)

	m.state = stateResults
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
			m.mode = m.filteredModes[m.menuCursor]
			m.snippet = GetRandomSnippet(m.mode)
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
		if m.isHIITMode {
			// In HIIT mode, esc quits the workout
			return m, tea.Quit
		}
		m.state = stateMenu
		return m, nil
	case " ":
		// Space bar toggles pause in HIIT mode
		if m.isHIITMode && m.workoutState != nil {
			if m.workoutState.Paused {
				// Resume
				pauseDuration := time.Since(m.workoutState.PausedAt)
				m.workoutState.PhasePausedDuration += pauseDuration
				m.workoutState.Paused = false
			} else {
				// Pause
				m.workoutState.Paused = true
				m.workoutState.PausedAt = time.Now()
			}
			return m, nil
		}
		// In freeform mode, space is a regular character
		if !m.isHIITMode {
			m = m.processChar(' ')
		}
	case "ctrl+n", "ctrl+s":
		// Only allow skip snippet in freeform mode
		if !m.isHIITMode {
			var newIndex int
			m.snippet, newIndex = GetNextSnippet(m.mode, m.snippetIndex)
			m.snippetIndex = newIndex
			m.typedText = ""
			m.currentPos = 0
			m.stats = TypingStats{
				StartTime: time.Now(),
				Errors:    []TypingError{},
			}
			return m, nil
		}
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
		Date:     time.Now(),
		Mode:     m.mode,
		Language: m.snippet.Language,
		Stats:    m.stats,
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
		m.snippet = GetRandomSnippet(m.mode)
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

	warmupStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			Bold(true)

	workStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	recoveryStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("78")).
			Bold(true)

	timerBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1).
			MarginBottom(1)
)

func renderProgressBar(progress float64, width int) string {
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}
	filled := int(progress * float64(width))
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func formatTime(seconds int) string {
	if seconds < 0 {
		seconds = 0
	}
	mins := seconds / 60
	secs := seconds % 60
	return fmt.Sprintf("%02d:%02d", mins, secs)
}

func (m model) renderPhaseTimer() string {
	if m.workoutState == nil {
		return ""
	}

	ws := m.workoutState
	var phaseStyle lipgloss.Style
	var phaseName string

	switch ws.CurrentPhase {
	case WarmupPhase:
		phaseStyle = warmupStyle
		phaseName = "WARMUP"
	case WorkPhase:
		phaseStyle = workStyle
		phaseName = "WORK"
	case RecoveryPhase:
		phaseStyle = recoveryStyle
		phaseName = "RECOVERY"
	}

	var b strings.Builder

	// Phase header with set number
	setInfo := fmt.Sprintf("%s - Set %d/%d", phaseName, ws.CurrentSet+1, ws.Workout.TotalSets)
	b.WriteString(phaseStyle.Render(setInfo))
	b.WriteString("\n")

	// Timer display
	remaining := ws.RemainingTime()
	if ws.Paused {
		b.WriteString(subtitleStyle.Render(fmt.Sprintf("⏸  PAUSED - %s remaining", formatTime(remaining))))
	} else {
		b.WriteString(fmt.Sprintf("⏱  %s remaining", formatTime(remaining)))
	}
	b.WriteString("\n")

	// Progress bar
	progress := ws.Progress()
	progressBar := renderProgressBar(progress, 40)
	b.WriteString(progressBar)
	b.WriteString(fmt.Sprintf(" %.0f%%", progress*100))

	return timerBoxStyle.Render(b.String())
}

func (m model) viewMenu() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("code-hiit"))
	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render("Select Mode:"))
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
		for i, mode := range m.filteredModes {
			cursor := " "
			if m.menuCursor == i {
				cursor = ">"
				b.WriteString(selectedStyle.Render(fmt.Sprintf(" %s %s", cursor, modeName(mode))))
			} else {
				b.WriteString(fmt.Sprintf(" %s %s", cursor, modeName(mode)))
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

	// Show HIIT timer if in HIIT mode
	if m.isHIITMode && m.workoutState != nil {
		b.WriteString(m.renderPhaseTimer())
		b.WriteString("\n")

		// Add phase-specific message
		var phaseMsg string
		switch m.workoutState.CurrentPhase {
		case WarmupPhase:
			phaseMsg = "Get your fingers moving with basic patterns..."
		case WorkPhase:
			phaseMsg = "PUSH! Maximum intensity!"
		case RecoveryPhase:
			phaseMsg = "Breathe... Reflect on the wisdom below..."
		}
		b.WriteString(subtitleStyle.Render(phaseMsg))
		b.WriteString("\n\n")
	} else {
		// Regular freeform mode header
		header := fmt.Sprintf("%s - %s", modeName(m.mode), m.snippet.Language)
		b.WriteString(titleStyle.Render(header))
		b.WriteString("\n\n")
	}

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

	// Show stats only in freeform mode (HIIT mode has timer instead)
	if !m.isHIITMode {
		progress := float64(m.currentPos) / float64(len(m.snippet.Content)) * 100
		b.WriteString(statsStyle.Render(fmt.Sprintf("\nProgress: %.0f%%", progress)))

		if m.stats.TotalChars > 0 {
			b.WriteString(statsStyle.Render(fmt.Sprintf(" | Accuracy: %.1f%%", m.stats.Accuracy())))
		}
	}

	b.WriteString("\n\n")

	// Show different controls for HIIT vs freeform mode
	if m.isHIITMode {
		b.WriteString(subtitleStyle.Render("space: pause/resume • esc: quit workout"))
	} else {
		b.WriteString(subtitleStyle.Render("ctrl+n: next snippet • esc: menu • ctrl+c: quit"))
	}

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

func demoHIITTimer() model {
	history, err := LoadHistory()
	if err != nil {
		history = &SessionHistory{Sessions: []Session{}}
	}

	// Create a Quick workout (3 sets) with EasyCode mode
	workout := &HIITWorkout{
		WorkoutType: QuickWorkout,
		FocusMode:   EasyCode,
		TotalSets:   QuickWorkout.Sets(),
		StartTime:   time.Now(),
		Sets:        []SetStats{},
	}

	// Select recovery quote once at start
	recoveryQuote := SelectWorkoutRecoveryQuote()

	// Initialize workout state
	workoutState := &WorkoutState{
		Workout:        workout,
		CurrentSet:     0,
		CurrentPhase:   WarmupPhase,
		RecoveryQuote:  recoveryQuote,
		Paused:         false,
	}

	// Start warmup phase
	workoutState.StartPhase(WarmupPhase, time.Now())
	workoutState.CurrentSnippet = GetWarmupSnippet()

	// Create model with HIIT mode enabled
	return model{
		state:        stateTyping,
		mode:         EasyCode,
		snippet:      workoutState.CurrentSnippet,
		typedText:    "",
		currentPos:   0,
		stats:        TypingStats{StartTime: time.Now(), Errors: []TypingError{}},
		history:      history,
		isHIITMode:   true,
		workoutState: workoutState,
	}
}

func main() {
	// Check if we should run HIIT demo
	if len(os.Args) > 1 && os.Args[1] == "--demo-hiit" {
		p := tea.NewProgram(demoHIITTimer())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Normal mode
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
