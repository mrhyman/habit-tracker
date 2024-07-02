package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"log/slog"
	"main/internal/domain"
)

type UserEventBus interface {
	UserCreated(user *domain.User) error
}

func NewUserEventBus(ctx context.Context, p sarama.SyncProducer) *UserEventBusImpl {
	return &UserEventBusImpl{ctx: ctx, producer: p}
}

type UserEventBusImpl struct {
	producer sarama.SyncProducer
	ctx      context.Context
}

// todo: команда?
func (b *UserEventBusImpl) UserCreated(user *domain.User) error {
	event, err := json.Marshal(eventFromDomain(user))
	if err != nil {
		slog.ErrorContext(b.ctx, "error parsing event", err.Error())
		return err
	}
	msg := &sarama.ProducerMessage{Topic: "user_created", Value: sarama.StringEncoder(event)}
	partition, offset, err := b.producer.SendMessage(msg)
	if err != nil {
		slog.ErrorContext(b.ctx, "FAILED to send user_created event", err.Error())
		return err
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
	slog.InfoContext(
		b.ctx,
		fmt.Sprintf("user_created event sent to partition %d at offset %d", partition, offset),
		event,
	)
	return nil
}
