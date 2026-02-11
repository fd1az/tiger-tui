# TigerBeetle TUI Theme & UI Guidelines

A TablePlus-like TUI for TigerBeetle should feel like a **ledger / operator console**: precise, calm, high-contrast, with **minimal accent color** and strong **semantic color meaning**.

---

## Design Goals (Normative)

- **Legibility first** (dense tables, fast scanning).
- **Few colors, strong semantics** (status colors mean something).
- **Operator vibe** (correctness + performance, not “gamer UI”).
- **Consistent focus model** (only one “active” area at a time).

### Non-Negotiables
- Keep decorative styling secondary to data scanning.
- Use status colors only for state semantics, never for arbitrary emphasis.
- Maintain ASCII fallback for symbols/icons in constrained terminals.
- Prioritize stable layouts over visual flourish.

---

## Recommended Theme: “Mainframe Modern”

### Color Tokens (Dark)

**Base**
- `bg`: `#0B0F14` (near-black)
- `panelBg`: `#101826`
- `border`: `#243041`
- `text`: `#E6EDF3`
- `textMuted`: `#9BA7B4`
- `textDim`: `#6B7785`

**Accent (aesthetic)**
- `accent`: `#F4B942` (TigerBeetle amber / gold)
  - Use for: focus ring, selected row, primary highlight

**Semantic (status)**
- `success`: `#2ECC71` (committed / posted / ok)
- `warning`: `#F39C12` (pending / unusual / needs attention)
- `error`: `#E74C3C` (rejected / invariant broken / failure)
- `info`: `#22C1C3` (replication / networking / state transitions)

> Rule: **1 aesthetic accent + up to 3 semantic colors** shown at once.

### Optional Light Theme

**Base**
- `bg`: `#F7F7F8`
- `panelBg`: `#FFFFFF`
- `border`: `#E5E7EB`
- `text`: `#111827`
- `textMuted`: `#4B5563`
- `textDim`: `#6B7280`

**Accent**
- `accent`: `#B45309` (darker amber for contrast)

**Semantic**
- `success`: `#15803D`
- `warning`: `#B45309`
- `error`: `#B91C1C`
- `info`: `#0E7490`

---

## UI Layout (TablePlus-like)

### 3-Zone Layout (MVP)

1. **Top Bar**
   - Connection status (cluster/env)
   - latency / mode (read-only vs write)
   - keybind hints (minimal)

2. **Primary Navigation (Top Tabs)**
   - `Accounts`
   - `Transfers`
   - `Balance Sheet`

3. **Main Area**
   - Table (primary)
   - Details panel (split view, toggleable)

`Left Navigation` for `Events/Queries/Metrics/Logs` is optional post-MVP.

---

## Interaction Model (Keybinds)

- `/` incremental search
- `f` filter builder (chips)
- `s` sort
- `Space` toggle details / preview row
- `Enter` drill-down
- `e` export (CSV / JSON)
- `:` command palette (vim-like command bar)
- `Esc` back / close modal

---

## Table Visual Rules

- Dense rows, consistent column widths, right-align numbers.
- Truncate with ellipsis in table (`…`), full view in Details.
- Use subtle borders/dividers (avoid heavy ASCII boxes).
- Only highlight:
  - selected row (accent background or accent left bar)
  - focused pane (accent border)
- Keep “decorative” color usage minimal.

---

## Semantic Status Markers (Ledger-Friendly)

Use **minimal symbols** + color (ASCII fallback shown first):

- `OK` (or `✓`) committed / durable
- `~` pending / in-flight
- `!` warning / unusual
- `ERR` (or `✗`) error / rejected

Example row prefix:
- `OK transfer posted`
- `~ transfer pending`
- `ERR invariant violation`

---

## Typography / Readability

- Prefer terminal default monospace.
- Avoid fancy unicode unless optional; keep ASCII fallback.
- Ensure contrast:
  - Primary text always readable on `bg` and `panelBg`
  - Dim text used only for hints, not for critical data

---

## Component Styling Guidelines

**Top Bar**
- `panelBg` background
- `textMuted` for hints
- `accent` only for active connection indicator

**Nav**
- Active item: `accent` + bold
- Inactive: `textMuted`
- Badges (counts): `textDim` background, `text` label

**Table**
- Header: `text` with `border` underline
- Row hover/selection: accent
- Numeric columns: right-aligned, `text`

**Details Panel**
- Key labels: `textMuted`
- Values: `text`
- Errors: `error` + short message + code

---

## Color Discipline Checklist

- [ ] Is color conveying meaning (status) or only decoration?
- [ ] Only one “focused” area at a time?
- [ ] Selected row uses `accent`, not random colors?
- [ ] Errors are red **only** when truly error?
- [ ] Warnings aren’t used for normal states?

---

## Style Direction Options

1) **Mainframe Modern (recommended)**
- Minimal, infra/ledger vibe, strong semantics.

2) **Aurora Devtool**
- More “product-y” feel (only if terminal renderer supports subtle effects).
- Still: keep semantic colors strict.

---

## Notes

TigerBeetle = correctness + performance. The UI should feel:
- deterministic
- calm under pressure
- built for scanning, filtering, exporting, drilling-down fast

If you choose a UI framework (Go Bubble Tea / Rust ratatui / etc.),
define these tokens in a single file and map components to tokens consistently.
