package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/internal/db/repository"
	"net/http"
)

func main() {
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

		config, err := pgxpool.ParseConfig('jdbc:postgresql://localhost:6432/postgres')
		conn, err := pgx.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		//pgx + repo + client
		repo := repository.NewUserRepository(config)

		fmt.Fprintf(w, "Person: %+v", u)

		//args[0] = user.Id
		//args[1] = user.Nickname
		//args[2] = user.CreatedAt
		//args[3] = user.Birthday
		//args[4] = user.ActiveHabitId
	})

	fmt.Println("Starting service at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
