package config

type Config struct {
	AppName  string         `mapstructure:"app_name" validate:"required"`
	Env      string         `mapstructure:"env" env-default:"local" validate:"required"`
	Database DatabaseConfig `mapstructure:"db" validate:"required"`
}
