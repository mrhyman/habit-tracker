package usercreated

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/ds248a/closer"
	"log/slog"
	"main/internal/config"
)

type Repo struct {
	topic string
	pr    sarama.SyncProducer
}

func NewRepo(ctx context.Context, addr []string, config config.ProducerConfig) (*Repo, error) {
	producer, err := sarama.NewSyncProducer(addr, configProducer(config))
	if err != nil {
		return nil, err
	}
	closer.Add(func() {
		if err = producer.Close(); err != nil {
			slog.ErrorContext(ctx, "error producer close", slog.String("err", err.Error()))
		}
	})

	return &Repo{
		topic: config.Topic,
		pr:    producer,
	}, nil
}

func configProducer(config config.ProducerConfig) *sarama.Config {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = config.RequiredAcks
	saramaConfig.Producer.Compression = config.CompressionCodec
	saramaConfig.Producer.Retry.Max = config.Retry.Max
	saramaConfig.Producer.Retry.Backoff = config.Retry.Backoff
	saramaConfig.Producer.Timeout = config.RequestTimeout
	saramaConfig.Producer.Idempotent = config.Idempotent
	saramaConfig.Producer.Return.Successes = config.ReturnSuccesses

	saramaConfig.Net.MaxOpenRequests = config.MaxInFlightRequestsPerConnection

	return saramaConfig
}
