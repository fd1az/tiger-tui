// Package ui provides the Bubble Tea TUI for tiger-tui.
package ui

import "github.com/charmbracelet/lipgloss"

// Mainframe Modern — color tokens (dark theme)
var (
	ColorBg      = lipgloss.Color("#0B0F14")
	ColorPanelBg = lipgloss.Color("#101826")
	ColorBorder  = lipgloss.Color("#243041")
	ColorText    = lipgloss.Color("#E6EDF3")
	ColorMuted   = lipgloss.Color("#9BA7B4")
	ColorDim     = lipgloss.Color("#6B7785")

	// Accent (aesthetic)
	ColorAccent = lipgloss.Color("#F4B942") // TigerBeetle amber

	// Semantic (status)
	ColorSuccess = lipgloss.Color("#2ECC71")
	ColorWarning = lipgloss.Color("#F39C12")
	ColorError   = lipgloss.Color("#E74C3C")
	ColorInfo    = lipgloss.Color("#22C1C3")
)

// Reusable styles
var (
	// Box / panel
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(0, 1)

	BoxFocusedStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorAccent).
			Padding(0, 1)

	// Title
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccent)

	// Text
	TextStyle  = lipgloss.NewStyle().Foreground(ColorText)
	MutedStyle = lipgloss.NewStyle().Foreground(ColorMuted)
	DimStyle   = lipgloss.NewStyle().Foreground(ColorDim)

	// Accent
	AccentStyle     = lipgloss.NewStyle().Foreground(ColorAccent)
	AccentBoldStyle = lipgloss.NewStyle().Foreground(ColorAccent).Bold(true)

	// Semantic
	SuccessStyle = lipgloss.NewStyle().Foreground(ColorSuccess)
	WarningStyle = lipgloss.NewStyle().Foreground(ColorWarning)
	ErrorStyle   = lipgloss.NewStyle().Foreground(ColorError)
	InfoStyle    = lipgloss.NewStyle().Foreground(ColorInfo)

	// Status indicators
	StatusConnected    = lipgloss.NewStyle().Foreground(ColorSuccess).Bold(true)
	StatusDisconnected = lipgloss.NewStyle().Foreground(ColorDim)
	StatusConnecting   = lipgloss.NewStyle().Foreground(ColorWarning).Bold(true)

	// Tabs
	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccent).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(ColorAccent).
			Padding(0, 2)

	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(ColorMuted).
				Border(lipgloss.NormalBorder(), false, false, true, false).
				BorderForeground(ColorDim).
				Padding(0, 2)

	// Table
	TableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorText).
				BorderBottom(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(ColorBorder)

	TableCellStyle = lipgloss.NewStyle().Padding(0, 1)

	// Help bar
	HelpStyle = lipgloss.NewStyle().Foreground(ColorDim)

	// Input field
	InputLabelStyle = lipgloss.NewStyle().Foreground(ColorMuted).Width(14)

	// Logo
	LogoStyle = lipgloss.NewStyle().Foreground(ColorAccent).Bold(true)
)
