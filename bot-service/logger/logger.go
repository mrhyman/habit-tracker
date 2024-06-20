package logger

import (
	"log/slog"
	"main/internal/config"
	"os"
	"runtime/debug"
)

func InitLogger(lc config.LoggerConfig) {
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
