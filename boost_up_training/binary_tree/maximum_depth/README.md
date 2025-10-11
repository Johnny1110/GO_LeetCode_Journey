# 104. Maximum Depth of Binary Tree

<br>

---


<br>

link: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

<br>

## Thinking

<br>

I remember I did something before, when I try to implement Ethereum Merkle Tree, when it comes to calculate tree height, I fund a equation for that.
but I forgot what it is...

<br>

### Asking AI to recap that equation:

Binary tree (like Merkle trees typically), the relationship between nodes and height is:Given n nodes:

```
height = ⌈log₂(n + 1)⌉ - 1
```

Or equivalently:

```
height = ⌊log₂(n)⌋
```

But LeetCode's Problem is Different!

**Key difference:** 

The LeetCode "Maximum Depth of Binary Tree" problem gives you an arbitrary binary tree that might be:

* Unbalanced (one side much deeper than the other)
* Sparse (missing nodes randomly)
* A single long chain

<br>

So nope! we can not using that formula to solve this tree height problem.

<br>
<br>

then I'm think about should I have to go through all the nodes. I guess the answer is yes.

That will be a `DFS` (Depth-First Search)

<br>
<br>

## Coding

```go
func maxDepth(root *TreeNode) int {
	return recursiveCalculateDepth(root, 0)
}

func recursiveCalculateDepth(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth
	}

	currentDepth++

	maxLeft := recursiveCalculateDepth(root.Left, currentDepth)
	maxRight := recursiveCalculateDepth(root.Right, currentDepth)
	return max(maxLeft, maxRight)
}
```

I did like above, and the result is:

![1.png](imgs/1.png)

<br>

Unfortunately, I'm using Top-Bottom Approach, I may try using Bottom-Top Approach.

<br>
<br>

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	a := maxDepth(root.Left)
	b := maxDepth(root.Right)

	return 1 + max(a, b)
}
```

![2.png](imgs/2.png)