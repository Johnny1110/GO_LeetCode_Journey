# 98. Validate BST

<br>

---

<br>

link: https://leetcode.com/problems/validate-binary-search-tree/description/

> Given the root of a binary tree, determine if it is a valid binary search tree (BST).

<br>
<br>

## Coding-1

> 這一題卡關需要重新寫 2 遍加深記憶．

```go
func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MaxInt64, math.MinInt64)
}

func validateBST(node *TreeNode, max, min int) bool {
	if node == nil {
		return true
	}

	if node.Val >= max || node.Val <= min {
		return false
	}

	return validateBST(node.Left, node.Val, min) && validateBST(node.Right, max, node.Val)
}
```

