# 236. Lowest Common Ancestor (LCA)

<br>

---

<br>

link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

<br>
<br>

## Coding-1

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// using dfs to find
	var dfs func(node *TreeNode) (bool, *TreeNode)
	dfs = func(node *TreeNode) (bool, *TreeNode) {
		if node == nil {
			return false, nil
		}

		foundthis := false

		if node.Val == p.Val || node.Val == q.Val {
			foundthis = true
		}

		// left side dfs
		foundleft, leftRes := dfs(node.Left)
		if leftRes != nil {
			// already found in sub-nodes
			return true, leftRes
		}

		foundRight, rightRes := dfs(node.Right)
		if rightRes != nil {
			// already found in sub-nodes
			return true, rightRes
		}

		if (foundleft && foundRight) || (foundRight && foundthis) || (foundleft && foundthis) {
			// this node will be the result
			return true, node
		}

		return foundthis || foundleft || foundRight, nil
	}

	_, result := dfs(root)

	return result
}
```

<br>
<br>

## Coding-2

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// dfs left & right
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}
```