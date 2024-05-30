package handler

import (
	"encoding/json"
	"errors"
	"main/internal/domain"
	"main/internal/usecase/getuserbyid"
	"net/http"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
)

func (h *HttpHandler) GetUserById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("id")

		q, err := getuserbyid.NewQuery(userId)
		if err != nil {
			http.Error(w, errors.Join(ErrInvalidArgument, err).Error(), http.StatusBadRequest)
			return
		}

		u, err := h.GetUserByIdHandler.Handle(q)
		if err != nil {
			if errors.Is(err, domain.ErrUserNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
