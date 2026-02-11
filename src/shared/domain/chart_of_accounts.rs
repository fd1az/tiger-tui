//! Mappings from `TigerBeetle` numeric IDs to human-readable labels.
//!
//! These correspond to the Terrace ledger-v2 chart of accounts:
//! ledger IDs → assets, account codes → types, transfer codes → types,
//! and `user_data_32` → venue names.

#![allow(dead_code, reason = "lookup functions used in Phase 3 table views")]

/// Ledger ID → (asset ticker, decimal scale).
pub fn ledger_info(id: u32) -> (&'static str, u8) {
    match id {
        // Fiat (1–99)
        1 => ("USD", 2),
        2 => ("EUR", 2),
        3 => ("ARS", 2),
        4 => ("BRL", 2),
        // Crypto (100+)
        100 => ("BTC", 8),
        101 => ("ETH", 18),
        102 => ("USDC", 6),
        103 => ("USDT", 6),
        104 => ("SOL", 9),
        105 => ("LTC", 8),
        106 => ("DOGE", 8),
        107 => ("AVAX", 18),
        108 => ("MATIC", 18),
        _ => ("???", 2),
    }
}

/// Returns only the asset ticker for a given ledger ID.
pub fn ledger_name(id: u32) -> &'static str {
    ledger_info(id).0
}

/// Returns the decimal scale for a given ledger ID.
pub fn ledger_scale(id: u32) -> u8 {
    ledger_info(id).1
}

/// Account code → human-readable type name.
pub fn account_type_name(code: u16) -> &'static str {
    match code {
        // User accounts (1–3)
        1 => "USER_WALLET_DEFI",
        2 => "USER_ACCOUNT_CEFI",
        3 => "USER_ACCOUNT_TRADFI",
        // Venue omnibus (100–105)
        100 => "VENUE_BINANCE",
        101 => "VENUE_OKX",
        102 => "VENUE_BYBIT",
        103 => "VENUE_GATE",
        104 => "VENUE_B2C2",
        105 => "CUSTODY_BITGO",
        // Hold accounts (200–202)
        200 => "HOLD_TRADE",
        201 => "HOLD_WITHDRAWAL",
        202 => "HOLD_SETTLEMENT",
        // Fee accounts (300–301)
        300 => "FEES_COLLECTED",
        301 => "FEES_VENUE",
        // Settlement / adjustment (400–402)
        400 => "SETTLEMENT_IN_TRANSIT",
        401 => "ADJUSTMENT",
        402 => "EXTERNAL_FUNDING",
        // Liquidity (500+)
        500 => "LIQUIDITY_BRIDGE",
        _ => "UNKNOWN",
    }
}

/// Transfer code → human-readable type name.
pub fn transfer_type_name(code: u16) -> &'static str {
    match code {
        1 => "DEPOSIT",
        2 => "WITHDRAWAL",
        3 => "TRADE_BUY",
        4 => "TRADE_SELL",
        5 => "INTERNAL_TRANSFER",
        6 => "VENUE_SETTLEMENT",
        7 => "FEE_CHARGE",
        8 => "FEE_REBATE",
        9 => "ADJUSTMENT",
        10 => "FUNDING_ALLOCATION",
        11 => "INTER_VENUE_TRANSFER",
        12 => "CUSTODY_MOVEMENT",
        13 => "HOLD_RESERVE",
        14 => "SETTLE_EXTERNAL",
        _ => "UNKNOWN",
    }
}

/// `user_data_32` venue enum → venue name.
pub fn venue_name(id: u32) -> &'static str {
    match id {
        0 => "Internal",
        1 => "Binance",
        2 => "OKX",
        3 => "Bybit",
        4 => "Gate",
        5 => "B2C2",
        6 => "BitGo",
        10 => "OnChain",
        _ => "Unknown",
    }
}

/// Formats a raw `u128` amount using the ledger's decimal scale.
///
/// Returns a string like `"1,500.00"` or `"0.05000000"`.
pub fn format_amount(raw: u128, scale: u8) -> String {
    if scale == 0 {
        return format_with_commas(raw);
    }

    let divisor = 10u128.pow(u32::from(scale));
    let whole = raw / divisor;
    let frac = raw % divisor;

    format!(
        "{}.{:0>width$}",
        format_with_commas(whole),
        frac,
        width = scale as usize
    )
}

fn format_with_commas(n: u128) -> String {
    let s = n.to_string();
    let mut result = String::with_capacity(s.len() + s.len() / 3);
    for (i, c) in s.chars().enumerate() {
        if i > 0 && (s.len() - i).is_multiple_of(3) {
            result.push(',');
        }
        result.push(c);
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn ledger_mappings() {
        assert_eq!(ledger_name(1), "USD");
        assert_eq!(ledger_name(100), "BTC");
        assert_eq!(ledger_scale(100), 8);
        assert_eq!(ledger_name(999), "???");
    }

    #[test]
    fn account_type_mappings() {
        assert_eq!(account_type_name(1), "USER_WALLET_DEFI");
        assert_eq!(account_type_name(200), "HOLD_TRADE");
        assert_eq!(account_type_name(999), "UNKNOWN");
    }

    #[test]
    fn transfer_type_mappings() {
        assert_eq!(transfer_type_name(1), "DEPOSIT");
        assert_eq!(transfer_type_name(13), "HOLD_RESERVE");
        assert_eq!(transfer_type_name(999), "UNKNOWN");
    }

    #[test]
    fn venue_mappings() {
        assert_eq!(venue_name(0), "Internal");
        assert_eq!(venue_name(1), "Binance");
        assert_eq!(venue_name(99), "Unknown");
    }

    #[test]
    fn amount_formatting() {
        assert_eq!(format_amount(150_000, 2), "1,500.00");
        assert_eq!(format_amount(5_000_000, 8), "0.05000000");
        assert_eq!(format_amount(150_000_000, 8), "1.50000000");
        assert_eq!(format_amount(0, 2), "0.00");
        assert_eq!(format_amount(1_000_000_000, 2), "10,000,000.00");
    }
}
