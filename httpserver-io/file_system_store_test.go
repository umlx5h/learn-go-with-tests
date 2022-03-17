package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanFile := createTempFile(t, `
	[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}
	]
	`)
	defer cleanFile()

	t.Run("league from a reader", func(t *testing.T) {
		store, err := NewFileSystemPlayerStore(database)
		require.NoError(t, err)
		got := store.GetLeague()

		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assert.Equal(t, want, got)

		got = store.GetLeague()
		assert.Equal(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		store, err := NewFileSystemPlayerStore(database)
		require.NoError(t, err)

		got, _ := store.GetPlayerScore("Chris")
		want := 33

		assert.Equal(t, want, got)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database2, cleanFile := createTempFile(t, `
		[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]
	`)
		defer cleanFile()

		store, err := NewFileSystemPlayerStore(database2)
		require.NoError(t, err)

		store.RecordWin("Chris")

		got, _ := store.GetPlayerScore("Chris")
		want := 34
		assert.Equal(t, want, got)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database2, cleanFile := createTempFile(t, `
		[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]
	`)
		defer cleanFile()

		store, err := NewFileSystemPlayerStore(database2)
		require.NoError(t, err)

		store.RecordWin("Pepper")

		got, _ := store.GetPlayerScore("Pepper")
		want := 1
		assert.Equal(t, want, got)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, clean := createTempFile(t, "")
		defer clean()

		_, err := NewFileSystemPlayerStore(database)
		require.NoError(t, err)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	require.NoError(t, err, "could not parse temp file")

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
