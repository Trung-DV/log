package log

import (
	"context"

	"github.com/rs/zerolog"
)

// kvLog is a wrapper around zerolog.Logger that enables method chaining for structured logging
type kvLog []any

// WithKeyValues adds key-value pairs to the logger's context and returns the updated logger
// for method chaining
func (l kvLog) WithKeyValues(keysAndValues ...any) kvLog {
	return append(l, keysAndValues...)
}

// Info logs a message at INFO level with the stored key-value pairs
func (l kvLog) Info(ctx context.Context, args ...any) {
	//logger := getLogContextFromContext(ctx).Logger().With().Fields(l).Logger()
	zerolog.Ctx(ctx).Info().Fields([]any(l)).Msg(formatMessage(args...))
}

// Debug logs a message at DEBUG level with the stored key-value pairs
func (l kvLog) Debug(ctx context.Context, args ...any) {
	//logger := getLogContextFromContext(ctx).Logger().With().Fields(l).Logger()
	zerolog.Ctx(ctx).Debug().Fields([]any(l)).Msg(formatMessage(args...))
}

// Warn logs a message at WARN level with the stored key-value pairs
func (l kvLog) Warn(ctx context.Context, args ...any) {
	//logger := getLogContextFromContext(ctx).Logger().With().Fields(l).Logger()
	zerolog.Ctx(ctx).Warn().Fields([]any(l)).Msg(formatMessage(args...))
}

// Error logs a message at ERROR level with the stored key-value pairs
func (l kvLog) Error(ctx context.Context, args ...any) {
	//logger := getLogContextFromContext(ctx).Logger().With().Fields(l).Logger()
	zerolog.Ctx(ctx).Error().Fields([]any(l)).Msg(formatMessage(args...))
}

// WithKeyValues creates a new kvLog with the provided key-value pairs
// for structured logging
func WithKeyValues(keysAndValues ...any) kvLog {
	return keysAndValues
}
