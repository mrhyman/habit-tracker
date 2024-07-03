//go:generate minimock -g -s .go -o ../../mocks/handler/http
package handler

import (
	"context"

	"main/internal/domain"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
)

type iCreateUser interface {
	Handle(ctx context.Context, cmd createuser.Command) error
}

type iGetUser interface {
	Handle(q getuserbyid.Query) (*domain.User, error)
}

type HttpHandler struct {
	CreateUserHandler  iCreateUser
	GetUserByIdHandler iGetUser
}

func New(
	createUserHandler iCreateUser,
	getUserByIdHandler iGetUser,
) *HttpHandler {
	return &HttpHandler{
		CreateUserHandler:  createUserHandler,
		GetUserByIdHandler: getUserByIdHandler,
	}
}
