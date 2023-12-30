package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test string"}
	t.Run("known word", func(t *testing.T) {
		search, want := "test", "test string"
		assertDefinition(t, dictionary, want, search)
	})
	t.Run("unknown word", func(t *testing.T) {
		search := "lorem"
		_, err := dictionary.Search(search)
		assertError(t, err, WordNotPresent)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		search, want := "test", "test string"
		err := dictionary.Add("test", "test string")
		if err != nil {
			t.Fatal("did not expect an error but received one")
		}
		assertDefinition(t, dictionary, want, search)
	})
	t.Run("existing word", func(t *testing.T) {
		search, want := "test", "test string"
		dictionary := Dictionary{search: want}
		err := dictionary.Add(search, want)
		assertError(t, err, wordAlreadyPresent)
		assertDefinition(t, dictionary, want, search)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		search, want := "test", "new test string"
		dictionary := Dictionary{search: "test string"}
		dictionary.Update(search, want)
		assertDefinition(t, dictionary, want, search)
	})
	t.Run("update new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "test string")
		assertError(t, err, UpdateNewWord)
	})
}

func TestDelete(t *testing.T) {
	search, want := "test", "test string"
	dictionary := Dictionary{search: want}
	dictionary.Delete(search)
	_, err := dictionary.Search(search)
	if err != WordNotPresent {
		t.Errorf("expected %q to be deleted", search)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, want, search string) {
	t.Helper()
	got, err := dictionary.Search(search)
	if err != nil {
		t.Fatal("did not expect an error but received one")
	}
	if got != want {
		t.Errorf("got %q, want %q, for search %q", got, want, search)
	}
}
