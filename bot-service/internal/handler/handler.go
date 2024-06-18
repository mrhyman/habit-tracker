//go:generate minimock -g -s .go -o ../../mocks/handler/http
package handler

import (
	"bytes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log/slog"
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

	//роутер прицепить
	mux.Handle("/hello", loggingMiddleware(h.Hello()))
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("POST /createUser", loggingMiddleware(h.CreateUser()))
	mux.Handle("GET /getUser", loggingMiddleware(h.GetUserById()))

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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		slog.Info("request received",
			"method", r.Method,
			"url", r.URL.String(),
			"body", string(bodyBytes),
		)

		next.ServeHTTP(w, r)
	})
}
