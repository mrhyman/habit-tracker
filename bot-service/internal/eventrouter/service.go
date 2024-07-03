package eventrouter

import (
	"context"

	"main/internal/domain"
)

//go:generate minimock -g -s .go -i iModerationTaskEventRepo -o ../../mocks/eventrouter
type iUserEventRepo interface {
	SendUserCreatedEvent(ctx context.Context, event domain.UserCreatedEvent) error
}

// Service сервис маршрутизации доменных событий.
type Service struct {
	userEvenRepo iUserEventRepo
}

// New возвращает новый Service.
func New(userEvenRepo iUserEventRepo) *Service {
	return &Service{
		userEvenRepo: userEvenRepo,
	}
}
