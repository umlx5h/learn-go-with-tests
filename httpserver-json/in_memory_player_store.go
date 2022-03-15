package main

import (
	"fmt"
	"sync"
)

type InMemoryPlayerStore struct {
	scores map[string]int
	mu     sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		scores: make(map[string]int),
	}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	if score, ok := i.scores[name]; !ok {
		return 0, fmt.Errorf("not found player %q", name)
	} else {
		return score, nil
	}
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	i.scores[name]++
	i.mu.Unlock()
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return []Player{}
}
