package main

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func NewMinHeapFromSlice(data []int) *MinHeap {
	h := NewMinHeap()
	n := len(data) - 1
	h.data = data
	for i := (n - 1) / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

func (h *MinHeap) push(item int) {
	h.data = append(h.data, item)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) pop() int {
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return last
}

func (h *MinHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if h.data[index] >= h.data[parent] {
		return
	}
	swap(&h.data[index], &h.data[parent])
	h.heapifyUp(parent)

}

func (h *MinHeap) heapifyDown(index int) {
	smallest := index
	for i := 2*index + 1; i <= 2*index+2; i++ {
		if i < len(h.data) && h.data[i] < h.data[smallest] {
			smallest = i
		}
	}
	if smallest == index {
		return
	}
	swap(&h.data[index], &h.data[smallest])
	h.heapifyDown(smallest)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
