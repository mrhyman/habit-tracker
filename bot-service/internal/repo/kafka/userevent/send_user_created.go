package userevent

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/google/uuid"

	"main/internal/domain"
)

type UserCreatedMessage struct {
	UserID uuid.UUID `json:"user_id"`
}

func (r Repo) SendUserCreatedEvent(ctx context.Context, event domain.UserCreatedEvent) error {
	// mapping from domain event to kafka message
	msg := &sarama.ProducerMessage{

		Topic: r.topic,                                     // topic from cfg
		Key:   sarama.StringEncoder(event.UserID.String()), // user id
		Value: sarama.StringEncoder("json"),                // marshal message
	}
	_, _, err := r.pr.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
