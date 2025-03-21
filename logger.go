package log

import (
	"context"
)

// kvLog is a wrapper around zerolog.Logger that enables method chaining for structured logging
type kvLog struct {
	fields []any
}

// WithKeyValues adds key-value pairs to the logger's context and returns the updated logger
// for method chaining
func (l kvLog) WithKeyValues(keysAndValues ...any) kvLog {
	l.fields = append(l.fields, keysAndValues...)
	return l
}

// Info logs a message at INFO level with the stored key-value pairs
func (l kvLog) Info(ctx context.Context, args ...interface{}) {
	logger := getLogContextFromContext(ctx).Logger().With().Fields(l.fields).Logger()
	logger.Info().Msg(formatMessage(args...))
}

// Debug logs a message at DEBUG level with the stored key-value pairs
func (l kvLog) Debug(ctx context.Context, args ...interface{}) {
	logger := getLogContextFromContext(ctx).Logger().With().Fields(l.fields).Logger()
	logger.Debug().Msg(formatMessage(args...))
}

// Warn logs a message at WARN level with the stored key-value pairs
func (l kvLog) Warn(ctx context.Context, args ...interface{}) {
	logger := getLogContextFromContext(ctx).Logger().With().Fields(l.fields).Logger()
	logger.Warn().Msg(formatMessage(args...))
}

// Error logs a message at ERROR level with the stored key-value pairs
func (l kvLog) Error(ctx context.Context, args ...interface{}) {
	logger := getLogContextFromContext(ctx).Logger().With().Fields(l.fields).Logger()
	logger.Error().Msg(formatMessage(args...))
}

// WithKeyValues creates a new kvLog with the provided key-value pairs
// for structured logging
func WithKeyValues(keysAndValues ...interface{}) kvLog {
	return kvLog{fields: keysAndValues}
}
