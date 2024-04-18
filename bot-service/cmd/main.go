package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/config"
	"main/internal/database"
	helloHandler "main/internal/handler/hello"
	userHandler "main/internal/handler/user"

	"net/http"
)

func main() {
	ctx := context.Background()
	conf := config.MustLoad()

	db, err := database.New(ctx, conf.Database.Connection)

	if err != nil {
		log.Fatal("unable to create connection pool:", err)
	}

	//TODO: router map - url: handler
	http.Handle("/hello", helloHandler.NewHelloHandler())
	http.Handle("/create", userHandler.NewCreateHandler(db, ctx))

	fmt.Println("Starting service at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
