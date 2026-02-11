package ui

import "github.com/fd1az/tiger-tui/business/connection/infra"

// ConnectedMsg signals successful TigerBeetle connection.
type ConnectedMsg struct {
	Client *infra.Client
}

// ConnectionFailedMsg signals a failed connection attempt.
type ConnectionFailedMsg struct {
	Err error
}

// ErrorMsg is sent when an error occurs.
type ErrorMsg struct {
	Err error
}

// StatusMsg sets a temporary status bar message.
type StatusMsg struct {
	Text  string
	Level StatusLevel
}

// StatusLevel represents the severity of a status message.
type StatusLevel int

const (
	StatusInfo StatusLevel = iota
	StatusSuccess
	StatusWarning
	StatusError
)
