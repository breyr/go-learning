package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brey", "")
		want := "Hello, Brey"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'hello world' with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Brey", "Spanish")
		want := "Hola, Brey"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Brey", "French")
		want := "Bonjour, Brey"
		assertCorrectMsg(t, got, want)
	})
}

// TB - interface that works with tests and benchmarks
func assertCorrectMsg(t testing.TB, got, want string) {
	// needed to tell function is a helper, line number reported if test fails goes to the function being tested, not this function
	t.Helper()
	// %q wraps values in ""
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
