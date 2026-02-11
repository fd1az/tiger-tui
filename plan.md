# tiger-tui: TUI para TigerBeetle

## Contexto

Crear una TUI para operar TigerBeetle, al estilo TablePlus pero en la terminal.
Objetivo: visualizar y operar el ledger financiero completo del sistema Terrace (chart of accounts, transfers, balances, holds).

Proyecto greenfield en `/Users/fd1az/Documents/code/go/tiger-tui/`.

Stack objetivo:
- Go 1.25+
- Bubble Tea (`github.com/charmbracelet/bubbletea`) + Lip Gloss
- Cliente oficial Go de TigerBeetle (`github.com/tigerbeetle/tigerbeetle-go`)
- Config con Viper
- Logging estructurado + errores tipados (`internal/logger`, `internal/apperror`)

Referencias:
- Arquitectura base: `/Users/fd1az/trud/arbitrage-bot`
- Dominio ledger: `/Users/fd1az/terrace/docs/ledger-v2-implementation-plan.md`

Auth:
- TigerBeetle no tiene auth nativa.
- Seguridad a nivel red (VPC/VPN).
- La TUI requiere `cluster_id` + `address` para conectar.

---

## Criterios de diseño (alineado a `design.md`)

Dirección visual: **Mainframe Modern**.
Los wireframes en este plan son funcionales, no cierran el estilo final.

Reglas obligatorias:
- Legibilidad primero: tablas densas, columnas numéricas alineadas a la derecha.
- Semántica de color estricta: `accent` para foco, `success|warning|error|info` solo para estado.
- Una sola zona activa por vez.
- Decoración mínima.
- Compatibilidad terminal: Unicode opcional con fallback ASCII.

Tokens base (dark):
- `bg`: `#0B0F14`
- `panelBg`: `#101826`
- `border`: `#243041`
- `text`: `#E6EDF3`
- `textMuted`: `#9BA7B4`
- `textDim`: `#6B7785`
- `accent`: `#F4B942`
- `success`: `#2ECC71`
- `warning`: `#F39C12`
- `error`: `#E74C3C`
- `info`: `#22C1C3`

Nota:
- Se mantienen los layouts de pantallas ya definidos (Connection, Dashboard Tabs, Details y overlays).
- `design.md` es la referencia visual de detalle.

---

## Arquitectura objetivo

Patrones combinados:
1. **Hexagonal + módulos por dominio** (alineado con `arbitrage-bot`).
2. **The Elm Architecture (TEA)** para la UI Bubble Tea: `Model -> Msg -> Update -> View`.

Decisiones de arquitectura:
- `cmd/` como entrypoint.
- `business/` para bounded contexts de producto (`connection`, `accounts`, `transfers`, `balancesheet`).
- `internal/` para servicios transversales (`config`, `logger`, `apperror`, `di`, `cache`, `circuitbreaker`, etc.).
- `pkg/ui` para TUI y componentes reutilizables.
- Contenedor tipo monolith para registrar/inicializar módulos, inspirado en `internal/monolith` de `arbitrage-bot`.

### Flujo de mensajes (Bubble Tea + workers)

- `Update` maneja teclado/UI state.
- Operaciones IO (TigerBeetle query/create) salen por `tea.Cmd` o goroutines.
- Resultados vuelven como `tea.Msg` (`ConnectedMsg`, `AccountsLoadedMsg`, `TransferCreatedMsg`, `ErrorMsg`).
- Cuando el worker esté fuera del loop principal, enviar eventos con `program.Send(msg)`.

Estados de conexión:
- `Disconnected`
- `Connecting`
- `Connected`
- `Degraded`

---

## Estructura propuesta

```text
tiger-tui/
├── cmd/
│   └── tiger-tui/
│       └── main.go
│
├── business/
│   ├── connection/
│   │   ├── module.go
│   │   ├── app/
│   │   │   └── service.go
│   │   └── infra/
│   │       └── tigerbeetle_client.go
│   ├── accounts/
│   │   ├── module.go
│   │   ├── domain/
│   │   │   └── types.go
│   │   ├── app/
│   │   │   └── service.go
│   │   └── infra/
│   │       └── repository_tb.go
│   ├── transfers/
│   │   ├── module.go
│   │   ├── domain/
│   │   │   └── types.go
│   │   ├── app/
│   │   │   └── service.go
│   │   └── infra/
│   │       └── repository_tb.go
│   └── balancesheet/
│       ├── module.go
│       └── app/
│           └── service.go
│
├── internal/
│   ├── apperror/
│   ├── cache/
│   ├── circuitbreaker/
│   ├── config/
│   ├── di/
│   ├── logger/
│   └── monolith/
│       └── monolith.go
│
├── pkg/
│   └── ui/
│       ├── tui.go
│       ├── messages.go
│       ├── keys.go
│       ├── styles.go
│       ├── models.go
│       └── components/
│           ├── connection_form.go
│           ├── accounts_table.go
│           ├── transfers_table.go
│           ├── balance_sheet.go
│           └── status_bar.go
│
├── design.md
├── plan.md
└── README.md
```

---

## Dependencias Go

`go.mod` base ya incluye Bubble Tea y utilidades internas.
Agregar/confirmar:

```bash
go get github.com/tigerbeetle/tigerbeetle-go
go get github.com/charmbracelet/lipgloss
```

Dependencias clave:
- `github.com/charmbracelet/bubbletea`
- `github.com/charmbracelet/lipgloss`
- `github.com/tigerbeetle/tigerbeetle-go`
- `github.com/spf13/viper`

---

## Modelo de mensajes UI (MVP)

```go
type ConnectedMsg struct{}
type ConnectionFailedMsg struct{ Err error }
type AccountsLoadedMsg struct{ Rows []AccountRow }
type TransfersLoadedMsg struct{ Rows []TransferRow }
type BalanceSheetLoadedMsg struct{ Summary BalanceSheetSummary }
type CreateTransferResultMsg struct{ Results []CreateTransferResult }
type ErrorMsg struct{ Err error }
```

Comandos principales:
- `ConnectCmd(clusterID uint32, address string)`
- `LoadAccountsCmd(filter AccountFilter)`
- `LoadTransfersCmd(filter TransferFilter)`
- `LoadBalanceSheetCmd()`
- `CreateTransferCmd(req CreateTransferRequest)`

---

## Chart of accounts (ledger-v2)

La UI debe traducir IDs/códigos de TigerBeetle a etiquetas legibles:
- `ledger_id -> asset`
- `account_code -> account_type`
- `transfer_code -> transfer_type`
- `user_data_32 -> venue`

Ubicación propuesta:
- `business/accounts/domain` para tipos/mappings de cuenta.
- `business/transfers/domain` para mappings de transfer.
- Helpers compartidos en `internal` si evita duplicación.

---

## Estrategia de errores operativos

Principios:
- Errores tipados en `internal/apperror`.
- Contexto y trazabilidad en logs (`operation`, `cluster_id`, `address`, `latency_ms`).
- Nunca romper el loop de UI por un error de IO.

Política:
- Lecturas (`query`, `lookup`): mantener último snapshot válido y mostrar error en status bar.
- Escrituras (`create`): feedback por item (`ok/error_code`) sin cerrar modal si hay errores.
- Conexión: retries con backoff acotado (250ms, 500ms, 1s, 2s, 5s).

---

## Fases de implementación

### Fase 1: Base Go + Bubble Tea + Connection
1. Consolidar entrypoint en `cmd/tiger-tui/main.go`.
2. Crear `pkg/ui` con `Model`, `Update`, `View`, keybindings y status bar.
3. Implementar pantalla Connection (`cluster_id`, `address`, `Connect`).
4. Integrar `internal/config`, `internal/logger`, `internal/apperror`.

DoD:
- `go run ./cmd/tiger-tui` abre TUI.
- Inputs editables y submit funcional.
- `q` y `Ctrl+C` cierran limpio.

### Fase 2: Integración TigerBeetle real
1. Agregar adapter `tigerbeetle-go` para connect/query.
2. Implementar pipeline de mensajes async en Bubble Tea.
3. Manejar `Connected/ConnectionFailed/Timeout` en status bar.

DoD:
- Conecta a TigerBeetle local (`cluster=0`, `address=3000`).
- Address inválida y server caído muestran error sin crash.

### Fase 3: Dashboard (Accounts / Transfers / Balance Sheet)
1. Tabla de Accounts con filtros (type/ledger/code).
2. Tabla de Transfers con filtros (type/ledger/pending).
3. Balance Sheet agrupado por tipo de cuenta + validación trial balance.
4. Navegación por tabs, selección y detalle.

DoD:
- Datos reales visibles con labels legibles.
- Trial balance reporta `balanced`/`unbalanced` claramente.

### Fase 4 (post-MVP): Formularios Create
1. Overlay Create Account.
2. Overlay Create Transfer (flags + user_data).
3. Validaciones y feedback por item.

### Fase 5 (post-MVP): Lookup y refinamientos
1. Lookup por ID (`/`).
2. Help overlay (`?`).
3. Auto-refresh configurable.
4. Sparkline de balance en account detail.

---

## Testing y aceptación

Suite mínima:
- Unit tests:
  - mapping chart of accounts
  - formatters (amount/timestamps)
  - reducers `Update()` por `tea.Msg`
- Integration tests (con TigerBeetle local):
  - `Connect -> QueryAccounts -> QueryTransfers`
- UI smoke:
  - render pantallas principales sin panic
  - navegación básica (`Tab`, `j/k`, `Esc`)

Comandos sugeridos:

```bash
go fmt ./...
go vet ./...
go test ./...
make check
```

---

## Verificación end-to-end

```bash
# 1. Crear data file TB
tigerbeetle format --cluster=0 --replica=0 --replica-count=1 0_0.tigerbeetle

# 2. Iniciar TigerBeetle
tigerbeetle start --addresses=3000 0_0.tigerbeetle

# 3. Ejecutar TUI
cd /Users/fd1az/Documents/code/go/tiger-tui
go run ./cmd/tiger-tui

# 4. Conectar con cluster_id=0 y address=3000
```

---

## Checklist de migración completada (Rust -> Go)

- [x] Contexto y path del proyecto actualizados a Go.
- [x] Stack actualizado a Bubble Tea + TigerBeetle Go client.
- [x] Arquitectura alineada a referencia `arbitrage-bot`.
- [x] Estructura de carpetas reescrita para `cmd/business/internal/pkg`.
- [x] Fases de implementación actualizadas a comandos Go/Make.
- [x] QA/testing actualizado (`go test`, `make check`).
