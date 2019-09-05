package main

import (
	"log"
	"net/http"

	poker "github.com/26huitailang/go_tdd/time"
)

const dbFilename = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFilename)
	defer closeFunc()

	if err != nil {
		log.Fatal(err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
