# Typing Trainer

A CLI typing trainer built with Go and Bubble Tea for practicing code patterns.

## Installation

```bash
go build -o typing-trainer
```

## Usage

Run the trainer:
```bash
./typing-trainer
```

Select a mode and start typing! The trainer tracks your WPM, accuracy, and specific metrics for numbers and symbols.

## Modes

- **Easy/Medium/Hard** - Code snippets of varying complexity
- **Numbers** - Practice typing numbers and numeric expressions
- **Hex Numbers** - Practice hexadecimal notation
- **Symbols** - Focus on special characters and operators
- **Brackets** - Practice matching pairs: (){}[]<>
- **Regex Patterns** - Complex regex patterns
- **Custom** - Your own custom snippets

## Custom Snippets

Create `~/.config/typing-trainer/snippets.json` to add your own practice material:

```json
{
  "snippets": [
    {
      "content": "your code here",
      "language": "Language Name"
    }
  ]
}
```

See `snippets.example.json` for more examples.

## Configuration

The app follows XDG Base Directory specification:
- Config: `~/.config/typing-trainer/` (or `$XDG_CONFIG_HOME/typing-trainer/`)
- History: `~/.config/typing-trainer/history.json`
- Custom snippets: `~/.config/typing-trainer/snippets.json`

Old files from `~/` are automatically migrated on first run.

## Controls

**Menu:**
- Type to search/filter modes
- ↑/↓ or ctrl-n/ctrl-p to navigate
- Enter to select
- ESC to clear search or quit

**Typing:**
- Type the code as shown
- ctrl-n to skip to next snippet
- ESC to return to menu

**Results:**
- Enter to try another snippet
- m to return to menu
- q to quit
