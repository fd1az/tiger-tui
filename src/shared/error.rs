//! Typed application errors for domain and infrastructure boundaries.

use thiserror::Error;

#[expect(dead_code, reason = "used in Phase 2 when TB client is wired")]
#[derive(Debug, Error)]
pub enum AppError {
    #[error("Connection error: {0}")]
    Connection(String),

    #[error("TigerBeetle error: {0}")]
    TigerBeetle(String),

    #[error("Validation error: {0}")]
    Validation(String),

    #[error("Internal error: {0}")]
    Internal(String),
}
