//! Global application state (TEA Model).

/// Root state for the entire application.
#[derive(Debug)]
pub struct AppState {
    pub screen: Screen,
    pub tab: Tab,
    pub should_quit: bool,
    pub connection: ConnectionForm,
    pub status: Option<StatusMessage>,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Screen {
    Connection,
    Dashboard,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Tab {
    Accounts,
    Transfers,
    BalanceSheet,
}

impl Tab {
    pub const ALL: [Self; 3] = [Self::Accounts, Self::Transfers, Self::BalanceSheet];

    pub fn label(self) -> &'static str {
        match self {
            Self::Accounts => "Accounts",
            Self::Transfers => "Transfers",
            Self::BalanceSheet => "Balance Sheet",
        }
    }

    pub fn index(self) -> usize {
        match self {
            Self::Accounts => 0,
            Self::Transfers => 1,
            Self::BalanceSheet => 2,
        }
    }

    pub fn next(self) -> Self {
        match self {
            Self::Accounts => Self::Transfers,
            Self::Transfers => Self::BalanceSheet,
            Self::BalanceSheet => Self::Accounts,
        }
    }

    pub fn prev(self) -> Self {
        match self {
            Self::Accounts => Self::BalanceSheet,
            Self::Transfers => Self::Accounts,
            Self::BalanceSheet => Self::Transfers,
        }
    }
}

// -- Connection form --------------------------------------------------------

/// Input state for the connection screen.
#[derive(Debug)]
pub struct ConnectionForm {
    pub cluster_id: String,
    pub address: String,
    pub focused: ConnectionField,
    pub status: ConnectionStatus,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ConnectionField {
    ClusterId,
    Address,
    ConnectButton,
}

impl ConnectionField {
    pub fn next(self) -> Self {
        match self {
            Self::ClusterId => Self::Address,
            Self::Address => Self::ConnectButton,
            Self::ConnectButton => Self::ClusterId,
        }
    }

    pub fn prev(self) -> Self {
        match self {
            Self::ClusterId => Self::ConnectButton,
            Self::Address => Self::ClusterId,
            Self::ConnectButton => Self::Address,
        }
    }

    pub fn is_text_input(self) -> bool {
        matches!(self, Self::ClusterId | Self::Address)
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
#[expect(
    dead_code,
    reason = "Connecting used in Phase 2 when TB client connects"
)]
pub enum ConnectionStatus {
    Disconnected,
    Connecting,
    Connected,
}

// -- Status bar messages ----------------------------------------------------

/// Feedback message shown in the status bar.
#[derive(Debug, Clone)]
pub struct StatusMessage {
    pub text: String,
    pub level: StatusLevel,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[expect(dead_code, reason = "all variants used once status feedback is wired")]
pub enum StatusLevel {
    Info,
    Success,
    Warning,
    Error,
}

#[expect(
    dead_code,
    reason = "constructors used once error/info feedback is wired"
)]
impl StatusMessage {
    pub fn info(text: impl Into<String>) -> Self {
        Self {
            text: text.into(),
            level: StatusLevel::Info,
        }
    }

    pub fn success(text: impl Into<String>) -> Self {
        Self {
            text: text.into(),
            level: StatusLevel::Success,
        }
    }

    pub fn error(text: impl Into<String>) -> Self {
        Self {
            text: text.into(),
            level: StatusLevel::Error,
        }
    }
}

// -- Defaults ---------------------------------------------------------------

impl Default for AppState {
    fn default() -> Self {
        Self {
            screen: Screen::Connection,
            tab: Tab::Accounts,
            should_quit: false,
            connection: ConnectionForm {
                cluster_id: "0".into(),
                address: "3000".into(),
                focused: ConnectionField::ClusterId,
                status: ConnectionStatus::Disconnected,
            },
            status: None,
        }
    }
}
