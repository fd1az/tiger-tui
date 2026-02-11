mod modules;
mod shared;
mod ui;

use std::io;
use std::time::Duration;

use crossterm::{
    execute,
    terminal::{EnterAlternateScreen, LeaveAlternateScreen, disable_raw_mode, enable_raw_mode},
};
use ratatui::{Terminal, backend::CrosstermBackend};

use shared::app_state::AppState;
use shared::event::{Event, EventHandler};
use shared::message::{handle_key, update};
use ui::view;

fn main() -> anyhow::Result<()> {
    // -- Tracing (to file so it doesn't interfere with TUI) -----------------
    let log_file = std::fs::File::create("tiger-tui.log")?;
    tracing_subscriber::fmt()
        .with_writer(log_file)
        .with_ansi(false)
        .with_env_filter(
            tracing_subscriber::EnvFilter::from_default_env()
                .add_directive("tiger_tui=info".parse()?),
        )
        .init();

    tracing::info!("tiger-tui starting");

    // -- Terminal setup -----------------------------------------------------
    enable_raw_mode()?;
    let mut stdout = io::stdout();
    execute!(stdout, EnterAlternateScreen)?;
    let backend = CrosstermBackend::new(stdout);
    let mut terminal = Terminal::new(backend)?;

    // Restore terminal on panic
    let original_hook = std::panic::take_hook();
    std::panic::set_hook(Box::new(move |info| {
        let _ = disable_raw_mode();
        let _ = execute!(io::stderr(), LeaveAlternateScreen);
        original_hook(info);
    }));

    // -- App state + event handler ------------------------------------------
    let mut app = AppState::default();
    let events = EventHandler::new(Duration::from_millis(100));

    // -- Main loop (TEA) ----------------------------------------------------
    loop {
        terminal.draw(|frame| view(&app, frame))?;

        match events.next()? {
            Event::Key(key) => {
                if let Some(msg) = handle_key(key, &app) {
                    update(&mut app, &msg);
                }
            }
            Event::Tick | Event::Resize(_, _) => {}
        }

        if app.should_quit {
            break;
        }
    }

    // -- Cleanup ------------------------------------------------------------
    disable_raw_mode()?;
    execute!(terminal.backend_mut(), LeaveAlternateScreen)?;
    terminal.show_cursor()?;
    tracing::info!("tiger-tui exiting");

    Ok(())
}
