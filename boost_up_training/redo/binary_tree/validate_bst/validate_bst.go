package validate_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	left := root.Left
	right := root.Right

	if left != nil && left.Val >= root.Val {
		return false
	}
	if right != nil && right.Val <= root.Val {
		return false
	}

	return validLeftBST(root.Left, root) && validRightBST(root.Right, root)
}

func validLeftBST(node, parent *TreeNode) bool {
	if node == nil {
		return true
	}

	left := node.Left
	right := node.Right

	if left != nil && (left.Val >= node.Val || left.Val >= parent.Val) {
		return false
	}

	if right != nil && (right.Val <= node.Val || right.Val >= parent.Val) {
		return false
	}

	return validLeftBST(node.Left, node) && validRightBST(node.Right, node)
}

func validRightBST(node, parent *TreeNode) bool {
	if node == nil {
		return true
	}

	left := node.Left
	right := node.Right

	if left != nil && (left.Val >= node.Val || left.Val <= parent.Val) {
		return false
	}

	if right != nil && (right.Val <= node.Val || right.Val <= parent.Val) {
		return false
	}

	return validLeftBST(node.Left, node) && validRightBST(node.Right, node)
}

// [5,4,6,null,null,3,7] -> false
// [45,42,null,null,44,43,null,41] -> false
