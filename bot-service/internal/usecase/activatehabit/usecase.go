//go:generate minimock -g -s .go -o ../../../mocks/usecase/activateHabit
package activatehabit

import (
	"context"

	"github.com/google/uuid"

	"main/internal/domain"
)

type iUserRepo interface {
	GetUserByID(userID uuid.UUID) (*domain.User, error)
	ActivateHabit(userID uuid.UUID, habitID uuid.UUID) error
}

type iEventRouter interface {
	RouteAllEvents(ctx context.Context, events []domain.Event) error
}

type CommandHandler struct {
	userRepo    iUserRepo
	eventRouter iEventRouter
}

func NewCommandHandler(userRepo iUserRepo, eventRouter iEventRouter) *CommandHandler {
	return &CommandHandler{userRepo: userRepo, eventRouter: eventRouter}
}
