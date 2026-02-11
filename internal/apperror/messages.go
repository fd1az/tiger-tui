package apperror

var messages = map[Code]string{
	// General
	CodeRequiredField:      "Required field is missing",
	CodeInvalidInput:       "Invalid input provided",
	CodeInvalidFormat:      "Invalid data format",
	CodeInvalidState:       "Invalid state for this operation",
	CodeNotFound:           "Resource not found",
	CodeValidationError:    "Validation error",
	CodeConfigurationError: "Configuration error",
	CodeInternalError:      "Internal error",
	CodeUnknownError:       "An unknown error occurred",

	// TigerBeetle
	CodeTBConnectionFailed: "Failed to connect to TigerBeetle",
	CodeTBConnectionClosed: "TigerBeetle connection closed",
	CodeTBRequestFailed:    "TigerBeetle request failed",
	CodeTBTimeout:          "TigerBeetle request timed out",
	CodeTBInvalidCluster:   "Invalid TigerBeetle cluster ID",
	CodeTBInvalidAddress:   "Invalid TigerBeetle address",

	// Account/Transfer
	CodeAccountNotFound:      "Account not found",
	CodeAccountCreateFailed:  "Failed to create account",
	CodeTransferNotFound:     "Transfer not found",
	CodeTransferCreateFailed: "Failed to create transfer",
	CodeInvalidAccountID:     "Invalid account ID",
	CodeInvalidTransferID:    "Invalid transfer ID",
	CodeInvalidLedger:        "Invalid ledger ID",
	CodeInsufficientBalance:  "Insufficient balance",

	// Circuit breaker
	CodeCircuitOpen:     "Circuit breaker is open",
	CodeCircuitHalfOpen: "Circuit breaker is half-open",
}
