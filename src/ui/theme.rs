//! Mainframe Modern color palette and reusable style helpers.
//!
//! All color tokens are defined here as a single source of truth,
//! matching the design spec in `design.md`.

#![expect(
    dead_code,
    reason = "style helpers and tokens used progressively across phases"
)]

use ratatui::style::{Color, Modifier, Style};

/// Complete set of Mainframe Modern color tokens.
#[derive(Debug)]
pub struct Theme {
    /// Near-black background.
    pub bg: Color,
    /// Slightly lighter panel background.
    pub panel_bg: Color,
    /// Subtle border / divider.
    pub border: Color,
    /// Primary text.
    pub text: Color,
    /// Secondary text (hints, labels).
    pub text_muted: Color,
    /// Tertiary text (least prominent).
    pub text_dim: Color,
    /// `TigerBeetle` amber — focus, selection, highlights.
    pub accent: Color,
    /// Committed / posted / ok.
    pub success: Color,
    /// Pending / needs attention.
    pub warning: Color,
    /// Rejected / invariant broken.
    pub error: Color,
    /// Networking / state transitions.
    pub info: Color,
}

pub const THEME: Theme = Theme {
    bg: Color::Rgb(11, 15, 20),            // #0B0F14
    panel_bg: Color::Rgb(16, 24, 38),      // #101826
    border: Color::Rgb(36, 48, 65),        // #243041
    text: Color::Rgb(230, 237, 243),       // #E6EDF3
    text_muted: Color::Rgb(155, 167, 180), // #9BA7B4
    text_dim: Color::Rgb(107, 119, 133),   // #6B7785
    accent: Color::Rgb(244, 185, 66),      // #F4B942
    success: Color::Rgb(46, 204, 113),     // #2ECC71
    warning: Color::Rgb(243, 156, 18),     // #F39C12
    error: Color::Rgb(231, 76, 60),        // #E74C3C
    info: Color::Rgb(34, 193, 195),        // #22C1C3
};

// -- reusable style helpers -------------------------------------------------

pub fn style_text() -> Style {
    Style::default().fg(THEME.text)
}

pub fn style_muted() -> Style {
    Style::default().fg(THEME.text_muted)
}

pub fn style_dim() -> Style {
    Style::default().fg(THEME.text_dim)
}

pub fn style_accent() -> Style {
    Style::default().fg(THEME.accent)
}

pub fn style_accent_bold() -> Style {
    Style::default()
        .fg(THEME.accent)
        .add_modifier(Modifier::BOLD)
}

pub fn style_border() -> Style {
    Style::default().fg(THEME.border)
}

pub fn style_border_focused() -> Style {
    Style::default().fg(THEME.accent)
}
