package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"example": "it is just a small example."}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("example")
		want := "it is just a small example."

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("no")
		assertError(t, got, ErrNotFOund)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "new item"
		defintion := "this is a new addition to the dictionary"

		err := dictionary.Add(word, defintion)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, defintion)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"new item": "this is a new addition to the dictionary"}

		word := "new item"
		defintion := "this is a new addition to the dictionary"

		err := dictionary.Add(word, defintion)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, defintion)

	})

}

func TestUpdate(t *testing.T) {

	t.Run("excistiong word", func(t *testing.T) {
		word := "new item"
		definition := "this is a new addition to the dictionary"
		updatedDefinition := "This is the updated definition"

		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, updatedDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updatedDefinition)

	})

	t.Run("new word", func(t *testing.T) {
		word := "new item"
		updatedDefinition := "This is the updated definition"

		dictionary := Dictionary{}

		err := dictionary.Update(word, updatedDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "new item"
	definition := "this is a new defintion"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFOund {
		t.Errorf("Expected %q to be deleted", word)
	}

}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
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
