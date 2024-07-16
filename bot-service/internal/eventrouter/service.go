package eventrouter

import (
	"context"

	"main/internal/domain"
)

//go:generate minimock -g -s .go -o ../../mocks/eventrouter
type iUserCreatedEventRepo interface {
	SendUserCreatedEvent(ctx context.Context, event domain.UserCreatedEvent) error
}

type iHabitActivatedEventRepo interface {
	SendHabitActivatedEvent(ctx context.Context, event domain.HabitActivatedEvent) error
}

// Service сервис маршрутизации доменных событий.
type Service struct {
	userCreatedEventRepo    iUserCreatedEventRepo
	habitActivatedEventRepo iHabitActivatedEventRepo
}

// NewService возвращает новый Service.
func NewService(userEvenRepo iUserCreatedEventRepo, habitActivatedEventRepo iHabitActivatedEventRepo) *Service {
	return &Service{
		userCreatedEventRepo:    userEvenRepo,
		habitActivatedEventRepo: habitActivatedEventRepo,
	}
}
