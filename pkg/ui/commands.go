package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/fd1az/tiger-tui/business/connection/infra"
)

// ConnectCmd returns a tea.Cmd that connects to TigerBeetle.
func ConnectCmd(clusterID string, address string) tea.Cmd {
	return func() tea.Msg {
		client, err := infra.Connect(clusterID, []string{address})
		if err != nil {
			return ConnectionFailedMsg{Err: err}
		}
		return ConnectedMsg{Client: client}
	}
}
