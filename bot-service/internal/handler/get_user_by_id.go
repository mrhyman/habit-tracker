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

// GetUserByIdHandler godoc
//
//	@Summary	Get bot user
//	@Tags		handler
//	@Produce	json
//	@Param		id	query	string	true	"uuid formatted ID"
//	@Router		/user/search [get]
//	@Success	200	{object}	UserModel
//	@Failure	400	{string}	Bad		Request
//	@Failure	404	{string}	User	Not	Found
//	@Failure	500	{string}	Server	Error
func (h *HttpHandler) GetUserById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("id")

		q, err := getuserbyid.NewQuery(userId)
		if err != nil {
			http.Error(w, errors.Join(ErrInvalidArgument, err).Error(), http.StatusBadRequest)
			return
		}

		u, err := h.GetUserByIdHandler.Handle(r.Context(), q)
		if err != nil {
			if errors.Is(err, domain.ErrUserNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(UserFromDomain(u))
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
