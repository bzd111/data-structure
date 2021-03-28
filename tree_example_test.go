package main

import "testing"

func TestMaxVal(t *testing.T) {
	t.Run("traditional way", func(t *testing.T) {
		root := NewBinaryTree()
		got := FindMaxValTree(root)
		want := 7
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("recursive way", func(t *testing.T) {
		root := NewBinaryTree()
		got := FindMaxValTree2(root)
		want := 7
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
