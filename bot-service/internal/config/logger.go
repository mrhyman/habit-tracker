package config

import (
	"log/slog"
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
