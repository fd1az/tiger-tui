# tiger-tui

A terminal UI for [TigerBeetle](https://tigerbeetle.com) — like TablePlus, but for your ledger.

```
 ████████╗██╗ ██████╗ ███████╗██████╗       ████████╗██╗   ██╗██╗
 ╚══██╔══╝██║██╔════╝ ██╔════╝██╔══██╗      ╚══██╔══╝██║   ██║██║
    ██║   ██║██║  ███╗█████╗  ██████╔╝  █████╗ ██║   ██║   ██║██║
    ██║   ██║██║   ██║██╔══╝  ██╔══██╗  ╚════╝ ██║   ██║   ██║██║
    ██║   ██║╚██████╔╝███████╗██║  ██║         ██║   ╚██████╔╝██║
    ╚═╝   ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝         ╚═╝    ╚═════╝ ╚═╝
```

Browse accounts, inspect transfers, and verify balance sheets — all from
your terminal with a keyboard-driven workflow designed for operators.

## Features

- **Connection screen** — connect to any TigerBeetle cluster by ID and address
- **Dashboard with tabs** — Accounts, Transfers, Balance Sheet
- **Chart of accounts** — human-readable labels for ledger IDs, account codes, transfer codes, and venues (Terrace ledger-v2 compatible)
- **Mainframe Modern theme** — high-contrast dark palette with semantic colors (amber accent, green/yellow/red status)
- **Keyboard-first** — Tab/Shift+Tab navigation, vim-style keybinds (planned)

## Requirements

- Rust 2024 edition (1.85+)
- A running [TigerBeetle](https://docs.tigerbeetle.com/) instance (for full functionality)

## Quick start

```bash
# Clone and build
git clone https://github.com/fd1az/tiger-tui.git
cd tiger-tui
cargo build --release

# Start a local TigerBeetle (if you don't have one running)
tigerbeetle format --cluster=0 --replica=0 --replica-count=1 0_0.tigerbeetle
tigerbeetle start --addresses=3000 0_0.tigerbeetle

# Run
cargo run --release
```

Connect with cluster ID `0` and address `3000`, then press **Connect**.

## Keybindings

| Key | Action |
|---|---|
| `Tab` / `Shift+Tab` | Navigate fields / cycle tabs |
| `Enter` | Submit / select |
| `q` / `Esc` | Quit / back |
| `Ctrl+C` | Force quit |

## Architecture

Two patterns combined:

- **The Elm Architecture (TEA)** for the UI loop — `AppState` (Model) → `Message` → `update()` → `view()`
- **Hexagonal Architecture** for modules — domain, ports, infra, and UI layers per feature

```
src/
├── main.rs                     # Terminal setup, TEA main loop
├── shared/
│   ├── app_state.rs            # Global model (Screen, Tab, ConnectionForm)
│   ├── message.rs              # Message enum + update() reducer
│   ├── event.rs                # Crossterm event polling (mpsc channel)
│   ├── error.rs                # AppError (thiserror)
│   └── domain/
│       └── chart_of_accounts.rs  # Ledger/account/transfer/venue mappings
├── modules/
│   └── connection/
│       └── ui/
│           └── connection_view.rs  # Connection screen renderer
└── ui/
    ├── mod.rs                  # View dispatcher
    ├── theme.rs                # Mainframe Modern color tokens
    ├── dashboard.rs            # Dashboard layout (tabs + content)
    └── status_bar.rs           # Status bar (connection info + hints)
```

## Development

```bash
# Format, lint, test
cargo fmt --all
cargo clippy --all-targets --all-features -- -D warnings
cargo test
```

Clippy runs with pedantic lints enabled. See `[lints]` in `Cargo.toml` for the full configuration.

## Roadmap

- [x] Phase 1 — Scaffold, connection screen, dashboard shell
- [ ] Phase 2 — TigerBeetle client integration (async bridge via tokio + mpsc)
- [ ] Phase 3 — Accounts table, transfers table, balance sheet with trial balance
- [ ] Phase 4 — Create account/transfer forms
- [ ] Phase 5 — Lookup by ID, auto-refresh, balance sparklines

## License

MIT
