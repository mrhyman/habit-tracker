package user_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"main/internal/database"
	"main/internal/database/repository"
	"net/http"
)

type Handler struct {
	UserRepo repository.UserRepositoryImpl
	Ctx      context.Context
}

func NewCreateHandler(db *database.DB, ctx context.Context) *Handler {
	return &Handler{*repository.NewUserRepository(db.Pool), ctx}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u repository.UserModel

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserRepo.CreateAndGetId(h.Ctx, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Person: %+v", u)
}
