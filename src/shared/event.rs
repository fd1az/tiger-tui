//! Unified event loop that merges crossterm input with async results.

use std::sync::mpsc;
use std::thread;
use std::time::Duration;

use crossterm::event::{self, Event as CtEvent, KeyEvent, KeyEventKind};

/// Unified event type consumed by the main loop.
#[expect(dead_code, reason = "Resize variant used by crossterm polling thread")]
pub enum Event {
    /// Key press (releases are filtered out).
    Key(KeyEvent),
    /// Terminal resized.
    Resize(u16, u16),
    /// Tick (no user input within the poll window).
    Tick,
}

/// Polls crossterm in a background thread and forwards events via `mpsc`.
///
/// The sender can be cloned to inject `TbResult` events from the tokio
/// runtime in Phase 2.
#[derive(Debug)]
pub struct EventHandler {
    rx: mpsc::Receiver<Event>,
    _tx: mpsc::Sender<Event>,
}

impl EventHandler {
    pub fn new(tick_rate: Duration) -> Self {
        let (tx, rx) = mpsc::channel();
        let event_tx = tx.clone();

        thread::spawn(move || {
            loop {
                if event::poll(tick_rate).unwrap_or(false) {
                    match event::read() {
                        Ok(CtEvent::Key(key)) if key.kind == KeyEventKind::Press => {
                            if event_tx.send(Event::Key(key)).is_err() {
                                break;
                            }
                        }
                        Ok(CtEvent::Resize(w, h)) => {
                            if event_tx.send(Event::Resize(w, h)).is_err() {
                                break;
                            }
                        }
                        _ => {}
                    }
                } else if event_tx.send(Event::Tick).is_err() {
                    break;
                }
            }
        });

        Self { rx, _tx: tx }
    }

    /// Blocking receive — returns the next event.
    pub fn next(&self) -> anyhow::Result<Event> {
        Ok(self.rx.recv()?)
    }
}
