package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/umlx5h/learn-go-with-tests/httpserver-websocket"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
