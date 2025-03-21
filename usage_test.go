package log_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Trung-DV/log"
)

func TestLog(t *testing.T) {
	ctx := context.Background()
	log.Info(ctx, "This is info log")
	log.Debug(ctx, "This is debug log, but will not be printed")
	log.Warn(ctx, "Warning log")
	log.Error(ctx, "Error happen")

	log.Setup(log.DebugLevel)
	log.Debug(ctx, "Debug information")

	log.WithKeyValues("user_id", 123, "role", []string{"admin", "user"}).Info(ctx, "User logged in")
	log.WithKeyValues("user_id", 123, "role", []string{"admin", "user"}).Debug(ctx, "Debug User logged in")
	log.WithKeyValues("user_id", 123, "role", []string{"admin", "user"}).Warn(ctx, "Warning")
	log.WithKeyValues("user_id", 123).WithKeyValues("role", "admin").Error(ctx, "User logged in")

	ctx = log.WithContextValues(ctx, "context", "logged in")
	log.Info(ctx, "Info with", "context")

	log.Info(ctx, "Mixed", "string", 123, errors.New("error"))

	ctx = log.SaveRequestID(ctx, "main-12345")
	log.Info(ctx, "Processing with request ID")
	ctx = log.WithContextValues(ctx, "user", "main")
	log.Error(ctx, "Error processing request")
}
