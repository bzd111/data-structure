package main

import (
	"testing"
)

func TestHeap(t *testing.T) {

	popItems := func(t *testing.T, h *MinHeap, want int) {
		got := h.pop()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("New Heap", func(t *testing.T) {
		h := NewMinHeap()
		h.push(3)
		h.push(2)
		h.push(4)
		h.push(1)
		popItems(t, h, 1)
		popItems(t, h, 2)
		popItems(t, h, 3)
		popItems(t, h, 4)
	})

	t.Run("New Heap From Slice", func(t *testing.T) {
		data := []int{4, 3, 2, 1}
		h := NewMinHeapFromSlice(data)
		popItems(t, h, 1)
		popItems(t, h, 2)
		popItems(t, h, 3)
		popItems(t, h, 4)
	})
}
