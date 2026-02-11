# tiger-tui

A terminal UI for [TigerBeetle](https://tigerbeetle.com): TablePlus-style ledger operations directly in the terminal.

## Estado actual

Proyecto migrado a Go + Bubble Tea.

Implementado hoy:
- Pantalla de conexión (`cluster_id`, `address`).
- Dashboard shell con tabs: Accounts, Transfers, Balance Sheet.
- Tema visual y status bar.
- Logging a archivo (`tiger-tui.log`).

Pendiente:
- Conexión real a TigerBeetle (hoy está simulada en la UI).
- Carga de accounts/transfers/balance sheet reales.

## Stack

- Go 1.25+
- Bubble Tea + Bubbles + Lip Gloss
- Viper
- Paquetes internos para logger, DI, cache, circuit breaker y errores tipados

## Requisitos

- Go 1.25 o superior
- (Opcional por ahora) TigerBeetle local para pruebas de integración futuras

## Quick start

```bash
git clone https://github.com/fd1az/tiger-tui.git
cd tiger-tui

# Ejecutar
go run ./cmd/tiger-tui

# o con Makefile
make run
```

## Keybindings

| Key | Acción |
|---|---|
| `Tab` / `Shift+Tab` | Navegar campos / cambiar tab |
| `Enter` | Submit / seleccionar |
| `Esc` | Volver a Connection desde Dashboard |
| `q` | Salir |
| `Ctrl+C` | Salir forzado |

## Estructura

```text
cmd/tiger-tui/main.go         # Entry point
pkg/ui/                       # Bubble Tea model, mensajes y componentes TUI
internal/config/              # Configuración
internal/logger/              # Logging estructurado
internal/apperror/            # Errores de aplicación
internal/di/                  # Contenedor DI
internal/cache/               # Cache utilitaria
internal/circuitbreaker/      # Circuit breaker
business/accounts/domain/     # Mapping y dominio de cuentas
```

## Desarrollo

```bash
make fmt
make vet
make test
make check
```

## Roadmap

- [x] Fase 1: scaffold TUI + pantalla de conexión + dashboard shell
- [ ] Fase 2: integración real con TigerBeetle Go client
- [ ] Fase 3: tablas de Accounts/Transfers + Balance Sheet
- [ ] Fase 4: formularios Create Account/Create Transfer
- [ ] Fase 5: lookup, auto-refresh y mejoras de detalle

## License

MIT
