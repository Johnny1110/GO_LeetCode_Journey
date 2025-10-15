package validate_bst

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return perform(root, math.MinInt64, math.MaxInt64)
}

func perform(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val >= max || node.Val <= min {
		return false
	}

	return perform(node.Left, min, node.Val) && perform(node.Right, node.Val, max)
}
