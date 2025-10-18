# 124. Binary Tree Maximum Path Sum

<br>


----

<br>

link: https://leetcode.com/problems/binary-tree-maximum-path-sum/description

<br>
<br>

> Given a binary tree, find the maximum path sum. A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
> The path sum of a path is the sum of the node's values in the path.
 
**Constraints**

* The number of nodes in the tree is in the range [1, 3×10⁴]
* -1000 ≤ Node.val ≤ 1000

<br>
<br>

## Thinking

* The path can start and end at any nodes (doesn't have to include root).
* Negative values are allowed, which adds complexity to the decision-making.
* A path must be connected.
* Each node can only be used once in a path

Now we summed up all the critical condition.

What's the theory and approach behind solving this problem?

<br>

**Topic**

* Dynamic Programming
* Tree
* Depth-First Search
* Binary Tree

<br>



<br>
<br>

## Coding

```go
func maxPathSum(root *TreeNode) int {
	return 0
}
```
