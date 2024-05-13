package server

import (
	"context"
	"fmt"
	"log"
	"main/internal/database"
	"main/internal/database/repository"
	"main/internal/handler"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
	"net/http"
)

type Server struct {
	Db       *database.DB
	Instance *http.Server
	Ctx      context.Context
}

func New(ctx context.Context, db *database.DB) *Server {
	mux := http.NewServeMux()
	userRepo := repository.NewUserRepository(db.Pool)

	httpHandler := handler.New(
		ctx, createuser.NewCommandHandler(userRepo), getuserbyid.NewQueryHandler(userRepo),
	)
	mux.Handle("/hello", httpHandler.Hello())
	mux.Handle("POST /create", httpHandler.CreateUser())
	mux.Handle("GET /create", httpHandler.GetUserById())

	return &Server{
		Ctx: ctx,
		Db:  db,
		Instance: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (s *Server) Start() {
	fmt.Println("Starting service at port 8080")
	if err := s.Instance.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Shutdown() {
	log.Println("Shutting service down")
	if err := s.Instance.Shutdown(s.Ctx); err != nil {
		log.Fatal(err)
	}
}
