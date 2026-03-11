# code-hiit

code-hiit is a CLI typing trainer built with Go and Bubble Tea for practicing code patterns.

## Installation

```bash
go build -o code-hiit
```

## Usage

Run the trainer:
```bash
./code-hiit
```

Select a mode and start typing! The trainer tracks your WPM, accuracy, and specific metrics for numbers and symbols.

## Modes

- **Easy Code / Medium Code / Hard Code** - Code snippets of varying complexity
- **Numbers Practice** - Practice typing numbers and numeric expressions
- **Symbols Practice** - Focus on special characters and operators
- **Hex Numbers** - Practice hexadecimal notation
- **Brackets Practice** - Practice matching pairs: (){}[]<>
- **Regex Patterns** - Complex regex patterns
- **Custom** - Your own custom snippets

## Custom Snippets

Create `~/.config/code-hiit/snippets.json` to add your own practice material:

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
- Config: `~/.config/code-hiit/` (or `$XDG_CONFIG_HOME/code-hiit/`)
- History: `~/.config/code-hiit/history.json`
- Custom snippets: `~/.config/code-hiit/snippets.json`

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

## License

MIT License — see `LICENSE` for details.
