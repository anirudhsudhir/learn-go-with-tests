package synchronization

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing counter thrice", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertValue(t, counter, 3)
	})
	t.Run("incrementing counter concurrently", func(t *testing.T) {
		counter := NewCounter()
		want := 1000
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 1; i <= want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertValue(t, counter, want)
	})
}

func assertValue(t testing.TB, counter *Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
