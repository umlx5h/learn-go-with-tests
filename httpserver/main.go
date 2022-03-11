package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return 123, nil
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
}

func main() {
	// handler := http.HandlerFunc(PlayerServer)
	server := &PlayerServer{store: &InMemoryPlayerStore{}}

	// http.Handle("/score", handler)
	// log.Fatal(http.ListenAndServe(":5555", nil))
	log.Fatal(http.ListenAndServe(":5555", server))
}
