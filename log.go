package log

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

// Context keys
type requestIDCtxKey struct{}

type logCtxKey struct{}

var defaultLogCtx zerolog.Context

func init() {
	// Skip additional frames to show the actual caller in logs
	zerolog.CallerSkipFrameCount = 3

	// Set custom format for caller path: dir/file.go:line
	// Match the correct function signature with the pc parameter
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		// Find the last two slashes
		last := -1
		secondLastSlash := len(file) - 1

		for ; secondLastSlash >= 0; secondLastSlash-- {
			if file[secondLastSlash] != '/' {
				continue
			}
			if last != -1 {
				break
			}
			last = secondLastSlash
		}

		return file[secondLastSlash+1:] + ":" + strconv.Itoa(line)
	}

	defaultLogger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(zerolog.InfoLevel)

	defaultLogCtx = defaultLogger.With()
}

type Level = zerolog.Level

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
	// Values less than TraceLevel are handled as numbers.
)

func ParseLevel(levelStr string) (Level, error) {
	return zerolog.ParseLevel(levelStr)
}

// Setup allows customization of the kvLog output and leve
// Setup allows customization of the kvLog output and level
func Setup(level Level) {
	defaultLogCtx = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(level).
		With()
}

// Info logs an info level message with variadic arguments
func Info(ctx context.Context, args ...interface{}) {
	logger := getLogCtxWithRequestID(ctx).Logger()
	logger.Info().Msg(formatMessage(args...))
}

func getLogCtxWithRequestID(ctx context.Context) zerolog.Context {
	logCtx := getLogContextFromContext(ctx)

	// Add request ID if present
	if requestID, ok := ctx.Value(requestIDCtxKey{}).(string); ok && requestID != "" {
		logCtx = logCtx.Str("request_id", requestID)
	}
	return logCtx
}

// Debug logs a debug level message with variadic arguments
func Debug(ctx context.Context, args ...interface{}) {
	logger := getLogCtxWithRequestID(ctx).Logger()
	logger.Debug().Msg(formatMessage(args...))
}

// Warn logs a warn level message with variadic arguments
func Warn(ctx context.Context, args ...interface{}) {
	logger := getLogCtxWithRequestID(ctx).Logger()
	logger.Warn().Msg(formatMessage(args...))
}

// Error logs an error level message with variadic arguments
func Error(ctx context.Context, args ...interface{}) {
	logger := getLogCtxWithRequestID(ctx).Logger()
	logger.Error().Msg(formatMessage(args...))
}

// SaveRequestID stores a request ID in the context
func SaveRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDCtxKey{}, requestID)
}

// GetRequestID retrieves the request ID from the context
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if requestID, ok := ctx.Value(requestIDCtxKey{}).(string); ok {
		return requestID
	}

	return ""
}

// WithContextValues adds key-value pairs to the context for logging
func WithContextValues(ctx context.Context, keysAndValues ...interface{}) context.Context {
	// Get existing zerolog context from context or use default
	logCtx := getLogContextFromContext(ctx)

	// Add new key-value pairs
	logCtx = logCtx.Fields(keysAndValues)

	// Store the context directly
	return context.WithValue(ctx, logCtxKey{}, logCtx)
}

// Helper functions

// getLogContextFromContext retrieves a zerolog.Context from context or creates a new one
func getLogContextFromContext(ctx context.Context) zerolog.Context {
	if ctx == nil {
		return defaultLogCtx
	}

	// Check if we already have a zerolog.Context stored
	if logCtx, ok := ctx.Value(logCtxKey{}).(zerolog.Context); ok {
		return logCtx
	}

	return defaultLogCtx
}

// Helper function to format variadic arguments into a single message
func formatMessage(args ...interface{}) string {
	if len(args) == 0 {
		return ""
	}

	result := fmt.Sprintln(args...)
	// Trim the trailing newline that fmt.Sprintln adds
	return result[:len(result)-1]
}
