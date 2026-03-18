package cgnibt

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0

	var dfs func(node *TreeNode, currentMax int)
	dfs = func(node *TreeNode, currentMax int) {
		if node == nil {
			return
		}

		thisVal := node.Val
		if thisVal >= currentMax {
			count++
			currentMax = thisVal
		}

		dfs(node.Left, currentMax)
		dfs(node.Right, currentMax)
	}

	dfs(root, root.Val)

	return count
}
