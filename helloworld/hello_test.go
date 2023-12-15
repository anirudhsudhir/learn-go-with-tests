package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := returnHello("Jim", "")
		want := "Hello, Jim!"

		assertCorrectMessage(got, want, t)
	})
	t.Run("Say 'Hello, World! when an empty string is provided", func(t *testing.T) {
		got := returnHello("", "")
		want := "Hello, World!"

		assertCorrectMessage(got, want, t)
	})
	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := returnHello("Robert", "Spanish")
		want := "Hola, Robert!"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Say hello in French", func(t *testing.T) {
		got := returnHello("Robert", "French")
		want := "Bonjour, Robert!"
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
