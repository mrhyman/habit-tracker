package handler

import (
	"context"
	"main/internal/database"
	"main/internal/database/repository"
)

type HttpHandler struct {
	Ctx            context.Context
	UserRepository *repository.UserRepositoryImpl
}

func New(ctx context.Context, db *database.DB) *HttpHandler {
	return &HttpHandler{ctx, repository.NewUserRepository(db.Pool)}
}
