package userupdated

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log/slog"
	"main/internal/domain"
)

func (r Repo) SendUserUpdatedEvent(ctx context.Context, event domain.UserUpdatedEvent) error {
	msg := &sarama.ProducerMessage{
		Topic: r.topic,
		Key:   sarama.StringEncoder(event.UserID.String()),
		Value: sarama.StringEncoder(fmt.Sprintf("%+v", event)),
	}
	p, o, err := r.pr.SendMessage(msg)
	if err != nil {
		return err
	}
	slog.InfoContext(
		ctx,
		fmt.Sprintf("user_uppdated event sent to partition %d at offset %d", p, o),
		slog.String("event_id", event.EventBase.ID()),
	)

	return nil
}
