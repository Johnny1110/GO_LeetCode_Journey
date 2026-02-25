package validate_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MaxInt64, math.MinInt64)
}

func validateBST(node *TreeNode, max, min int) bool {
	if node == nil {
		return true
	}

	if node.Val >= max || node.Val <= min {
		return false
	}

	return validateBST(node.Left, node.Val, min) && validateBST(node.Right, max, node.Val)
}
