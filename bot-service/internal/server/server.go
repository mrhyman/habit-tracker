package server

import (
	"context"
	"fmt"
	"log"
	"main/internal/database"
	"main/internal/handler"
	"net/http"
)

type Server struct {
	Db       *database.DB
	Instance *http.Server
	Ctx      context.Context
}

func New(ctx context.Context, db *database.DB) *Server {
	mux := http.NewServeMux()

	httpHandler := handler.New(ctx, db)
	mux.Handle("/hello", httpHandler.Hello())
	mux.Handle("/create", httpHandler.CreateUser())

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
