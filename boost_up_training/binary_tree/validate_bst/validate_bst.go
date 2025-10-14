package validate_bst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	rootVal := root.Val

	if root.Left != nil && root.Left.Val >= rootVal {
		return false
	}

	if root.Right != nil && root.Right.Val <= rootVal {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
