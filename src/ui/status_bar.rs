use ratatui::{
    Frame,
    layout::Rect,
    layout::{Constraint, Layout},
    style::Style,
    text::Span,
    widgets::Paragraph,
};

use super::theme::THEME;
use crate::shared::app_state::{AppState, ConnectionStatus, Screen, StatusLevel};

pub fn render_status_bar(app: &AppState, frame: &mut Frame, area: Rect) {
    let [left_area, right_area] =
        Layout::horizontal([Constraint::Fill(1), Constraint::Fill(1)]).areas(area);

    // -- Left side: version or connection info ------------------------------
    let left_text = match app.screen {
        Screen::Connection => {
            format!(" tiger-tui v{}", env!("CARGO_PKG_VERSION"))
        }
        Screen::Dashboard => match &app.connection.status {
            ConnectionStatus::Connected => {
                format!(
                    " ● {}:{}",
                    app.connection.cluster_id, app.connection.address
                )
            }
            ConnectionStatus::Connecting => " ◌ Connecting…".into(),
            ConnectionStatus::Disconnected => " ○ Disconnected".into(),
        },
    };

    let left_style = match app.screen {
        Screen::Dashboard if app.connection.status == ConnectionStatus::Connected => {
            Style::default().fg(THEME.success)
        }
        _ => Style::default().fg(THEME.text_dim),
    };
    frame.render_widget(
        Paragraph::new(Span::styled(left_text, left_style)),
        left_area,
    );

    // -- Right side: contextual status or key hints -------------------------
    let right_text = if let Some(msg) = &app.status {
        msg.text.clone()
    } else {
        match app.screen {
            Screen::Connection => "Tab Navigate  Enter Connect  q Quit ".into(),
            Screen::Dashboard => "Tab Switch  r Refresh  ? Help  q Quit ".into(),
        }
    };

    let right_style = if let Some(msg) = &app.status {
        match msg.level {
            StatusLevel::Info => Style::default().fg(THEME.info),
            StatusLevel::Success => Style::default().fg(THEME.success),
            StatusLevel::Warning => Style::default().fg(THEME.warning),
            StatusLevel::Error => Style::default().fg(THEME.error),
        }
    } else {
        Style::default().fg(THEME.text_dim)
    };

    frame.render_widget(
        Paragraph::new(Span::styled(right_text, right_style))
            .alignment(ratatui::layout::Alignment::Right),
        right_area,
    );
}
