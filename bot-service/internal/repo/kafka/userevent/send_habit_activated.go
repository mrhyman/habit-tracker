package userevent

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/google/uuid"

	"main/internal/domain"
)

type HabitActivatedMessage struct {
	HabitId uuid.UUID `json:"habit_id"`
	UserID  uuid.UUID `json:"user_id"`
}

func (r Repo) SendHabitActivatedEvent(ctx context.Context, event domain.HabitActivatedEvent) error {
	msg := &sarama.ProducerMessage{
		Topic: r.topic,
		Key:   sarama.StringEncoder(event.HabitId.String()),
		Value: sarama.StringEncoder(fmt.Sprintf("%+v", event)),
	}
	p, o, err := r.pr.SendMessage(msg)
	if err != nil {
		slog.ErrorContext(
			ctx, "FAILED to send habit_created event", slog.String("err", err.Error()),
		)
		return err
	}
	slog.InfoContext(
		ctx,
		fmt.Sprintf("habit_activated event sent to partition %d at offset %d", p, o),
		slog.String("event_id", event.EventBase.ID()),
	)

	return nil
}
