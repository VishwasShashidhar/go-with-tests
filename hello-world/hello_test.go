package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Vish", english)
		want := "Hello, Vish"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is passed", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Ami", spanish)
		want := "Hola, Ami"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Annu", french)
		want := "Bonjour, Annu"

		assertCorrectMessage(t, got, want)
	})

}
