package userevent

import (
	"context"
	"log/slog"
	"os"

	"github.com/IBM/sarama"
	"github.com/ds248a/closer"

	"main/internal/config"
)

type Repo struct {
	topic string
	pr    sarama.SyncProducer
}

func NewRepo(ctx context.Context, addr []string, config config.ProducerConfig) (*Repo, error) {
	// Kafka producer
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Compression = sarama.CompressionSnappy
	saramaConfig.Producer.Timeout = config.Timeout

	producer, err := sarama.NewSyncProducer(addr, saramaConfig)
	if err != nil {
		return nil, err
	}
	closer.Add(func() {
		if err = producer.Close(); err != nil {
			slog.ErrorContext(ctx, "error close kafka producer", err)
			os.Exit(1)
		}
	})

	return &Repo{
		topic: config.Topic,
		pr:    producer,
	}, nil
}
