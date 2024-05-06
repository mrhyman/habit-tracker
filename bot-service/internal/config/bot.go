package config

import "time"

type BotConfig struct {
	Token         string        `mapstructure:"token"`
	PollerTimeout time.Duration `mapstructure:"poller_timeout"`
}

func (bc BotConfig) GetToken() string {
	return bc.Token
}

func (bc BotConfig) GetPollerTimeout() time.Duration {
	return bc.PollerTimeout
}
