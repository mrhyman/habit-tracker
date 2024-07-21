package usercreated

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"log/slog"
	"main/internal/domain"
)

type Message struct {
	UserID uuid.UUID `json:"user_id"`
}

func (r Repo) SendUserCreatedEvent(ctx context.Context, event domain.UserCreatedEvent) error {
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
		fmt.Sprintf("user_created event sent to partition %d at offset %d", p, o),
		slog.String("event_id", event.EventBase.ID()),
	)

	return nil
}
