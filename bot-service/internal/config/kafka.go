package config

import "time"

type KafkaConfig struct {
	Host []string `mapstructure:"host"`
}

func (kc KafkaConfig) GetHost() []string {
	return kc.Host
}

type ProducerConfig struct {
	Topic   string        `mapstructure:"topic"`
	Timeout time.Duration `mapstructure:"timeout"`
}
