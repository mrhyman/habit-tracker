package handler

import (
	"fmt"
	"github.com/google/uuid"
	"main/internal/domain"
	"main/internal/usecase/getuserbyid"
	"net/http"
)

func (h *HttpHandler) GetUserById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("id")

		if userId == "" {
			http.Error(w, domain.ErrIdInvalid.Error(), http.StatusBadRequest)
			return
		}

		userUuid, err := uuid.FromBytes([]byte(userId))
		if err != nil {
			http.Error(w, domain.ErrInvalidUserID.Error(), http.StatusBadRequest)
		}

		q, err := getuserbyid.NewQuery(userUuid)

		u, err := h.GetUserByIdHandler.Handle(h.Ctx, q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Person: %+v", u)
	})
}
