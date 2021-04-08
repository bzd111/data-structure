package main

import (
	"reflect"
	"testing"
)

func TestPreOrder(t *testing.T) {
	root := NewBinaryTree()
	result := []int{}
	PreOrder(root, &result)
	ans := []int{5, 3, 1, 4, 7, 6}
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("got %v, want %v", result, ans)
	}
}

func TestInOrder(t *testing.T) {
	root := NewBinaryTree()
	result := []int{}
	InOrder(root, &result)
	ans := []int{1, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("got %v, want %v", result, ans)
	}
}
func TestPostOrder(t *testing.T) {
	root := NewBinaryTree()
	result := []int{}
	PostOrder(root, &result)
	ans := []int{1, 4, 3, 6, 7, 5}
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("got %v, want %v", result, ans)
	}
}

func TestLevelOrder(t *testing.T) {
	root := NewBinaryTree()
	result := [][]int{}
	LeverlOrder(root, &result)
	ans := [][]int{{5}, {3, 7}, {1, 4, 6}}
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("got %v, want %v", result, ans)
	}
}

func TestNewBST(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := NewBST(nums)
	result := []int{}
	InOrder(root, &result)
	if !reflect.DeepEqual(nums, result) {
		t.Errorf("got %v, want %v", result, result)
	}
}

func TestBST(t *testing.T) {
	t.Run("create BST", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		root := NewBST(nums)
		result := []int{}
		InOrder(root, &result)
		if !reflect.DeepEqual(nums, result) {
			t.Errorf("got %v, want %v", result, result)
		}
	})
	t.Run("BST search", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		root := NewBST(nums)
		exist := search(root, 3)
		if exist == false {
			t.Errorf("search not exist")
		}
		notExist := search(root, 8)
		if notExist == true {
			t.Errorf("search exist")
		}
	})

}
