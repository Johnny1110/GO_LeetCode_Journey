package binary_tree_max_imumpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	bestSum := math.MinInt32

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)

		bestPathSum := max(leftSum, rightSum, 0)

		oneWaySum := node.Val + bestPathSum // could be only node itself, if both left&right are negative

		// v way option:
		vWaySum := node.Val + leftSum + rightSum

		bestSum = max(bestSum, oneWaySum, vWaySum)

		return oneWaySum

	}

	dfs(root)

	return bestSum
}
