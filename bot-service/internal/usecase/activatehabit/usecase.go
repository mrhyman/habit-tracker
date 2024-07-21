//go:generate minimock -g -s .go -o ../../../mocks/usecase/activateHabit
package activatehabit

import (
	"context"

	"github.com/google/uuid"

	"main/internal/domain"
)

type iUserRepo interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	ActivateHabit(ctx context.Context, userID uuid.UUID, habitID uuid.UUID) error
}

type iEventRepo interface {
	CreateEvent(ctx context.Context, ev *domain.Event) error
}

type iEventRouter interface {
	RouteAllEvents(ctx context.Context, events []domain.Event) error
}

type CommandHandler struct {
	userRepo    iUserRepo
	eventRepo   iEventRepo
	eventRouter iEventRouter
}

func NewCommandHandler(userRepo iUserRepo, eventRepo iEventRepo, eventRouter iEventRouter) *CommandHandler {
	return &CommandHandler{userRepo: userRepo, eventRepo: eventRepo, eventRouter: eventRouter}
}
