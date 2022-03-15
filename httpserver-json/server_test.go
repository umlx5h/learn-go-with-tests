package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
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

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPalyers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Floyd":  10,
			"Pepper": 20,
		},
		winCalls: nil}

	server := NewPlayerServer(store)

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

	server := NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusAccepted, response.Code)

		require.Equal(t, 1, len(store.winCalls))
		assert.Equal(t, "Pepper", store.winCalls[0])
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{league: wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		// if response.Result().Header.Get("content-type") != jsonContentType {
		// 	t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		// }

		server.ServeHTTP(response, request)
		require.Equal(t, http.StatusOK, response.Code)

		got := getLeagueFromResponse(t, response.Body)

		assert.Equal(t, wantedLeague, got)
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	// bytes, err := io.ReadAll(response.Body)
	// json.Unmarshal(bytes, &got)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return request
}
