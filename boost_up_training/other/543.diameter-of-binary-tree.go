/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftLen, leftBest := dfs(node.Left)
		rightLen, rightBest := dfs(node.Right)

		retLen := max(leftLen, rightLen) + 1

		thisBest := leftLen + rightLen + 1

		return retLen, max(thisBest, leftBest, rightBest)
	}

	_, best := dfs(root)
	return best - 1
}

// @lc code=end

