package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Barry", "French")
		want := "Bonjour, Barry"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello' world, if not is not provided", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
