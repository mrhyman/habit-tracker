package config

import "time"

type Config struct {
	AppName         string         `mapstructure:"app_name"`
	Env             string         `mapstructure:"env"`
	Port            int            `mapstructure:"port"`
	Database        DatabaseConfig `mapstructure:"db"`
	GracefulTimeout time.Duration  `mapstructure:"graceful_timeout"`
}
