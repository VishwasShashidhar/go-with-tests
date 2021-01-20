package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word missing", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("blaa")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("not expecting an error")
		}

		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		got, _ := dictionary.Search(word)

		assertError(t, err, ErrWordExists)
		assertStrings(t, got, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "new test"

		_ = dictionary.Update(word, newDefinition)
		got, _ := dictionary.Search(word)

		assertStrings(t, got, newDefinition)
	})

	t.Run("word doesn't exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "new test"

		err := dictionary.Update("notest", newDefinition)
		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		_ = dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("word doesn't exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		err := dictionary.Delete("blah")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected to get an error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
