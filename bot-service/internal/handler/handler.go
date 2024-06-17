//go:generate minimock -g -s .go -o ../../mocks/handler/http
package handler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"main/internal/domain"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
	"net/http"
)

type iCreateUser interface {
	Handle(cmd createuser.Command) error
}

type iGetUser interface {
	Handle(q getuserbyid.Query) (*domain.User, error)
}

type HttpHandler struct {
	CreateUserHandler  iCreateUser
	GetUserByIdHandler iGetUser
}

func (h *HttpHandler) SetupMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/hello", h.Hello())
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("POST /createUser", h.CreateUser())
	mux.Handle("GET /getUser", h.GetUserById())

	return mux
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
