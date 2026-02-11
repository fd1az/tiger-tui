package apperror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// AppError implements the error interface and provides structured error handling.
type AppError struct {
	Code      Code      `json:"code"`
	Message   string    `json:"message"`
	Context   string    `json:"context,omitempty"`
	TraceID   string    `json:"traceId,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	cause     error
	stack     []uintptr
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if e.Context != "" {
		return fmt.Sprintf("%s: %s (code: %s, context: %s)", e.Code, e.Message, e.Code, e.Context)
	}
	return fmt.Sprintf("%s: %s (code: %s)", e.Code, e.Message, e.Code)
}

// Unwrap implements the errors.Unwrap interface.
func (e *AppError) Unwrap() error {
	return e.cause
}

// Is implements errors.Is for error comparison.
func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// WithTraceID sets the trace ID.
func (e *AppError) WithTraceID(traceID string) *AppError {
	e.TraceID = traceID
	return e
}

// ToLog serializes the error for structured logging.
func (e *AppError) ToLog() map[string]interface{} {
	log := map[string]interface{}{
		"code":      e.Code,
		"message":   e.Message,
		"timestamp": e.Timestamp.Format(time.RFC3339),
	}

	if e.Context != "" {
		log["context"] = e.Context
	}
	if e.TraceID != "" {
		log["traceId"] = e.TraceID
	}
	if e.cause != nil {
		log["cause"] = e.cause.Error()
	}
	if len(e.stack) > 0 {
		log["stack"] = e.formatStack()
	}

	return log
}

func (e *AppError) formatStack() string {
	var sb strings.Builder
	frames := runtime.CallersFrames(e.stack)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			sb.WriteString(fmt.Sprintf("\n\t%s:%d %s", frame.File, frame.Line, frame.Function))
		}
		if !more {
			break
		}
	}
	return sb.String()
}

func captureStack() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}

// New creates a new AppError with the given code and options.
func New(code Code, opts ...Option) *AppError {
	err := &AppError{
		Code:      code,
		Message:   messages[code],
		Timestamp: time.Now(),
		stack:     captureStack(),
	}

	for _, opt := range opts {
		opt(err)
	}

	if err.Message == "" {
		err.Message = string(code)
	}

	return err
}

// Option is a functional option for AppError.
type Option func(*AppError)

// WithMessage sets a custom message.
func WithMessage(message string) Option {
	return func(e *AppError) {
		e.Message = message
	}
}

// WithContext adds context information.
func WithContext(context string) Option {
	return func(e *AppError) {
		e.Context = context
	}
}

// WithCause wraps an underlying error.
func WithCause(cause error) Option {
	return func(e *AppError) {
		e.cause = cause
	}
}

// Internal creates an internal error.
func Internal(code Code, context string, cause error) *AppError {
	return New(code, WithContext(context), WithCause(cause))
}

// Wrap wraps a standard error into AppError.
func Wrap(err error, code Code, context string) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		if context != "" && appErr.Context == "" {
			appErr.Context = context
		}
		return appErr
	}

	return Internal(code, context, err)
}

// IsAppError checks if an error is an AppError.
func IsAppError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr)
}

// GetCode extracts the error code from an error.
func GetCode(err error) Code {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code
	}
	return CodeUnknownError
}
