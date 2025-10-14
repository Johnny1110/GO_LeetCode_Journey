# 98. Validate Binary Search Tree

<br>

---

<br>

link: https://leetcode.com/problems/validate-binary-search-tree/description/

> Given the root of a binary tree, determine if it is a valid binary search tree (BST).

<br>

## Thinking

A valid BST is defined as follows:

* The left subtree of a node contains only nodes with keys strictly less than the node's key.
* The right subtree of a node contains only nodes with keys strictly greater than the node's key.
* Both the left and right subtrees must also be binary search trees.

<br>

**Topic**

* Tree
* Depth-First Search
* Binary Search Tree
* Binary Tree

<br>

So left leaf must smaller than root, and right leaf must bigger than root.

And I try like this:

```go
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	rootVal := root.Val

	if root.Left != nil && root.Left.Val >= rootVal {
		return false
	}
	if root.Right != nil && root.Right.Val <= rootVal {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
```

It is wrong, because it can not identify a invalid binary-tree like this:

```go
    5
  /   \
 4     6
      / \
     3   7
```

3 is lower than 4, it should be stay in left side of node-4.

<br>

The Key Insight: __Think about what valid range each node can have:__

* The root can be any value
* When you go left, you're adding an upper bound (must be less than parent)
* When you go right, you're adding a lower bound (must be greater than parent)

As you traverse down the tree, each node inherits constraints from all its ancestors, not just its immediate parent.

<br>

**Hint**:

1. Track the valid range: As you recurse, pass along what the valid range is for each node
2. Initial range: Start with `(-∞, +∞)` for the root
3. Update the range: When going left, what changes? When going right, what changes?

<br>
<br>

## Coding

```go
func isValidBST(root *TreeNode) bool {
    
}
```