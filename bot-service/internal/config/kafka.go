package config

import (
	"github.com/IBM/sarama"
	"time"
)

type KafkaConfig struct {
	Host []string `mapstructure:"host"`
}

func (kc KafkaConfig) GetHost() []string {
	return kc.Host
}

type ProducerRetry struct {
	Max     int           `mapstructure:"max"`
	Backoff time.Duration `mapstructure:"backoff"`
}

type ProducerConfig struct {
	Topic                            string                  `mapstructure:"topic"`
	RequiredAcks                     sarama.RequiredAcks     `mapstructure:"required_acks"`
	CompressionCodec                 sarama.CompressionCodec `mapstructure:"compression_codec"`
	Retry                            ProducerRetry           `mapstructure:"retry"`
	RequestTimeout                   time.Duration           `mapstructure:"request_timeout"`
	Idempotent                       bool                    `mapstructure:"idempotent"`
	ReturnSuccesses                  bool                    `mapstructure:"return_successes"`
	MaxInFlightRequestsPerConnection int                     `mapstructure:"max_in_flight_requests_per_connection"`
}
