package config

import "time"

type Config struct {
	AppName                           string         `mapstructure:"app_name"`
	Env                               string         `mapstructure:"env"`
	Port                              int            `mapstructure:"port"`
	Database                          DatabaseConfig `mapstructure:"db"`
	GracefulTimeout                   time.Duration  `mapstructure:"graceful_timeout"`
	BusinessMetricsScrapeInterval     time.Duration  `mapstructure:"scrape_interval"`
	Logger                            LoggerConfig   `mapstructure:"pkg"`
	Kafka                             KafkaConfig    `mapstructure:"eventbus"`
	UserCreatedEventProducerConfig    ProducerConfig `mapstructure:"user_created_event_producer_config"`
	HabitActivatedEventProducerConfig ProducerConfig `mapstructure:"habit_activated_event_producer_config"`
}
