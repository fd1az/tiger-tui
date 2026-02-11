// Package components provides reusable TUI components for tiger-tui.
package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Colors (imported from parent but kept local to avoid circular imports)
var (
	colorAccent  = lipgloss.Color("#F4B942")
	colorText    = lipgloss.Color("#E6EDF3")
	colorMuted   = lipgloss.Color("#9BA7B4")
	colorDim     = lipgloss.Color("#6B7785")
	colorBorder  = lipgloss.Color("#243041")
	colorSuccess = lipgloss.Color("#2ECC71")
	colorWarning = lipgloss.Color("#F39C12")
	colorError   = lipgloss.Color("#E74C3C")
)

// ConnectionForm is the connection screen component.
type ConnectionForm struct {
	clusterInput textinput.Model
	addressInput textinput.Model
	focused      int // 0=cluster, 1=address, 2=button
	status       int // 0=disconnected, 1=connecting, 2=connected
	errorMsg     string
	width        int
}

// NewConnectionForm creates a new connection form.
func NewConnectionForm() ConnectionForm {
	ci := textinput.New()
	ci.Placeholder = "0"
	ci.SetValue("0")
	ci.CharLimit = 39 // uint128 max
	ci.Width = 30
	ci.Focus()
	ci.PromptStyle = lipgloss.NewStyle().Foreground(colorAccent)
	ci.TextStyle = lipgloss.NewStyle().Foreground(colorText)
	ci.PlaceholderStyle = lipgloss.NewStyle().Foreground(colorDim)
	ci.Cursor.Style = lipgloss.NewStyle().Foreground(colorAccent)

	ai := textinput.New()
	ai.Placeholder = "3000"
	ai.SetValue("3000")
	ai.CharLimit = 64
	ai.Width = 30
	ai.PromptStyle = lipgloss.NewStyle().Foreground(colorAccent)
	ai.TextStyle = lipgloss.NewStyle().Foreground(colorText)
	ai.PlaceholderStyle = lipgloss.NewStyle().Foreground(colorDim)
	ai.Cursor.Style = lipgloss.NewStyle().Foreground(colorAccent)

	return ConnectionForm{
		clusterInput: ci,
		addressInput: ai,
		focused:      0,
	}
}

// SetWidth sets the available width.
func (f *ConnectionForm) SetWidth(w int) {
	f.width = w
}

// SetStatus sets the connection status (0=disconnected, 1=connecting, 2=connected).
func (f *ConnectionForm) SetStatus(s int) {
	f.status = s
}

// SetError sets the error message.
func (f *ConnectionForm) SetError(msg string) {
	f.errorMsg = msg
}

// ClusterID returns the current cluster ID value.
func (f *ConnectionForm) ClusterID() string {
	return f.clusterInput.Value()
}

// Address returns the current address value.
func (f *ConnectionForm) Address() string {
	return f.addressInput.Value()
}

// FocusNext moves focus to the next field.
func (f *ConnectionForm) FocusNext() {
	f.focused = (f.focused + 1) % 3
	f.updateFocus()
}

// FocusPrev moves focus to the previous field.
func (f *ConnectionForm) FocusPrev() {
	f.focused = (f.focused + 2) % 3
	f.updateFocus()
}

// IsButtonFocused returns true if the connect button is focused.
func (f *ConnectionForm) IsButtonFocused() bool {
	return f.focused == 2
}

func (f *ConnectionForm) updateFocus() {
	f.clusterInput.Blur()
	f.addressInput.Blur()

	switch f.focused {
	case 0:
		f.clusterInput.Focus()
	case 1:
		f.addressInput.Focus()
	}
}

// Update handles input for the connection form.
func (f *ConnectionForm) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch f.focused {
	case 0:
		var cmd tea.Cmd
		f.clusterInput, cmd = f.clusterInput.Update(msg)
		cmds = append(cmds, cmd)
	case 1:
		var cmd tea.Cmd
		f.addressInput, cmd = f.addressInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}

// View renders the connection form.
func (f *ConnectionForm) View() string {
	labelStyle := lipgloss.NewStyle().Foreground(colorMuted).Width(14)
	accentBold := lipgloss.NewStyle().Foreground(colorAccent).Bold(true)
	dimStyle := lipgloss.NewStyle().Foreground(colorDim)
	errStyle := lipgloss.NewStyle().Foreground(colorError)

	var sb strings.Builder

	// Logo
	logo := `
 ████████╗██╗ ██████╗ ███████╗██████╗       ████████╗██╗   ██╗██╗
 ╚══██╔══╝██║██╔════╝ ██╔════╝██╔══██╗      ╚══██╔══╝██║   ██║██║
    ██║   ██║██║  ███╗█████╗  ██████╔╝ ─────   ██║   ██║   ██║██║
    ██║   ██║██║   ██║██╔══╝  ██╔══██╗         ██║   ██║   ██║██║
    ██║   ██║╚██████╔╝███████╗██║  ██║         ██║   ╚██████╔╝██║
    ╚═╝   ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝         ╚═╝    ╚═════╝ ╚═╝`
	sb.WriteString(accentBold.Render(logo))
	sb.WriteString("\n")
	sb.WriteString(dimStyle.Render("              the best client for tigerbeetle"))
	sb.WriteString("\n\n")

	// Form box
	formWidth := 48
	var form strings.Builder

	// Cluster ID field
	form.WriteString(labelStyle.Render("Cluster ID:"))
	form.WriteString(" ")
	form.WriteString(f.clusterInput.View())
	form.WriteString("\n\n")

	// Address field
	form.WriteString(labelStyle.Render("Address:"))
	form.WriteString(" ")
	form.WriteString(f.addressInput.View())
	form.WriteString("\n\n")

	// Connect button
	var btnText string
	switch f.status {
	case 1:
		btnText = lipgloss.NewStyle().Foreground(colorWarning).Bold(true).Render("  Connecting...  ")
	default:
		if f.focused == 2 {
			btnText = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#0B0F14")).
				Background(colorAccent).
				Bold(true).
				Padding(0, 2).
				Render("● Connect")
		} else {
			btnText = lipgloss.NewStyle().
				Foreground(colorAccent).
				Border(lipgloss.NormalBorder()).
				BorderForeground(colorBorder).
				Padding(0, 2).
				Render("● Connect")
		}
	}
	// Center the button within the form width
	btnWidth := lipgloss.Width(btnText)
	innerWidth := formWidth - 6 // account for box padding + border
	btnPad := (innerWidth - btnWidth) / 2
	if btnPad < 0 {
		btnPad = 0
	}
	form.WriteString(fmt.Sprintf("%s%s", strings.Repeat(" ", btnPad), btnText))

	// Error message
	if f.errorMsg != "" {
		form.WriteString("\n\n")
		form.WriteString(errStyle.Render("  " + f.errorMsg))
	}

	// Render form in a box
	formBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorBorder).
		Padding(1, 2).
		Width(formWidth).
		Render(form.String())

	sb.WriteString(formBox)

	return sb.String()
}
