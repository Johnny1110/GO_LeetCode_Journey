1448. Count Good Nodes in Binary Tree

<br>

---

<br>

## Coding

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
```

<br>
<br>

## Time & Space Complexity

Assume n = tree size.

Time Complexity: we have to iterate through all the nodes, so it is O(n).
Space Complexity: O(h) where h is the height of the tree


In Go (and most languages), every time a function calls itself, a new stack frame is added to the call stack to keep track of local variables and the return address.

**h** is the maximum depth of the recursion.

