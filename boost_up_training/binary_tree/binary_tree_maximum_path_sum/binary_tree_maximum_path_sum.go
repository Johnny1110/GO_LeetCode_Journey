package binary_tree_maximum_path_sum

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, sumMax := perform(root)
	return sumMax
}

// perform return 2 int
// 1. int: the value that could be extended by parent node.
// 2. int: current maximum value.
func perform(node *TreeNode) (int, int) {
	if node == nil {
		return 0, math.MinInt
	}

	// goes left:
	leftExtVal, leftMax := perform(node.Left)
	// goes right:
	rightExtVal, rightMax := perform(node.Right)

	currentExtVal := max(node.Val, leftExtVal+node.Val, rightExtVal+node.Val)
	currentV := node.Val + leftExtVal + rightExtVal

	currentMax := max(currentExtVal, leftMax, rightMax, currentV, node.Val)

	return currentExtVal, currentMax
}
