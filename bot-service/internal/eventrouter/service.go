package eventrouter

import (
	"context"

	"main/internal/domain"
)

//go:generate minimock -g -s .go -o ../../mocks/eventrouter
type iUserEventRepo interface {
	SendUserCreatedEvent(ctx context.Context, event domain.UserCreatedEvent) error
}

// Service сервис маршрутизации доменных событий.
type Service struct {
	userEventRepo iUserEventRepo
}

// NewService возвращает новый Service.
func NewService(userEvenRepo iUserEventRepo) *Service {
	return &Service{
		userEventRepo: userEvenRepo,
	}
}
