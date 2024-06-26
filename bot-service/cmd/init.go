package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"main/internal/config"
)

func initDefaultLogger() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

func initLogger(lc config.LoggerConfig) {
	opts := &slog.HandlerOptions{
		AddSource: lc.GetAddSource(),
		Level:     lc.GetLevel(),
	}
	info, _ := debug.ReadBuildInfo()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts)).With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", info.GoVersion),
		),
	)

	slog.SetDefault(logger)
}
