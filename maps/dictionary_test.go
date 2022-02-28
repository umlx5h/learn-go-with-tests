package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known_word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assert.Equal(t, want, got)
	})

	t.Run("unknown_word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assert.EqualError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new_word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assert.NoError(t, err)

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assert.Equal(t, definition, got)
	})

	t.Run("existing_word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assert.EqualError(t, err, ErrWordExists.Error())

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assert.Equal(t, definition, got)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
