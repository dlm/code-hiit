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

// WorkoutType represents predefined HIIT workout configurations
type WorkoutType int

const (
	QuickWorkout    WorkoutType = iota // 3 sets
	StandardWorkout                    // 5 sets
	ExtendedWorkout                    // 8 sets
)

func (wt WorkoutType) Sets() int {
	switch wt {
	case QuickWorkout:
		return 3
	case StandardWorkout:
		return 5
	case ExtendedWorkout:
		return 8
	default:
		return 0
	}
}

func (wt WorkoutType) String() string {
	switch wt {
	case QuickWorkout:
		return "Quick (3 sets)"
	case StandardWorkout:
		return "Standard (5 sets)"
	case ExtendedWorkout:
		return "Extended (8 sets)"
	default:
		return "Unknown"
	}
}

// WorkoutPhase represents the three phases of a HIIT set
type WorkoutPhase int

const (
	WarmupPhase WorkoutPhase = iota
	WorkPhase
	RecoveryPhase
)

func (wp WorkoutPhase) String() string {
	switch wp {
	case WarmupPhase:
		return "Warmup"
	case WorkPhase:
		return "Work"
	case RecoveryPhase:
		return "Recovery"
	default:
		return "Unknown"
	}
}

func (wp WorkoutPhase) Duration() int {
	switch wp {
	case WarmupPhase:
		return 15 // seconds
	case WorkPhase:
		return 30 // seconds
	case RecoveryPhase:
		return 10 // seconds
	default:
		return 0
	}
}

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

// PhaseStats tracks statistics for a single HIIT phase (warmup/work/recovery)
type PhaseStats struct {
	Phase        WorkoutPhase `json:"phase"`
	Stats        TypingStats  `json:"stats"`
	Completed    bool         `json:"completed"`
	SkippedEarly bool         `json:"skipped_early"`
}

// SetStats tracks statistics for one complete HIIT set (all 3 phases)
type SetStats struct {
	SetNumber int        `json:"set_number"`
	Warmup    PhaseStats `json:"warmup"`
	Work      PhaseStats `json:"work"`
	Recovery  PhaseStats `json:"recovery"`
}

func NewSetStats(setNumber int) SetStats {
	return SetStats{
		SetNumber: setNumber,
		Warmup:    PhaseStats{Phase: WarmupPhase},
		Work:      PhaseStats{Phase: WorkPhase},
		Recovery:  PhaseStats{Phase: RecoveryPhase},
	}
}

// HIITWorkout represents a complete HIIT workout session
type HIITWorkout struct {
	WorkoutType WorkoutType `json:"workout_type"`
	FocusMode   Mode        `json:"focus_mode"` // Mode used for work phase
	Sets        []SetStats  `json:"sets"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	Completed   bool        `json:"completed"`
	TotalSets   int         `json:"total_sets"`
}

func (hw *HIITWorkout) CompletedSets() int {
	count := 0
	for _, set := range hw.Sets {
		if set.Warmup.Completed && set.Work.Completed && set.Recovery.Completed {
			count++
		}
	}
	return count
}

func (hw *HIITWorkout) AverageWPM(phase WorkoutPhase) float64 {
	total := 0.0
	count := 0
	for _, set := range hw.Sets {
		var stats *PhaseStats
		switch phase {
		case WarmupPhase:
			stats = &set.Warmup
		case WorkPhase:
			stats = &set.Work
		case RecoveryPhase:
			stats = &set.Recovery
		}
		if stats != nil && stats.Completed {
			total += stats.Stats.WPM()
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return total / float64(count)
}

func (hw *HIITWorkout) AverageAccuracy(phase WorkoutPhase) float64 {
	total := 0.0
	count := 0
	for _, set := range hw.Sets {
		var stats *PhaseStats
		switch phase {
		case WarmupPhase:
			stats = &set.Warmup
		case WorkPhase:
			stats = &set.Work
		case RecoveryPhase:
			stats = &set.Recovery
		}
		if stats != nil && stats.Completed {
			total += stats.Stats.Accuracy()
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return total / float64(count)
}

type SessionHistory struct {
	Sessions     []Session     `json:"sessions"`
	HIITWorkouts []HIITWorkout `json:"hiit_workouts"`
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

// RecoveryQuote represents a quote used during recovery phase
type RecoveryQuote struct {
	Content  string `json:"content"`
	Author   string `json:"author"`
	Category string `json:"category"` // "motivational", "educational", "meditative"
}

// WorkoutState represents the current state during a HIIT workout
type WorkoutState struct {
	Workout             *HIITWorkout
	CurrentSet          int
	CurrentPhase        WorkoutPhase
	PhaseStartTime      time.Time
	PhaseEndTime        time.Time
	CurrentSnippet      CodeSnippet
	SnippetIndex        int
	TypedText           string
	CurrentPos          int
	CurrentStats        TypingStats
	RecoveryQuote       RecoveryQuote // Quote selected for this workout (same for all recovery phases)
	Paused              bool
	PausedAt            time.Time
	PhasePausedDuration time.Duration
}

func (ws *WorkoutState) StartPhase(phase WorkoutPhase, startTime time.Time) {
	ws.CurrentPhase = phase
	ws.PhaseStartTime = startTime
	ws.PhaseEndTime = time.Time{}
	ws.Paused = false
	ws.PausedAt = time.Time{}
	ws.PhasePausedDuration = 0
}

func (ws *WorkoutState) phasePausedDuration() time.Duration {
	total := ws.PhasePausedDuration
	if ws.Paused && !ws.PausedAt.IsZero() {
		total += time.Since(ws.PausedAt)
	}
	if total < 0 {
		return 0
	}
	return total
}

// RemainingTime returns seconds remaining in current phase
func (ws *WorkoutState) RemainingTime() int {
	phaseDuration := time.Duration(ws.CurrentPhase.Duration()) * time.Second
	if phaseDuration <= 0 || ws.PhaseStartTime.IsZero() {
		return 0
	}

	elapsed := time.Since(ws.PhaseStartTime) - ws.phasePausedDuration()
	if elapsed < 0 {
		elapsed = 0
	}

	remaining := phaseDuration - elapsed
	if remaining <= 0 {
		return 0
	}
	return int((remaining + time.Second - 1) / time.Second)
}

// ElapsedTime returns seconds elapsed in current phase
func (ws *WorkoutState) ElapsedTime() int {
	if ws.PhaseStartTime.IsZero() {
		return 0
	}

	elapsed := time.Since(ws.PhaseStartTime) - ws.phasePausedDuration()
	if elapsed < 0 {
		return 0
	}

	return int(elapsed.Seconds())
}

// Progress returns 0.0 to 1.0 representing phase completion
func (ws *WorkoutState) Progress() float64 {
	duration := float64(ws.CurrentPhase.Duration())
	if duration == 0 {
		return 1.0
	}
	elapsed := float64(ws.ElapsedTime())
	progress := elapsed / duration
	if progress > 1.0 {
		return 1.0
	}
	if progress < 0 {
		return 0
	}
	return progress
}
