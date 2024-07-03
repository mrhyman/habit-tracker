package config

import "time"

type Config struct {
	AppName                       string         `mapstructure:"app_name"`
	Env                           string         `mapstructure:"env"`
	Port                          int            `mapstructure:"port"`
	Database                      DatabaseConfig `mapstructure:"db"`
	GracefulTimeout               time.Duration  `mapstructure:"graceful_timeout"`
	BusinessMetricsScrapeInterval time.Duration  `mapstructure:"scrape_interval"`
	Logger                        LoggerConfig   `mapstructure:"logger"`
	Kafka                         KafkaConfig    `mapstructure:"kafka"`
	UserEventProducerConfig       ProducerConfig `mapstructure:"user_event_producer_config"`
}
