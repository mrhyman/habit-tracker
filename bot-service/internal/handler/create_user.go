package handler

import (
	"encoding/json"
	"fmt"
	"main/internal/domain"
	"main/internal/usecase/createuser"
	"net/http"
)

func (h *HttpHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u domain.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd, err := createuser.NewCommand(u)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.CreateUserHandler.Handle(h.Ctx, cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Person: %+v", u)
	})
}
