package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Dashboard renders the main dashboard shell with tabs.
type Dashboard struct {
	activeTab int // 0=Accounts, 1=Transfers, 2=Balance Sheet
	width     int
	height    int
}

var tabNames = []string{"Accounts", "Transfers", "Balance Sheet"}

// NewDashboard creates a new dashboard.
func NewDashboard() Dashboard {
	return Dashboard{}
}

// SetSize sets the available dimensions.
func (d *Dashboard) SetSize(w, h int) {
	d.width = w
	d.height = h
}

// ActiveTab returns the current active tab index.
func (d *Dashboard) ActiveTab() int {
	return d.activeTab
}

// NextTab cycles to the next tab.
func (d *Dashboard) NextTab() {
	d.activeTab = (d.activeTab + 1) % len(tabNames)
}

// PrevTab cycles to the previous tab.
func (d *Dashboard) PrevTab() {
	d.activeTab = (d.activeTab + len(tabNames) - 1) % len(tabNames)
}

// View renders the dashboard.
func (d *Dashboard) View() string {
	activeStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorAccent).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(colorAccent).
		Padding(0, 2)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(colorDim).
		Padding(0, 2)

	// Render tabs
	var tabs []string
	for i, name := range tabNames {
		if i == d.activeTab {
			tabs = append(tabs, activeStyle.Render(name))
		} else {
			tabs = append(tabs, inactiveStyle.Render(name))
		}
	}
	tabBar := lipgloss.JoinHorizontal(lipgloss.Bottom, tabs...)

	// Content area (placeholder)
	dimStyle := lipgloss.NewStyle().Foreground(colorDim)
	var content string
	switch d.activeTab {
	case 0:
		content = dimStyle.Render("  Accounts will appear here after connecting.")
	case 1:
		content = dimStyle.Render("  Transfers will appear here after connecting.")
	case 2:
		content = dimStyle.Render("  Balance Sheet will appear here after connecting.")
	}

	// Content box
	contentHeight := d.height - 6 // Reserve space for tabs + status
	if contentHeight < 3 {
		contentHeight = 3
	}
	contentBox := lipgloss.NewStyle().
		Width(d.width - 4).
		Height(contentHeight).
		Render(content)

	var sb strings.Builder
	sb.WriteString(tabBar)
	sb.WriteString("\n\n")
	sb.WriteString(contentBox)

	return sb.String()
}
