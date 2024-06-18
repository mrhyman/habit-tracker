package config

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"
)

type LoggerConfig struct {
	AddSource bool       `mapstructure:"add_source"`
	Level     slog.Level `mapstructure:"level"`
}

func (lc LoggerConfig) GetAddSource() bool {
	return lc.AddSource
}

func (lc LoggerConfig) GetLevel() slog.Level {
	return lc.Level
}

func InitLogger(lc LoggerConfig) {
	opts := &slog.HandlerOptions{
		AddSource: lc.GetAddSource(),
		Level:     lc.GetLevel(),
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	info, _ := debug.ReadBuildInfo()
	logger.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		"runtime info",
		slog.String("go_version", info.GoVersion),
	)

	slog.SetDefault(logger)
}
