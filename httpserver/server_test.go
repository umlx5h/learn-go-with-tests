package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	score, ok := s.scores[name]
	if !ok {
		return 0, fmt.Errorf("not found player %q", name)
	}
	return score, nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPalyers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Floyd":  10,
			"Pepper": 20,
		},
		winCalls: nil}

	server := &PlayerServer{store: store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "20", response.Body.String())
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "10", response.Body.String())
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "", response.Body.String())
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusAccepted, response.Code)

		require.Equal(t, 1, len(store.winCalls), "record win must be counted")
		assert.Equal(t, player, store.winCalls[0])
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: nil,
	}

	server := &PlayerServer{store: store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusAccepted, response.Code)

		require.Equal(t, 1, len(store.winCalls))
		assert.Equal(t, "Pepper", store.winCalls[0])
	})

}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return request
}
