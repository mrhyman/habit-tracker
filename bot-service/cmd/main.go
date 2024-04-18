package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/internal/config"
	"main/internal/db/repository"
	"net/http"
)

func main() {

	ctx := context.Background()

	conf := config.MustLoad()

	conn, err := pgxpool.New(ctx, conf.Database.Connection)
	if err != nil {
		log.Fatal("unable to create connection pool:", err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {

		var u repository.UserModel

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(u)

		repo := repository.NewUserRepository(conn)
		err = repo.CreateAndGetId(ctx, &u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Person: %+v", u)
	})

	fmt.Println("Starting service at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
