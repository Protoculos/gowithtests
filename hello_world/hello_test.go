package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Add name", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("Add empty name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Franch", func(t *testing.T) {
		got := Hello("Et'en", "Franch")
		want := "Bonjour, Et'en"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
