package components

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// StatusBar renders the bottom status bar.
type StatusBar struct {
	connectionStatus int    // 0=disconnected, 1=connecting, 2=connected
	clusterID        string
	address          string
	message          string
	messageLevel     int // 0=info, 1=success, 2=warning, 3=error
	messageTime      time.Time
	width            int
}

// NewStatusBar creates a new status bar.
func NewStatusBar() StatusBar {
	return StatusBar{}
}

// SetConnection updates the connection info displayed.
func (s *StatusBar) SetConnection(status int, clusterID, address string) {
	s.connectionStatus = status
	s.clusterID = clusterID
	s.address = address
}

// SetMessage sets a temporary message.
func (s *StatusBar) SetMessage(text string, level int) {
	s.message = text
	s.messageLevel = level
	s.messageTime = time.Now()
}

// SetWidth sets the available width.
func (s *StatusBar) SetWidth(w int) {
	s.width = w
}

// View renders the status bar.
func (s *StatusBar) View() string {
	barStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Width(s.width)

	dimStyle := lipgloss.NewStyle().Foreground(colorDim)

	var parts []string

	// Connection status
	switch s.connectionStatus {
	case 2: // Connected
		connStyle := lipgloss.NewStyle().Foreground(colorSuccess).Bold(true)
		parts = append(parts, connStyle.Render(fmt.Sprintf("● Connected %s:%s", s.clusterID, s.address)))
	case 1: // Connecting
		connStyle := lipgloss.NewStyle().Foreground(colorWarning)
		parts = append(parts, connStyle.Render("○ Connecting..."))
	default: // Disconnected
		parts = append(parts, dimStyle.Render("○ Disconnected"))
	}

	// Status message (show for 10 seconds)
	if s.message != "" && time.Since(s.messageTime) < 10*time.Second {
		var style lipgloss.Style
		switch s.messageLevel {
		case 1:
			style = lipgloss.NewStyle().Foreground(colorSuccess)
		case 2:
			style = lipgloss.NewStyle().Foreground(colorWarning)
		case 3:
			style = lipgloss.NewStyle().Foreground(colorError)
		default:
			style = lipgloss.NewStyle().Foreground(colorMuted)
		}
		parts = append(parts, style.Render(s.message))
	}

	left := strings.Join(parts, "  │  ")

	// Right side: version + help hint
	right := dimStyle.Render("? Help  q Quit")

	// Calculate spacing
	leftLen := lipgloss.Width(left)
	rightLen := lipgloss.Width(right)
	gap := s.width - leftLen - rightLen - 2
	if gap < 1 {
		gap = 1
	}

	return barStyle.Render(left + strings.Repeat(" ", gap) + right)
}
