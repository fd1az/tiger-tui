//! Top-level UI layer: view dispatcher, theme, and shared widgets.

pub mod dashboard;
pub mod status_bar;
pub mod theme;

use ratatui::Frame;

use crate::shared::app_state::{AppState, Screen};

/// Routes rendering to the active screen.
pub fn view(app: &AppState, frame: &mut Frame) {
    match app.screen {
        Screen::Connection => {
            crate::modules::connection::ui::connection_view::render_connection(app, frame);
        }
        Screen::Dashboard => {
            dashboard::render_dashboard(app, frame);
        }
    }
}
