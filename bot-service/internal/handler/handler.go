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
	Handle(ctx context.Context, q getuserbyid.Query) (*domain.User, error)
}

type HttpHandler struct {
	Ctx                context.Context
	CreateUserHandler  iCreateUser
	GetUserByIdHandler iGetUser
}

func New(
	ctx context.Context,
	createUserHandler iCreateUser,
	getUserByIdHandler iGetUser,
) *HttpHandler {
	return &HttpHandler{
		Ctx:                ctx,
		CreateUserHandler:  createUserHandler,
		GetUserByIdHandler: getUserByIdHandler,
	}
}
