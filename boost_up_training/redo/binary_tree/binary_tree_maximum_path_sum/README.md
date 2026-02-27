# 124. Binary Tree Maximum Path Sum

<br>

---

<br>

link: https://leetcode.com/problems/binary-tree-maximum-path-sum

<br>
<br>

Using DFS + DP

```go

```



## Coding-1

```go

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
```