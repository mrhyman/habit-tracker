package config

import "time"

type Config struct {
	AppName         string         `mapstructure:"app_name"`
	Env             string         `mapstructure:"env"`
	Database        DatabaseConfig `mapstructure:"db"`
	GracefulTimeout time.Duration  `mapstructure:"graceful_timeout"`
	Bot             BotConfig      `mapstructure:"bot"`
}
