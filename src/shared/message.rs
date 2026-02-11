//! TEA messages and the central `update` reducer.

use crossterm::event::{KeyCode, KeyEvent, KeyModifiers};

use super::app_state::{AppState, ConnectionField, ConnectionStatus, Screen, StatusMessage};

/// Every user-intent and async result is expressed as a `Message`.
#[derive(Debug)]
#[expect(dead_code, reason = "Tick constructed by the event loop")]
pub enum Message {
    Quit,
    Tick,
    FocusNext,
    FocusPrev,
    CharInput(char),
    Backspace,
    Submit,
}

// -- Key → Message mapping --------------------------------------------------

/// Maps a raw key event to a domain `Message` based on the current screen.
pub fn handle_key(key: KeyEvent, app: &AppState) -> Option<Message> {
    // Ctrl+C always quits
    if key.code == KeyCode::Char('c') && key.modifiers.contains(KeyModifiers::CONTROL) {
        return Some(Message::Quit);
    }

    match app.screen {
        Screen::Connection => handle_connection_key(key, app),
        Screen::Dashboard => handle_dashboard_key(key),
    }
}

fn handle_connection_key(key: KeyEvent, app: &AppState) -> Option<Message> {
    let editing = app.connection.focused.is_text_input();

    match key.code {
        KeyCode::Tab => Some(Message::FocusNext),
        KeyCode::BackTab => Some(Message::FocusPrev),
        KeyCode::Enter => Some(Message::Submit),
        KeyCode::Esc => Some(Message::Quit),
        KeyCode::Char('q') if !editing => Some(Message::Quit),
        KeyCode::Char(c) if editing => Some(Message::CharInput(c)),
        KeyCode::Backspace if editing => Some(Message::Backspace),
        _ => None,
    }
}

fn handle_dashboard_key(key: KeyEvent) -> Option<Message> {
    match key.code {
        KeyCode::Char('q') | KeyCode::Esc => Some(Message::Quit),
        KeyCode::Tab => Some(Message::FocusNext),
        KeyCode::BackTab => Some(Message::FocusPrev),
        _ => None,
    }
}

// -- Update (TEA) -----------------------------------------------------------

/// Central reducer: applies a `Message` to `AppState`.
pub fn update(app: &mut AppState, msg: &Message) {
    match msg {
        Message::Quit => app.should_quit = true,
        Message::Tick => {}
        msg => match app.screen {
            Screen::Connection => update_connection(app, msg),
            Screen::Dashboard => update_dashboard(app, msg),
        },
    }
}

fn update_connection(app: &mut AppState, msg: &Message) {
    match msg {
        Message::CharInput(c) => match app.connection.focused {
            ConnectionField::ClusterId => app.connection.cluster_id.push(*c),
            ConnectionField::Address => app.connection.address.push(*c),
            ConnectionField::ConnectButton => {}
        },
        Message::Backspace => match app.connection.focused {
            ConnectionField::ClusterId => {
                app.connection.cluster_id.pop();
            }
            ConnectionField::Address => {
                app.connection.address.pop();
            }
            ConnectionField::ConnectButton => {}
        },
        Message::FocusNext => app.connection.focused = app.connection.focused.next(),
        Message::FocusPrev => app.connection.focused = app.connection.focused.prev(),
        Message::Submit => {
            if matches!(app.connection.focused, ConnectionField::ConnectButton) {
                // Phase 2: send TbCommand::Connect via channel.
                // For now, simulate a successful connection.
                app.connection.status = ConnectionStatus::Connected;
                app.screen = Screen::Dashboard;
                app.status = Some(StatusMessage::success("Connected (simulated)"));
            }
        }
        _ => {}
    }
}

fn update_dashboard(app: &mut AppState, msg: &Message) {
    match msg {
        Message::FocusNext => app.tab = app.tab.next(),
        Message::FocusPrev => app.tab = app.tab.prev(),
        _ => {}
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::shared::app_state::Tab;

    fn default_app() -> AppState {
        AppState::default()
    }

    #[test]
    fn quit_sets_flag() {
        let mut app = default_app();
        update(&mut app, &Message::Quit);
        assert!(app.should_quit);
    }

    #[test]
    fn connection_focus_cycles() {
        let mut app = default_app();
        assert_eq!(app.connection.focused, ConnectionField::ClusterId);
        update(&mut app, &Message::FocusNext);
        assert_eq!(app.connection.focused, ConnectionField::Address);
        update(&mut app, &Message::FocusNext);
        assert_eq!(app.connection.focused, ConnectionField::ConnectButton);
        update(&mut app, &Message::FocusNext);
        assert_eq!(app.connection.focused, ConnectionField::ClusterId);
    }

    #[test]
    fn connection_char_input_and_backspace() {
        let mut app = default_app();
        // Default cluster_id is "0"
        update(&mut app, &Message::CharInput('1'));
        assert_eq!(app.connection.cluster_id, "01");
        update(&mut app, &Message::Backspace);
        assert_eq!(app.connection.cluster_id, "0");
    }

    #[test]
    fn submit_on_button_transitions_to_dashboard() {
        let mut app = default_app();
        app.connection.focused = ConnectionField::ConnectButton;
        update(&mut app, &Message::Submit);
        assert_eq!(app.screen, Screen::Dashboard);
        assert_eq!(app.connection.status, ConnectionStatus::Connected);
    }

    #[test]
    fn submit_on_input_does_nothing() {
        let mut app = default_app();
        assert_eq!(app.connection.focused, ConnectionField::ClusterId);
        update(&mut app, &Message::Submit);
        assert_eq!(app.screen, Screen::Connection);
    }

    #[test]
    fn dashboard_tab_cycles() {
        let mut app = default_app();
        app.screen = Screen::Dashboard;
        assert_eq!(app.tab, Tab::Accounts);
        update(&mut app, &Message::FocusNext);
        assert_eq!(app.tab, Tab::Transfers);
        update(&mut app, &Message::FocusNext);
        assert_eq!(app.tab, Tab::BalanceSheet);
        update(&mut app, &Message::FocusNext);
        assert_eq!(app.tab, Tab::Accounts);
    }

    #[test]
    fn dashboard_tab_cycles_prev() {
        let mut app = default_app();
        app.screen = Screen::Dashboard;
        update(&mut app, &Message::FocusPrev);
        assert_eq!(app.tab, Tab::BalanceSheet);
    }
}
