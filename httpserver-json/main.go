package main

import (
	"log"
	"net/http"
)

func main() {
	// handler := http.HandlerFunc(PlayerServer)
	server := NewPlayerServer(NewInMemoryPlayerStore())

	// http.Handle("/score", handler)
	// log.Fatal(http.ListenAndServe(":5555", nil))
	log.Fatal(http.ListenAndServe(":5555", server))
}
