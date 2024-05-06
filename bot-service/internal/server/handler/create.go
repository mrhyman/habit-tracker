package handler

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
	"main/internal/database/repository"
	"net/http"
	"time"
)

var ErrNicknameIsMissing = errors.New("user nickname is missing")

func (h *HttpHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u repository.UserDTO
		var record repository.UserRecord

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(u.Nickname) == 0 {
			http.Error(w, ErrNicknameIsMissing.Error(), http.StatusBadRequest)
			return
		}

		u.Id = uuid.New()
		u.CreatedAt = time.Now()

		err = h.UserRepository.CreateAndGetId(h.Ctx, record.FromDTO(u))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Person: %+v\n", u)
	})
}
