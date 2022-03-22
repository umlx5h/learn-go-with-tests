package poker

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	t.Run("sync", func(t *testing.T) {
		db, clean := createTempFile(t, `[]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(db)
		require.NoError(t, err)

		server := NewPlayerServer(store)
		player := "Pepper"

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		t.Run("get score", func(t *testing.T) {
			response := httptest.NewRecorder()
			server.ServeHTTP(response, newGetScoreRequest(player))

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, "3", response.Body.String())
		})

		t.Run("get league", func(t *testing.T) {
			response := httptest.NewRecorder()
			server.ServeHTTP(response, newLeagueRequest())

			assert.Equal(t, http.StatusOK, response.Code)

			got := getLeagueFromResponse(t, response.Body)
			want := []Player{
				{"Pepper", 3},
			}

			assert.Equal(t, want, got)

		})
	})

	t.Run("concurrently", func(t *testing.T) {
		db, clean := createTempFile(t, `[]`)
		defer clean()
		store, err := NewFileSystemPlayerStore(db)
		require.NoError(t, err)
		server := NewPlayerServer(store)
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
