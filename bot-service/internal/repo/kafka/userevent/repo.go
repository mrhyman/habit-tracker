package userevent

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/ds248a/closer"
	"main/utils"

	"main/internal/config"
)

type Repo struct {
	topic string
	pr    sarama.SyncProducer
}

func NewRepo(ctx context.Context, addr []string, config config.ProducerConfig) (*Repo, error) {
	// Kafka producer
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = config.RequiredAcks
	saramaConfig.Producer.Compression = config.CompressionCodec
	saramaConfig.Producer.Retry.Max = config.Retry.Max
	saramaConfig.Producer.Retry.Backoff = config.Retry.Backoff
	saramaConfig.Producer.Timeout = config.RequestTimeout
	saramaConfig.Producer.Idempotent = config.Idempotent
	saramaConfig.Producer.Return.Successes = config.ReturnSuccesses

	saramaConfig.Net.MaxOpenRequests = config.MaxInFlightRequestsPerConnection

	producer, err := sarama.NewSyncProducer(addr, saramaConfig)
	if err != nil {
		return nil, err
	}
	closer.Add(func() {
		if err = producer.Close(); err != nil {
			utils.LogFatal(ctx, "error close kafka producer", err)
		}
	})

	return &Repo{
		topic: config.Topic,
		pr:    producer,
	}, nil
}
