# Code HIIT Lab

**High-Intensity Interval Training for developers.**

🌐 **[Visit the project site](https://dlm.github.io/code-hiit)** for screenshots, features, and installation guide.

---

Code HIIT Lab is a terminal-based typing trainer that uses HIIT methodology to build muscle memory for code patterns. Short, intense work sets with recovery periods help you master symbols, brackets, numbers, and real code snippets — the characters that trip up coders most.

## Why HIIT?

Traditional typing trainers focus on sustained speed. Code HIIT Lab uses **High-Intensity Interval Training** — alternating between focused work phases and recovery periods — to improve both speed and accuracy under pressure. Just like HIIT workouts build athletic performance, Code HIIT Lab builds coding performance.

## Installation

### Quick Install (Recommended)

**macOS and Linux:**
```bash
curl -fsSL https://raw.githubusercontent.com/dlm/code-hiit/main/install.sh | sh
```

### Homebrew

```bash
brew install dlm/tap/code-hiit
```

### Manual Installation

Download pre-built binaries from the [releases page](https://github.com/dlm/code-hiit/releases/latest):

1. Download the appropriate binary for your system:
   - **Linux AMD64**: `code-hiit-linux-amd64`
   - **macOS Apple Silicon**: `code-hiit-darwin-arm64`

2. Make it executable and move to your PATH:
   ```bash
   chmod +x code-hiit-*
   sudo mv code-hiit-* /usr/local/bin/code-hiit
   ```

**Supported Platforms:** Linux AMD64 (x86_64), macOS ARM64 (Apple Silicon)

### Build from Source

Requires Go 1.23 or later:

```bash
git clone https://github.com/dlm/code-hiit.git
cd code-hiit
make local
# or: go build -o code-hiit
```

## Quick Start

```bash
code-hiit
```

Pick a mode, choose your workout length, and start typing. The timer runs fixed 15s/30s/10s intervals (warmup/work/recovery) while tracking your WPM and accuracy.

## Workout Modes

- **Easy Code / Medium Code / Hard Code** - Real code snippets at varying difficulty
- **Numbers** - Numeric expressions and data patterns
- **Symbols** - Operators and special characters
- **Hex** - Hexadecimal notation practice
- **Brackets** - Matching pairs: `(){}[]<>`
- **Regex** - Complex pattern matching expressions
- **Custom** - Your own snippets from your codebase

## HIIT Phases

Each workout includes:
- **Warmup** (15s) - Get your fingers ready
- **Work** (30s) - Maximum intensity typing
- **Recovery** (10s) - Brief rest between sets
- **Summary** - Review your performance stats

## Controls

**Menu:**
- Type to search/filter modes
- ↑/↓ or `ctrl-n`/`ctrl-p` to navigate
- Enter to select
- ESC to clear search or quit

**During workout:**
- Type the code as shown
- `ctrl-space` to pause/resume
- `ctrl-n` to skip to next snippet
- ESC to return to menu

**Results:**
- Enter to try another snippet
- `m` to return to menu
- `q` to quit

## Custom Snippets

Add your own practice material from your codebase:

Create `~/.config/code-hiit/snippets.json`:

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

See `snippets.example.json` for examples.

## Configuration

Code HIIT Lab follows XDG Base Directory specification:
- **Config:** `~/.config/code-hiit/`
- **History:** `~/.config/code-hiit/history.json`
- **Custom snippets:** `~/.config/code-hiit/snippets.json`

Old files from `~/` are automatically migrated on first run.

## Stats & Progress

Your workout history is automatically saved to `~/.config/code-hiit/history.json`. Track your progress over time and see how your speed and accuracy improve with each session.

## License

MIT License — see `LICENSE` for details.
