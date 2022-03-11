package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	// handler := http.HandlerFunc(PlayerServer)
	server := &PlayerServer{store: &InMemoryPlayerStore{}}

	// http.Handle("/score", handler)
	// log.Fatal(http.ListenAndServe(":5555", nil))
	log.Fatal(http.ListenAndServe(":5555", server))
}
