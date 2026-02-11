package apperror

// Code represents a unique error code for the application.
type Code string

// General error codes.
const (
	CodeRequiredField      Code = "REQUIRED_FIELD"
	CodeInvalidInput       Code = "INVALID_INPUT"
	CodeInvalidFormat      Code = "INVALID_FORMAT"
	CodeInvalidState       Code = "INVALID_STATE"
	CodeNotFound           Code = "NOT_FOUND"
	CodeValidationError    Code = "VALIDATION_ERROR"
	CodeConfigurationError Code = "CONFIGURATION_ERROR"
	CodeInternalError      Code = "INTERNAL_ERROR"
	CodeUnknownError       Code = "UNKNOWN_ERROR"
)

// TigerBeetle error codes.
const (
	CodeTBConnectionFailed Code = "TB_CONNECTION_FAILED"
	CodeTBConnectionClosed Code = "TB_CONNECTION_CLOSED"
	CodeTBRequestFailed    Code = "TB_REQUEST_FAILED"
	CodeTBTimeout          Code = "TB_TIMEOUT"
	CodeTBInvalidCluster   Code = "TB_INVALID_CLUSTER"
	CodeTBInvalidAddress   Code = "TB_INVALID_ADDRESS"
)

// Account/Transfer error codes.
const (
	CodeAccountNotFound       Code = "ACCOUNT_NOT_FOUND"
	CodeAccountCreateFailed   Code = "ACCOUNT_CREATE_FAILED"
	CodeTransferNotFound      Code = "TRANSFER_NOT_FOUND"
	CodeTransferCreateFailed  Code = "TRANSFER_CREATE_FAILED"
	CodeInvalidAccountID      Code = "INVALID_ACCOUNT_ID"
	CodeInvalidTransferID     Code = "INVALID_TRANSFER_ID"
	CodeInvalidLedger         Code = "INVALID_LEDGER"
	CodeInsufficientBalance   Code = "INSUFFICIENT_BALANCE"
)

// Circuit breaker error codes.
const (
	CodeCircuitOpen     Code = "CIRCUIT_OPEN"
	CodeCircuitHalfOpen Code = "CIRCUIT_HALF_OPEN"
)
