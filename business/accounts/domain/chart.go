// Package domain provides account and ledger domain types for the Terrace chart of accounts.
package domain

// LedgerAsset maps a TigerBeetle ledger ID to an asset symbol and display info.
type LedgerAsset struct {
	Symbol   string
	Name     string
	Decimals int
}

// Ledgers maps ledger IDs to asset info (Terrace ledger-v2).
var Ledgers = map[uint32]LedgerAsset{
	1:   {Symbol: "USD", Name: "US Dollar", Decimals: 2},
	2:   {Symbol: "EUR", Name: "Euro", Decimals: 2},
	100: {Symbol: "BTC", Name: "Bitcoin", Decimals: 8},
	101: {Symbol: "ETH", Name: "Ethereum", Decimals: 8},
	102: {Symbol: "USDC", Name: "USD Coin", Decimals: 6},
	103: {Symbol: "USDT", Name: "Tether", Decimals: 6},
	104: {Symbol: "SOL", Name: "Solana", Decimals: 9},
}

// LedgerSymbol returns the asset symbol for a ledger ID, or the ID as string.
func LedgerSymbol(id uint32) string {
	if a, ok := Ledgers[id]; ok {
		return a.Symbol
	}
	return ""
}

// AccountType maps an account code to a human-readable type name.
type AccountType struct {
	Code uint16
	Name string
}

// AccountTypes maps account codes to type names (Terrace ledger-v2).
var AccountTypes = map[uint16]string{
	1:   "USER_WALLET_DEFI",
	2:   "USER_ACCT_CEFI",
	100: "VENUE_BINANCE",
	101: "VENUE_OKX",
	102: "VENUE_BYBIT",
	103: "VENUE_BITGO",
	200: "HOLD_TRADE",
	201: "HOLD_WITHDRAWAL",
	202: "HOLD_SETTLEMENT",
	300: "FEES_COLLECTED",
	301: "FEES_VENUE",
	400: "SETTLEMENT_TRANSIT",
}

// AccountTypeName returns the type name for an account code, or "UNKNOWN".
func AccountTypeName(code uint16) string {
	if name, ok := AccountTypes[code]; ok {
		return name
	}
	return "UNKNOWN"
}

// TransferTypes maps transfer codes to type names (Terrace ledger-v2).
var TransferTypes = map[uint16]string{
	1:  "DEPOSIT",
	2:  "WITHDRAWAL",
	3:  "TRADE_BUY",
	4:  "TRADE_SELL",
	5:  "SETTLEMENT",
	6:  "FEE",
	7:  "REBATE",
	10: "INTERNAL_MOVE",
	11: "REBALANCE",
	12: "HOLD_RESERVE",
	13: "HOLD_RELEASE",
	14: "HOLD_TIMEOUT",
}

// TransferTypeName returns the type name for a transfer code, or "UNKNOWN".
func TransferTypeName(code uint16) string {
	if name, ok := TransferTypes[code]; ok {
		return name
	}
	return "UNKNOWN"
}

// Venues maps venue IDs (user_data_32) to venue names.
var Venues = map[uint32]string{
	0: "Internal",
	1: "Binance",
	2: "OKX",
	3: "Bybit",
	4: "BitGo",
	5: "Coinbase",
}

// VenueName returns the venue name for a user_data_32 value, or "Unknown".
func VenueName(id uint32) string {
	if name, ok := Venues[id]; ok {
		return name
	}
	return "Unknown"
}
