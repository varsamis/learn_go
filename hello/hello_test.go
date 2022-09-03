package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Philip", "French")
		want := "Bonjour, Philip"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Greek", func(t *testing.T) {
		got := Hello("Βαρσάμη", "Greek")
		want := "Γειά σου, Βαρσάμη"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Holla, Elodie"

		assertCorrectMessage(t, got, want)
	})
}

func TestGreetingPrefix(t *testing.T) {
	t.Run("default prefix", func(t *testing.T) {
		got := greetingPrefix("")
		want := "Hello, "

		assertCorrectMessage(t, got, want)
	})

	t.Run("french prefix", func(t *testing.T) {
		got := greetingPrefix("French")
		want := "Bonjour, "

		assertCorrectMessage(t, got, want)
	})

	t.Run("greek greeting", func(t *testing.T) {
		got := greetingPrefix("Greek")
		want := "Γειά σου, "

		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish greeting", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Holla, "

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
