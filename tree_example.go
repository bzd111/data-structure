package main

import "math"

func FindMaxValTree(root *TreeNode) int {

	_max := math.MinInt16
	traverse(root, &_max)
	return _max
}

func traverse(root *TreeNode, _max *int) {
	if root == nil {
		return
	}
	*_max = max(root.Val, *_max)
	traverse(root.Left, _max)
	traverse(root.Right, _max)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMaxValTree2(root *TreeNode) int {
	if root == nil {
		return math.MinInt16
	}
	maxLeft := FindMaxValTree2(root.Left)
	maxRight := FindMaxValTree2(root.Right)
	bigger := max(maxLeft, maxRight)
	return max(bigger, root.Val)
}
