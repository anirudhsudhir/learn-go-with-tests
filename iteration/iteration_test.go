package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Print 5 characters iteratively", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleRepeat() {
	repeatedChar := Repeat("A", 3)
	fmt.Println(repeatedChar)
	// Output: AAA
}
