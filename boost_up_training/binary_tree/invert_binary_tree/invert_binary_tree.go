package invert_binary_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	perform(root)
	return root
}

func perform(root *TreeNode) {
	if root == nil {
		return
	}

	perform(root.Left)
	perform(root.Right)

	root.Left, root.Right = root.Right, root.Left
}
