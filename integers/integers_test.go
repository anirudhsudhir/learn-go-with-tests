package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Add 2 numbers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		assertCorrectAnswer(t, got, want)
	})
}

func assertCorrectAnswer(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Printf("%d", sum)
	// Output:7
}
