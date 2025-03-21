# log
A structured logging package for Go applications - zerolog wrapper

## Introduction
This package is a lightweight wrapper around [zerolog](https://github.com/rs/zerolog) designed to provide a more natural, user-friendly logging interface. While zerolog offers excellent performance and structured logging capabilities, this package simplifies its usage with intuitive methods and context-aware functions that feel more natural to use in Go applications.

The wrapper maintains zerolog's performance advantages while offering:

- Structured JSON logs for better parsing and analysis
- Context-aware logging with request ID tracking
- Multiple log levels (Debug, Info, Warn, Error)
- Key-value enrichment for detailed log context
- Caller information for easier debugging
- A simplified API that feels more natural to use

## Installation

```bash
go get github.com/Trung-DV/log
```

## Usage

### Basic Logging

```go
ctx := context.Background()

// Different log levels
log.Info(ctx, "This is info log")
log.Debug(ctx, "This is debug log") // Only prints if debug level is enabled
log.Warn(ctx, "Warning log")
log.Error(ctx, "Error happen")

// Enable debug level logging
log.Setup(log.DebugLevel)
```

### Mixed Arguments Logging

The logger supports multiple arguments of different types which will be formatted together:

```go
// Mix strings, numbers and errors in a single log message
log.Info(ctx, "Mixed", "string", 123, errors.New("error"))

// Output: {"level":"info","time":"...","caller":"...","message":"Mixed string 123 error"}
```

### Structured Logging with Key-Value Pairs

```go
// Add key-value pairs to your logs
log.WithKeyValues("user_id", 123, "role", []string{"admin", "user"}).Info(ctx, "User logged in")

// Chain multiple WithKeyValues calls
log.WithKeyValues("user_id", 123).WithKeyValues("role", "admin").Error(ctx, "User logged in")
```

### Context-Aware Logging

```go
// Add values to the context
ctx = log.WithContextValues(ctx, "context", "logged in")
log.Info(ctx, "Info with context")

// Track request IDs
ctx = log.SaveRequestID(ctx, "main-12345")
log.Info(ctx, "Processing with request ID")

// Add more context values
ctx = log.WithContextValues(ctx, "user", "main")
log.Error(ctx, "Error processing request")
```

### Log Levels

Available log levels in order of increasing severity:

```go
log.TraceLevel  // -1 (lowest level, most verbose)
log.DebugLevel  // 0
log.InfoLevel   // 1 (default)
log.WarnLevel   // 2
log.ErrorLevel  // 3
log.FatalLevel  // 4
log.PanicLevel  // 5
log.NoLevel     // 6
log.Disabled    // 7 (highest level, no logging)
```

> **Note:** Log levels are imported directly from the zerolog package, ensuring compatibility with zerolog's level system.

Set the log level:

```go
// Parse from string
level, err := log.ParseLevel("debug")
if err == nil {
    log.Setup(level)
}

// Or use constants
log.Setup(log.DebugLevel)
```

## Sample Output

```json
{"level":"info","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:13","message":"This is info log"}
{"level":"warn","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:15","message":"Warning log"}
{"level":"error","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:16","message":"Error happen"}
{"level":"debug","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:19","message":"Debug information"}
{"level":"info","user_id":123,"role":["admin","user"],"time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:21","message":"User logged in"}
{"level":"debug","user_id":123,"role":["admin","user"],"time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:22","message":"Debug User logged in"}
{"level":"warn","user_id":123,"role":["admin","user"],"time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:23","message":"Warning"}
{"level":"error","user_id":123,"role":"admin","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:24","message":"User logged in"}
{"level":"info","context":"logged in","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:27","message":"Info with context"}
{"level":"info","context":"logged in","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:29","message":"Mixed string 123 error"}
{"level":"info","context":"logged in","request_id":"main-12345","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:32","message":"Processing with request ID"}
{"level":"error","context":"logged in","user":"main","request_id":"main-12345","time":"2025-03-21T23:44:16+07:00","caller":"log/usage_test.go:34","message":"Error processing request"}
```
