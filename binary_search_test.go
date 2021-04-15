package main

import "testing"

func TestBinarySearch(t *testing.T) {
	s := []int{1, 2, 5, 7, 8, 12}
	got := BinarySearch(s, 8)
	want := 4
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	got = BinarySearch(s, 6)
	want = -1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBound(t *testing.T) {

	t.Run("lowerBound", func(t *testing.T) {
		s := []int{1, 2, 2, 2, 4, 4, 5}
		got := lowerBound(s, 2)
		want := 1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("upperBound", func(t *testing.T) {
		s := []int{1, 2, 2, 2, 4, 4, 5}
		got := upperBound(s, 2)
		want := 4
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
