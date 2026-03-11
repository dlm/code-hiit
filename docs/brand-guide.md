# Code HIIT Lab – Brand Guide (Web)

Visual direction for the Code HIIT Lab site (GitHub Pages). Theme: HIIT gym energy meets terminal clarity. Tone: confident, kinetic, concise.

## Positioning
- What: High-intensity interval typing for coders; CLI is `code-hiit`.
- Feel: Interval timer + command line; sweat-but-smart, experimental lab.
- Voice: Short, imperative, motivational (“Run a set”, “Hit start”, “Recovery”).

## Palette
- **Primary Surge** `#00E8A9` — CTA fill, links, highlight bars.
- **Heat Accent** `#FF5E5B` — warnings, active intervals, hover rings.
- **Dark Base** `#0C111B` — hero background, footer.
- **Graphite** `#141927` — cards, panels.
- **Slate** `#1F2535` — nav, secondary panels.
- **Light Text** `#E9EDF5` — body text on dark.
- **Muted Text** `#A8B3C6` — secondary copy.
- **Border** `#2C3346` — strokes/dividers.
- **Success/Recovery** `#39C07F` — “recovered” states, success toasts.
- Gradient suggestion: `linear-gradient(135deg, #0C111B 0%, #11182A 40%, #14253D 100%)` with thin neon accents (1–2px) in `#00E8A9`.

Contrast pairs that clear WCAG AA on large text: Primary Surge on Dark Base, Light Text on Graphite/Slate, Heat Accent on Dark Base.

## Typography
- **Headlines**: `Space Grotesk` (600/700). Wide, modern; set letter-spacing slightly negative (-0.01em).
- **Body**: `Manrope` (400/500). Clean, tech-friendly without feeling default.
- **Mono**: `JetBrains Mono` (500) for code callouts, stat readouts, and buttons with key labels.
- CSS import example:
  ```css
  @import url('https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@500;600;700&family=Manrope:wght@400;500;600&family=JetBrains+Mono:wght@500&display=swap');

  :root {
    --font-display: 'Space Grotesk', system-ui, -apple-system, sans-serif;
    --font-body: 'Manrope', system-ui, -apple-system, sans-serif;
    --font-mono: 'JetBrains Mono', SFMono-Regular, Consolas, monospace;
  }
  ```

## Logo & Icon
- **Wordmark**: “Code HIIT Lab” in Space Grotesk, uppercase or small caps. Emphasize “HIIT” with Heat Accent underline or block.
- **Mark/Favicon**: Monogram `CH` inside a rounded square “plate” with a subtle inner stroke; Primary Surge fill, Heat Accent stroke. Alternate: stylized interval timer bars (3 vertical ticks) inside the plate.
- **CLI Lockup**: `code-hiit` set in JetBrains Mono with a thin accent bar under “HIIT”.

## UI Elements
- **Buttons**: Primary fill `#00E8A9`, text Dark Base, radius 8px, bold mono text. Hover: lift + 2px outline in Heat Accent; Active: slight inset shadow.
- **Secondary Button**: Outline in Border color, text Light Text; hover fills Slate with Primary outline.
- **Pills/Tags**: Background Slate, text Muted Text; active tag uses Heat Accent text and Primary outline.
- **Cards**: Graphite background, Border stroke, 12–16px radius, top accent bar 3px in Primary Surge for key sections (Features, Modes).
- **CTA Bar**: Horizontal band with gradient background and split layout for copy + button; include a small interval timer icon.
- **Dividers**: 1px Border; occasional neon hairline (1px Primary) for section separators.

## Layout & Sections (for Pages)
- **Hero**: Dark Base gradient, left-aligned headline (“Code HIIT Lab”), subline (“High-intensity intervals for coders”), CTA buttons (`Run code-hiit`, `View on GitHub`). Right side: terminal mockup or animated GIF of the CLI. Add small “interval chips” floating with Primary and Heat dots.
- **Key Stats/Features**: 3–4 cards with icon circles (Primary/Heat) and short copy (Intervals, Symbols/Numbers focus, History/Stats).
- **Modes**: Grid of cards/pills listing modes (Easy/Medium/Hard code, Numbers, Symbols, Hex, Brackets, Regex, Custom).
- **How to Run**: Step list with mono code blocks; include the new name `code-hiit`.
- **Brand Bar**: A slim bar showing palette chips and font stack as a quick brand cue.
- **Footer**: Dark Base, muted links, GitHub icon.

## Motion
- Keep motion meaningful: staggered reveal on hero text (60–90ms), glow pulse on Primary buttons every ~8s, slide-in for interval chips. Avoid overshooting easings; prefer `cubic-bezier(0.16, 1, 0.3, 1)`.

## Imagery
- Use CLI captures on dark background with slight blur shadow. Overlay a thin neon grid or subtle noise texture on hero background (2–4% opacity) to avoid flatness.

## Tone & Copy Starters
- Headline: “Intervals for coders.” / “Short bursts, sharper keystrokes.”
- CTA: “Run a set” / “Start an interval” / “Open code-hiit”.
- Section labels: “Warmup”, “Work Interval”, “Recovery”, “Progress”.

## Assets
- Mark: `docs/assets/code-hiit-mark.svg` (rounded plate + interval bars; works for hero/lockups).
- Favicon/icon: `docs/assets/code-hiit-favicon.svg` (simplified for 16–64px).
- Usage: set `rel="icon"` to the favicon; use the mark inline in hero/CTA bars; keep on dark backgrounds or add a 2px Primary outline on light.

## Deliverables for Build Task
- Apply palette, typography imports, and component styles above.
- Use provided SVG assets for mark/favicon (rounded square + interval bars).
- Use the layout outline for the Pages site sections.
