# 236. Lowest Common Ancestor of a Binary Tree (LCA)

<br>

---

<br>

link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

> Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
> According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

<br>
<br>


## Thinking

<br>

When it comes to recursive call, the first thing is define function signature.

I need a func, and tell me how many node does it found in child nodes.

<br>
<br>

## Coding

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, node := find(root, p, q)
	return node
}

func find(node, p, q *TreeNode) (int, *TreeNode) {

	count := 0

	if node == nil {
		return 0, nil
	}

	if node == p || node == q {
		count++
	}

	countLeft, A := find(node.Left, p, q)
	countRight, B := find(node.Right, p, q)

	count = count + countLeft + countRight

	if count == 2 && A == nil && B == nil {
		return 2, node
	}

	if A != nil {
		return count, A
	}
	if B != nil {
		return count, B
	}

	return count, nil
}
```

<br>

![1.png](imgs/1.png)

<br>

Although I solved this problem, but I don't think I'm on the right track.

My code is so disgusting, definitely need a revamp.

<br>
<br>

## Revamp

- If we find `p` or `q`, return it
- If a subtree found nothing, return `nil`
- If **both** subtrees returned something → they found `p` and `q` separately → **current node is LCA**

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// why we don't need go further cuz we only find 1 node (p or q)?
		// because if root == p the answer at least is p or above.
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// if both side subtree returned nodes, then current node is LCA
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}
```
