package main

import (
	"log"
	"net/http"

	poker "github.com/umlx5h/learn-go-with-tests/httpserver-cli"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5555", server))
}
