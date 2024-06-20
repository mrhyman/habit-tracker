package handler

import (
	"encoding/json"
	"errors"
	"main/internal/domain"
	"main/internal/usecase/createuser"
	"net/http"
)

func (h *HttpHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u UserModel

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd, err := createuser.NewCommand(u.Id, u.Nickname, u.Birthday, u.ActiveHabitId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.CreateUserHandler.Handle(cmd)
		if err != nil {
			if errors.Is(err, domain.ErrUserAlreadyExists) {
				http.Error(w, err.Error(), http.StatusConflict)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}
