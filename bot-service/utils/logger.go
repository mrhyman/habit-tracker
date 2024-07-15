package utils

import (
	"context"
	"log/slog"
	"os"
)

func LogFatal(ctx context.Context, message string, err error) {
	slog.ErrorContext(ctx, message, slog.String("err", err.Error()))
	os.Exit(1)
}
