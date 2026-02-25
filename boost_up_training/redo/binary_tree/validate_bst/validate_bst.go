package validate_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var validate func(node *TreeNode, max, min int) bool
	validate = func(node *TreeNode, max, min int) bool {
		if node == nil {
			return true
		}

		if node.Val >= max || node.Val <= min {
			return false
		}

		return validate(node.Left, node.Val, min) && validate(node.Right, max, node.Val)
	}

	return validate(root, math.MaxInt64, math.MinInt64)
}
