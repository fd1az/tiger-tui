package ui

// Screen represents the current screen.
type Screen int

const (
	ScreenConnection Screen = iota
	ScreenDashboard
)

// Tab represents the active dashboard tab.
type Tab int

const (
	TabAccounts Tab = iota
	TabTransfers
	TabBalanceSheet
)

// String returns the tab display name.
func (t Tab) String() string {
	switch t {
	case TabAccounts:
		return "Accounts"
	case TabTransfers:
		return "Transfers"
	case TabBalanceSheet:
		return "Balance Sheet"
	default:
		return ""
	}
}

// AllTabs returns all tabs in order.
func AllTabs() []Tab {
	return []Tab{TabAccounts, TabTransfers, TabBalanceSheet}
}

// Next returns the next tab (wrapping).
func (t Tab) Next() Tab {
	return (t + 1) % Tab(len(AllTabs()))
}

// Prev returns the previous tab (wrapping).
func (t Tab) Prev() Tab {
	n := Tab(len(AllTabs()))
	return (t - 1 + n) % n
}

// ConnectionStatus represents the state of the TB connection.
type ConnectionStatus int

const (
	Disconnected ConnectionStatus = iota
	Connecting
	Connected
)

// String returns a display string for the connection status.
func (s ConnectionStatus) String() string {
	switch s {
	case Disconnected:
		return "Disconnected"
	case Connecting:
		return "Connecting..."
	case Connected:
		return "Connected"
	default:
		return ""
	}
}

// ConnectionField is the focused field in the connection form.
type ConnectionField int

const (
	FieldClusterID ConnectionField = iota
	FieldAddress
	FieldConnectButton
)

// Next returns the next field (wrapping).
func (f ConnectionField) Next() ConnectionField {
	return (f + 1) % 3
}

// Prev returns the previous field (wrapping).
func (f ConnectionField) Prev() ConnectionField {
	return (f + 2) % 3
}
