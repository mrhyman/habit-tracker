package userevent

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/google/uuid"

	"main/internal/domain"
)

type UserCreatedMessage struct {
	UserID uuid.UUID `json:"user_id"`
}

func (r Repo) SendUserCreatedEvent(ctx context.Context, event domain.UserCreatedEvent) error {
	msg := &sarama.ProducerMessage{
		Topic: r.topic,                                     // topic from cfg
		Key:   sarama.StringEncoder(event.UserID.String()), // user id
		Value: sarama.StringEncoder("json"),                // marshal message
	}
	p, o, err := r.pr.SendMessage(msg)
	if err != nil {
		slog.ErrorContext(
			ctx, "FAILED to send user_created event", slog.String("err", err.Error()),
		)
		return err
	}
	slog.InfoContext(
		ctx,
		fmt.Sprintf("user_created event sent to partition %d at offset %d", p, o),
		slog.String("event_id", event.EventBase.ID()),
	)

	return nil
}
