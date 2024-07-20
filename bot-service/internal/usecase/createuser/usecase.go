//go:generate minimock -g -s .go -o ../../../mocks/usecase/createUser
package createuser

import (
	"context"

	"github.com/google/uuid"

	"main/internal/domain"
)

type iUserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
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
