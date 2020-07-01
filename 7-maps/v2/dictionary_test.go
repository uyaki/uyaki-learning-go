package v1

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	assertError := func(t *testing.T, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error '%s' want '%s'", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		_, got := dictionary.Search("test")
		assertError(t, got, ErrNotFound)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}
