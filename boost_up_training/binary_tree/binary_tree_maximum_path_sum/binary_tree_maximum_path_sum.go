package binary_tree_maximum_path_sum

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
	extendableVal, submax := perform(root, root.Val)
	return maxOf3(root.Val, extendableVal, submax)
}

func maxOf3(a int, b int, c int) int {
	val := max(a, b)
	return max(val, c)
}

// perform return 2 int
// 1. int: the value that could be extended by parent node.
// 2. int: current maximum value.
func perform(node *TreeNode, currentMax int) (int, int) {
	if node == nil {
		return 0, currentMax
	}

	// goes left:
	leftExtVal, leftSubMax := perform(node.Left, currentMax)

	// goes right:
	rightExtVal, rightSubMax := perform(node.Right, currentMax)

	currentExtVal := maxOf3(node.Val, leftExtVal+node.Val, rightExtVal+node.Val)
	currentV := node.Val + leftExtVal + rightExtVal

	bestSubMax := max(leftSubMax, rightSubMax)

	currentMax = maxOf3(max(currentMax, bestSubMax), currentV, currentExtVal)

	return currentExtVal, currentMax
}
