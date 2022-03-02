package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	t.Run("returns fast server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// defer slowServer.Close()
		// defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer3(slowURL, fastURL, time.Second)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.Equal(t, want, got)
	})

	t.Run("timeout server", func(t *testing.T) {
		timeoutServer := makeDelayedServer(2 * time.Second)
		timeout2Server := makeDelayedServer(3 * time.Second)

		_, err := Racer3(timeoutServer.URL, timeout2Server.URL, time.Millisecond*30)

		assert.Error(t, err, "must be timed out")

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
