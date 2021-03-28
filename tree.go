package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 6}
	return root
}

func NewBST(nums []int) *TreeNode {
	var root *TreeNode
	for _, num := range nums {
		root = insert(root, num)
	}
	return root
}

func PreOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	*ans = append(*ans, root.Val)
	PreOrder(root.Left, ans)
	PreOrder(root.Right, ans)
}

func InOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, ans)
	fmt.Println(root.Val)
	*ans = append(*ans, root.Val)
	InOrder(root.Right, ans)
}

func PostOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.Left, ans)
	PostOrder(root.Right, ans)
	fmt.Println(root.Val)
	*ans = append(*ans, root.Val)
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

func search(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.Val {
		return true
	}
	if val < root.Val {
		return search(root.Left, val)
	} else {
		return search(root.Right, val)
	}
}
