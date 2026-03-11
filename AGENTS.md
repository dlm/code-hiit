## MANDATORY: Use td for Task Management

You must run `td usage --new-session` at conversation start (or after /clear) to see current work.
Use `td usage -q` for subsequent reads.

## Project Overview

This is a typing trainer CLI application built in Go using Charm's Bubble Tea framework.
The goal is to help users improve typing skills for numbers, symbols, and code patterns.

## Current State

- Working typing trainer with 5 modes: Easy, Medium, Hard, Numbers, Symbols
- Built with Go, Bubble Tea, and Lipgloss
- Uses Nix flakes for development environment
- Session history saved to `~/.typing-history.json`

## Workflow for New Agents

1. **Check for active work:**
   ```bash
   td status              # See current session and focus
   td resume              # View handoff from previous agent
   ```

2. **Pick up or start a task:**
   ```bash
   td ready               # See highest priority open tasks
   td start <task-id>     # Begin work on a task
   ```

3. **During development:**
   - Build: `go build -o code-hiit`
   - Test: `./code-hiit`
   - Main files:
     - `types.go` - Data structures and difficulty levels
     - `snippets.go` - Code snippet library
     - `main.go` - Bubble Tea UI and logic
     - `history.go` - Session persistence

4. **When handing off or completing work:**

   **CRITICAL**: ALWAYS do a proper handoff BEFORE submitting for review!

   ```bash
   # Step 1: Create detailed handoff (REQUIRED)
   td handoff <task-id> \
     --done "Detailed list of what you completed" \
     --remaining "What's left to do (or 'None' if complete)" \
     --uncertain "Any blockers or questions (or 'None')"

   # Step 2: ONLY AFTER handoff, submit for review
   td review <task-id>
   ```

   **BAD PRACTICE**: Never run `td review` without doing `td handoff` first!
   The auto-generated handoff is minimal and doesn't document your work properly.

   **GOOD HANDOFF EXAMPLE**:
   ```bash
   td handoff td-abc123 \
     --done "Created custom.go with LoadCustomSnippets() function. \
             Added Custom difficulty to types.go. Updated main.go menu. \
             Modified GetRandomSnippet, GetSnippet, GetNextSnippet functions. \
             Created snippets.example.json under XDG config. Tested functionality." \
     --remaining "None - feature complete and ready for review" \
     --uncertain "None - all functionality tested and working"
   ```

5. **If task is complete:**
   ```bash
   td approve <task-id>   # Approve and close
   ```

## Task Structure

All features are organized under epic `td-ce6be4`:
- More Training Modes (P1-P3)
- Statistics & Progress Tracking (P2-P3)
- Practice Features (P2-P3)
- Quality of Life Improvements (P3-P4)
- Advanced Statistics (P3-P4)
- Social & Competition Features (P4)

Use `td tree td-ce6be4` to see the full roadmap.

## Key Commands Reference

```bash
td list --type feature --priority P2    # Filter by priority
td show <task-id>                       # View task details
td next                                 # Show next highest priority task
td tree <epic-id>                       # View task hierarchy
td info                                 # Project overview
```
