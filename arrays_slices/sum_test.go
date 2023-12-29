package arraysslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got := Sum(arr)
	want := 15

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4, 5})
	want := []int{3, 12}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checksum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("sum tails of collections of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4}, []int{5, 6})
		want := []int{2, 4, 6}
		checksum(t, got, want)
	})

	t.Run("sum tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 6}, []int{7, 8})
		want := []int{0, 6, 8}
		checksum(t, got, want)
	})
}
