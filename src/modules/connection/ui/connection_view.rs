use ratatui::{
    Frame,
    layout::{Alignment, Constraint, Layout, Rect},
    style::{Modifier, Style},
    text::{Line, Span},
    widgets::{Block, BorderType, Paragraph},
};

use crate::shared::app_state::{AppState, ConnectionField};
use crate::ui::theme::THEME;

const LOGO: &str = r"
 ████████╗██╗ ██████╗ ███████╗██████╗       ████████╗██╗   ██╗██╗
 ╚══██╔══╝██║██╔════╝ ██╔════╝██╔══██╗      ╚══██╔══╝██║   ██║██║
    ██║   ██║██║  ███╗█████╗  ██████╔╝  █████╗ ██║   ██║   ██║██║
    ██║   ██║██║   ██║██╔══╝  ██╔══██╗  ╚════╝ ██║   ██║   ██║██║
    ██║   ██║╚██████╔╝███████╗██║  ██║         ██║   ╚██████╔╝██║
    ╚═╝   ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝         ╚═╝    ╚═════╝ ╚═╝";

pub fn render_connection(app: &AppState, frame: &mut Frame) {
    let area = frame.area();

    // Full-screen background
    frame.render_widget(Block::default().style(Style::default().bg(THEME.bg)), area);

    let [_, logo_area, _, subtitle_area, _, form_area, _] = Layout::vertical([
        Constraint::Min(1),    // top pad
        Constraint::Length(8), // logo
        Constraint::Length(1), // gap
        Constraint::Length(1), // subtitle
        Constraint::Length(1), // gap
        Constraint::Length(9), // form block
        Constraint::Min(1),    // bottom pad
    ])
    .areas(area);

    // -- Logo ---------------------------------------------------------------
    let logo = Paragraph::new(LOGO)
        .style(
            Style::default()
                .fg(THEME.accent)
                .add_modifier(Modifier::BOLD),
        )
        .alignment(Alignment::Center);
    frame.render_widget(logo, logo_area);

    // -- Subtitle -----------------------------------------------------------
    let subtitle = Paragraph::new("t i g e r  -  t u i")
        .style(Style::default().fg(THEME.text_dim))
        .alignment(Alignment::Center);
    frame.render_widget(subtitle, subtitle_area);

    // -- Form (centered 46-col box) -----------------------------------------
    let form_rect = center_h(form_area, 46);
    render_form(app, frame, form_rect);
}

fn render_form(app: &AppState, frame: &mut Frame, area: Rect) {
    let focused_field = app.connection.focused;

    let block = Block::bordered()
        .title(" Connection ")
        .title_style(
            Style::default()
                .fg(THEME.accent)
                .add_modifier(Modifier::BOLD),
        )
        .border_type(BorderType::Rounded)
        .border_style(Style::default().fg(THEME.border));
    let inner = block.inner(area);
    frame.render_widget(block, area);

    let [_, row_cid, _, row_addr, _, row_btn, _] = Layout::vertical([
        Constraint::Length(1), // pad
        Constraint::Length(1), // cluster id
        Constraint::Length(1), // gap
        Constraint::Length(1), // address
        Constraint::Length(1), // gap
        Constraint::Length(1), // button
        Constraint::Min(0),    // pad
    ])
    .areas(inner);

    // Input rows
    render_input_row(
        frame,
        row_cid,
        "Cluster ID",
        &app.connection.cluster_id,
        focused_field == ConnectionField::ClusterId,
    );
    render_input_row(
        frame,
        row_addr,
        "Address   ",
        &app.connection.address,
        focused_field == ConnectionField::Address,
    );

    // Connect button
    let btn_focused = focused_field == ConnectionField::ConnectButton;
    let btn_style = if btn_focused {
        Style::default()
            .fg(THEME.bg)
            .bg(THEME.accent)
            .add_modifier(Modifier::BOLD)
    } else {
        Style::default().fg(THEME.text_muted)
    };
    let btn_label = if btn_focused {
        " ▶ Connect "
    } else {
        "   Connect  "
    };
    let btn = Paragraph::new(btn_label)
        .style(btn_style)
        .alignment(Alignment::Center);
    frame.render_widget(btn, row_btn);
}

fn render_input_row(frame: &mut Frame, area: Rect, label: &str, value: &str, focused: bool) {
    let [pad, lbl_area, val_area, pad2] = Layout::horizontal([
        Constraint::Length(2),
        Constraint::Length(12),
        Constraint::Min(4),
        Constraint::Length(2),
    ])
    .areas(area);

    // Blank pads
    let _ = (pad, pad2);

    let label_style = if focused {
        Style::default().fg(THEME.accent)
    } else {
        Style::default().fg(THEME.text_muted)
    };
    frame.render_widget(Paragraph::new(label).style(label_style), lbl_area);

    let val_style = if focused {
        Style::default().fg(THEME.text)
    } else {
        Style::default().fg(THEME.text_dim)
    };

    let display: Line = if focused {
        Line::from(vec![
            Span::styled(value, val_style),
            Span::styled("▌", Style::default().fg(THEME.accent)),
        ])
    } else {
        Line::from(Span::styled(value, val_style))
    };

    frame.render_widget(Paragraph::new(display), val_area);
}

fn center_h(area: Rect, width: u16) -> Rect {
    let clamped = width.min(area.width);
    let [_, center, _] = Layout::horizontal([
        Constraint::Fill(1),
        Constraint::Length(clamped),
        Constraint::Fill(1),
    ])
    .areas(area);
    center
}
