package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log/slog"
	"main/internal/config"
)

type Kafka struct {
	Ctx context.Context
	sarama.SyncProducer
}

func New(ctx context.Context, cfg config.KafkaConfig) (*Kafka, error) {
	producer, err := sarama.NewSyncProducer(
		[]string{fmt.Sprintf("%s:%d", cfg.GetHost(), cfg.GetPort())},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Kafka{ctx, producer}, nil
}

func (k *Kafka) Close() {
	slog.InfoContext(k.Ctx, "Closing kafka connection")
	_ = k.SyncProducer.Close()
}
