# tiger-tui

A terminal UI for [TigerBeetle](https://tigerbeetle.com): TablePlus-style ledger operations directly in the terminal.

## Current status

Built with Go + Bubble Tea.

Implemented:
- Connection screen (`cluster_id`, `address`)
- Dashboard shell with tabs: Accounts, Transfers, Balance Sheet
- Mainframe Modern theme and status bar
- File logging (`tiger-tui.log`)

Pending:
- Real TigerBeetle connection (currently simulated in the UI)
- Load actual accounts/transfers/balance sheet data

## Stack

- Go 1.25+
- Bubble Tea + Bubbles + Lip Gloss
- Viper
- Internal packages for logger, DI, cache, circuit breaker, and typed errors

## Requirements

- Go 1.25 or higher
- (Optional for now) Local TigerBeetle instance for future integration testing

## Quick start

```bash
git clone https://github.com/fd1az/tiger-tui.git
cd tiger-tui

# Run directly
go run ./cmd/tiger-tui

# Or with Makefile
make run
```

## Keybindings

| Key | Action |
|---|---|
| `Tab` / `Shift+Tab` | Navigate fields / cycle tabs |
| `Enter` | Submit / select |
| `Esc` | Return to Connection from Dashboard |
| `q` | Quit |
| `Ctrl+C` | Force quit |

## Structure

```text
cmd/tiger-tui/main.go         # Entry point
pkg/ui/                       # Bubble Tea model, messages, and TUI components
internal/config/              # Configuration
internal/logger/              # Structured logging
internal/apperror/            # Application errors
internal/di/                  # DI container
internal/cache/               # Utility cache
internal/circuitbreaker/      # Circuit breaker
business/accounts/domain/     # Account mapping and domain
```

## Development

```bash
make fmt
make vet
make test
make check
```

## Roadmap

- [x] Phase 1: TUI scaffold + connection screen + dashboard shell
- [ ] Phase 2: Real TigerBeetle Go client integration
- [ ] Phase 3: Accounts/Transfers tables + Balance Sheet
- [ ] Phase 4: Create Account/Create Transfer forms
- [ ] Phase 5: Lookup, auto-refresh, and detail improvements

## License

MIT
