package main

import (
	"log"
	"net/http"

	"github.com/go-pg/pg"
)

func init() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "go_tdd",
	})
	defer db.Close()
	createSchema(db)
}

func main() {

	// server := &PlayerServer{NewInMemoryPlayerStore()}
	server := &PlayerServer{NewPostgresPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
