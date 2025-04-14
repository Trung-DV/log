package log

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

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

	Setup(InfoLevel)

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
func Setup(level Level) {
	defaultLogger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(level)
	zerolog.DefaultContextLogger = &defaultLogger
}

// Info logs an info level message with variadic arguments
func Info(ctx context.Context, args ...interface{}) {
	zerolog.Ctx(ctx).Info().Msg(formatMessage(args...))
}

// Debug logs a debug level message with variadic arguments
func Debug(ctx context.Context, args ...interface{}) {
	zerolog.Ctx(ctx).Debug().Msg(formatMessage(args...))
}

// Warn logs a warn level message with variadic arguments
func Warn(ctx context.Context, args ...interface{}) {
	zerolog.Ctx(ctx).Warn().Msg(formatMessage(args...))
}

// Error logs an error level message with variadic arguments
func Error(ctx context.Context, args ...interface{}) {
	zerolog.Ctx(ctx).Error().Msg(formatMessage(args...))
}

// WithRequestID stores a request ID in a new zerolog context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return zerolog.Ctx(ctx).With(). // Create a child logger to hold request_id
					Str("request_id", requestID).
					Logger().
					WithContext(ctx) // New context with the child logger that has the request ID
}

// WithContextValues adds key-value pairs to the context for logging
func WithContextValues(ctx context.Context, keysAndValues ...interface{}) context.Context {
	return zerolog.Ctx(ctx).With(). // Create a child logger to hold key-value pairs
					Fields(keysAndValues).
					Logger().
					WithContext(ctx) // New context with the child logger
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
