//go:generate minimock -g -s .go -o ../../mocks/handler/http
package handler

import (
	"context"
	"main/internal/usecase/activatehabit"

	"main/internal/domain"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
)

type iCreateUser interface {
	Handle(ctx context.Context, cmd createuser.Command) error
}

type iGetUser interface {
	Handle(ctx context.Context, q getuserbyid.Query) (*domain.User, error)
}

type iActivateHabit interface {
	Handle(ctx context.Context, cmd activatehabit.Command) error
}

type HttpHandler struct {
	CreateUserHandler    iCreateUser
	GetUserByIdHandler   iGetUser
	ActivateHabitHandler iActivateHabit
}

func New(
	createUserHandler iCreateUser,
	getUserByIdHandler iGetUser,
	activateHabitHandler iActivateHabit,
) *HttpHandler {
	return &HttpHandler{
		CreateUserHandler:    createUserHandler,
		GetUserByIdHandler:   getUserByIdHandler,
		ActivateHabitHandler: activateHabitHandler,
	}
}
