use ratatui::{
    Frame,
    layout::{Alignment, Constraint, Layout},
    style::{Modifier, Style},
    text::{Line, Span},
    widgets::{Block, BorderType, Paragraph, Tabs},
};

use super::status_bar::render_status_bar;
use super::theme::THEME;
use crate::shared::app_state::{AppState, Tab};

pub fn render_dashboard(app: &AppState, frame: &mut Frame) {
    let area = frame.area();

    // Background
    frame.render_widget(Block::default().style(Style::default().bg(THEME.bg)), area);

    let [title_area, tabs_area, content_area, status_area] = Layout::vertical([
        Constraint::Length(1), // title bar
        Constraint::Length(2), // tabs
        Constraint::Fill(1),   // main content
        Constraint::Length(1), // status bar
    ])
    .areas(area);

    // -- Title bar ----------------------------------------------------------
    let title = Line::from(vec![
        Span::styled(
            " tiger-tui ",
            Style::default()
                .fg(THEME.accent)
                .add_modifier(Modifier::BOLD),
        ),
        Span::styled("──", Style::default().fg(THEME.border)),
        Span::styled(
            format!(
                " ● {}:{} ",
                app.connection.cluster_id, app.connection.address
            ),
            Style::default().fg(THEME.success),
        ),
    ]);
    frame.render_widget(Paragraph::new(title), title_area);

    // -- Tabs ---------------------------------------------------------------
    let tab_titles: Vec<Line> = Tab::ALL.iter().map(|t| Line::from(t.label())).collect();
    let tabs = Tabs::new(tab_titles)
        .select(app.tab.index())
        .highlight_style(
            Style::default()
                .fg(THEME.accent)
                .add_modifier(Modifier::BOLD),
        )
        .style(Style::default().fg(THEME.text_muted))
        .divider(Span::styled(" │ ", Style::default().fg(THEME.border)));
    frame.render_widget(tabs, tabs_area);

    // -- Content area (placeholder for Phase 3) -----------------------------
    let placeholder = match app.tab {
        Tab::Accounts => "Accounts table — Phase 3",
        Tab::Transfers => "Transfers table — Phase 3",
        Tab::BalanceSheet => "Balance Sheet — Phase 3",
    };

    let block = Block::bordered()
        .title(format!(" {} ", app.tab.label()))
        .title_style(Style::default().fg(THEME.text_muted))
        .border_type(BorderType::Rounded)
        .border_style(Style::default().fg(THEME.border));

    let content = Paragraph::new(placeholder)
        .style(Style::default().fg(THEME.text_dim))
        .alignment(Alignment::Center)
        .block(block);
    frame.render_widget(content, content_area);

    // -- Status bar ---------------------------------------------------------
    render_status_bar(app, frame, status_area);
}
