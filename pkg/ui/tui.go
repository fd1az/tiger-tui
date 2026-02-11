package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/fd1az/tiger-tui/pkg/ui/components"
)

// Model is the main Bubble Tea model for tiger-tui.
type Model struct {
	// Components
	connForm  components.ConnectionForm
	dashboard components.Dashboard
	statusBar components.StatusBar

	// State
	screen    Screen
	connStatus ConnectionStatus
	keys      KeyMap
	width     int
	height    int
	ready     bool
	quitting  bool
}

// New creates a new TUI model.
func New() Model {
	return Model{
		connForm:  components.NewConnectionForm(),
		dashboard: components.NewDashboard(),
		statusBar: components.NewStatusBar(),
		screen:    ScreenConnection,
		keys:      DefaultKeyMap(),
	}
}

// Init initializes the TUI model.
func (m Model) Init() tea.Cmd {
	return textinputBlink()
}

// textinputBlink returns a command that starts the textinput cursor blinking.
func textinputBlink() tea.Cmd {
	return nil
}

// Update handles messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		m.connForm.SetWidth(msg.Width)
		m.dashboard.SetSize(msg.Width, msg.Height)
		m.statusBar.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		// Global: always allow quit
		if key.Matches(msg, m.keys.Quit) {
			m.quitting = true
			return m, tea.Quit
		}

		// Route to screen-specific handler
		switch m.screen {
		case ScreenConnection:
			return m.updateConnection(msg)
		case ScreenDashboard:
			return m.updateDashboard(msg)
		}

	// --- App messages ---
	case ConnectedMsg:
		m.connStatus = Connected
		m.screen = ScreenDashboard
		m.connForm.SetStatus(2)
		m.statusBar.SetConnection(2, m.connForm.ClusterID(), m.connForm.Address())
		m.statusBar.SetMessage("Connected to TigerBeetle", 1)
		return m, nil

	case ConnectionFailedMsg:
		m.connStatus = Disconnected
		m.connForm.SetStatus(0)
		m.connForm.SetError(msg.Err.Error())
		m.statusBar.SetMessage(fmt.Sprintf("Connection failed: %s", msg.Err), 3)
		return m, nil

	case ErrorMsg:
		m.statusBar.SetMessage(msg.Err.Error(), 3)
		return m, nil

	case StatusMsg:
		m.statusBar.SetMessage(msg.Text, int(msg.Level))
		return m, nil
	}

	return m, nil
}

// updateConnection handles keys on the connection screen.
func (m Model) updateConnection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Tab):
		m.connForm.FocusNext()
		return m, nil

	case key.Matches(msg, m.keys.ShiftTab):
		m.connForm.FocusPrev()
		return m, nil

	case key.Matches(msg, m.keys.Enter):
		if m.connForm.IsButtonFocused() {
			// Validate and submit
			if m.connForm.ClusterID() == "" {
				m.connForm.SetError("Cluster ID is required")
				return m, nil
			}
			if m.connForm.Address() == "" {
				m.connForm.SetError("Address is required")
				return m, nil
			}

			m.connStatus = Connecting
			m.connForm.SetStatus(1)
			m.connForm.SetError("")
			m.statusBar.SetMessage("Connecting...", 0)

			// Phase 1: simulate connect (Phase 2 will do real TB connection)
			return m, func() tea.Msg {
				return ConnectedMsg{}
			}
		}
		// Enter on text fields moves to next
		m.connForm.FocusNext()
		return m, nil

	case msg.String() == "q":
		// On connection screen, q in a text field types 'q'; on button it quits
		if m.connForm.IsButtonFocused() {
			m.quitting = true
			return m, tea.Quit
		}
	}

	// Forward to text inputs
	cmd := m.connForm.Update(msg)
	return m, cmd
}

// updateDashboard handles keys on the dashboard screen.
func (m Model) updateDashboard(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Tab):
		m.dashboard.NextTab()
		return m, nil

	case key.Matches(msg, m.keys.ShiftTab):
		m.dashboard.PrevTab()
		return m, nil

	case key.Matches(msg, m.keys.Escape):
		// Return to connection screen
		m.screen = ScreenConnection
		m.connStatus = Disconnected
		m.connForm.SetStatus(0)
		m.statusBar.SetConnection(0, "", "")
		return m, nil

	case msg.String() == "q":
		m.quitting = true
		return m, tea.Quit
	}

	return m, nil
}

// View renders the TUI.
func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if !m.ready {
		return "\n  Initializing..."
	}

	var content string

	switch m.screen {
	case ScreenConnection:
		content = m.viewConnection()
	case ScreenDashboard:
		content = m.viewDashboard()
	}

	return content
}

// viewConnection renders the connection screen.
func (m Model) viewConnection() string {
	// Center the form
	formContent := m.connForm.View()

	// Center horizontally
	formWidth := lipgloss.Width(formContent)
	leftPad := (m.width - formWidth) / 2
	if leftPad < 0 {
		leftPad = 0
	}

	// Center vertically
	formHeight := lipgloss.Height(formContent)
	topPad := (m.height - formHeight - 3) / 2 // -3 for status bar
	if topPad < 1 {
		topPad = 1
	}

	var sb strings.Builder
	sb.WriteString(strings.Repeat("\n", topPad))
	for _, line := range strings.Split(formContent, "\n") {
		sb.WriteString(strings.Repeat(" ", leftPad))
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	// Fill remaining space, then status bar
	currentHeight := topPad + formHeight + 1
	remaining := m.height - currentHeight - 1
	if remaining > 0 {
		sb.WriteString(strings.Repeat("\n", remaining))
	}

	sb.WriteString(m.statusBar.View())

	return sb.String()
}

// viewDashboard renders the dashboard screen.
func (m Model) viewDashboard() string {
	// Top bar
	topBar := m.renderTopBar()

	var sb strings.Builder
	sb.WriteString(topBar)
	sb.WriteString("\n\n")
	sb.WriteString(m.dashboard.View())

	// Fill remaining space, then status bar
	currentHeight := lipgloss.Height(sb.String()) + 1
	remaining := m.height - currentHeight - 1
	if remaining > 0 {
		sb.WriteString(strings.Repeat("\n", remaining))
	}
	sb.WriteString(m.statusBar.View())

	return sb.String()
}

// renderTopBar renders the dashboard top bar.
func (m Model) renderTopBar() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(ColorAccent)
	title := titleStyle.Render(" tiger-tui ")

	connStyle := StatusConnected
	connText := connStyle.Render(fmt.Sprintf("● Connected %s:%s", m.connForm.ClusterID(), m.connForm.Address()))

	gap := m.width - lipgloss.Width(title) - lipgloss.Width(connText) - 2
	if gap < 1 {
		gap = 1
	}

	return title + strings.Repeat(" ", gap) + connText
}

// Program holds the Bubble Tea program instance for external access.
var Program *tea.Program

// Run starts the Bubble Tea program.
func Run() error {
	Program = tea.NewProgram(New(), tea.WithAltScreen())
	_, err := Program.Run()
	return err
}

// Send sends a message to the running program.
func Send(msg tea.Msg) {
	if Program != nil {
		Program.Send(msg)
	}
}
