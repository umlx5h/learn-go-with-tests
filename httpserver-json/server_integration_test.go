package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	t.Run("sync", func(t *testing.T) {
		store := NewInMemoryPlayerStore()
		server := PlayerServer{store: store}
		player := "Pepper"

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "3", response.Body.String())
	})

	t.Run("concurrently", func(t *testing.T) {
		store := NewInMemoryPlayerStore()
		server := PlayerServer{store: store}
		player := "Pepper"

		num := 1000

		var wg sync.WaitGroup
		wg.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				defer wg.Done()
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
			}()
		}

		wg.Wait()

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, strconv.Itoa(num), response.Body.String())
	})
}
