package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("say 'Hello, Mike!' when the string 'Mike' is supplied", func(t *testing.T) {
		got := Hello("Mike", "English")
		want := "Hello, Mike!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, world' when an empty strign is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Michael", "Spanish")
		want := "Hola, Michael!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Michael", "French")
		want := "Bonjour, Michael!"

		assertCorrectMessage(t, got, want)
	})

}
