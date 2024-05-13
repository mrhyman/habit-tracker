package config

import "time"

type DatabaseConfig struct {
	Connection string        `mapstructure:"connection"`
	DbName     string        `mapstructure:"dbName"`
	UserName   string        `mapstructure:"user_name"`
	Timeout    time.Duration `mapstructure:"timeout"`
}

func (dc DatabaseConfig) GetConnection() string {
	return dc.Connection
}

func (dc DatabaseConfig) GetDbName() string {
	return dc.DbName
}

func (dc DatabaseConfig) GetUserName() string {
	return dc.UserName
}

func (dc DatabaseConfig) GetTimeout() time.Duration {
	return dc.Timeout
}
