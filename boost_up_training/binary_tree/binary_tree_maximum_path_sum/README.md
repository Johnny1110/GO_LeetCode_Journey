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

### Problem Deconstruction

1. Understanding What a "Path" Really Means

What constitutes a valid path:

* A path is a connected sequence of nodes where we travel alone edges.
* We can move in any direction (up from child to parent, or down for parent to child)
* Can't visit same node twice.
* The path can be any shape: straight up or down, or V-shape (up then down), or just a single node.

<br>

2. **Types of Paths Through a Node**

For any given node there are four possible scenarios for how a maximum path might involve that node.

* Type 1: Node alone
* Type 2: Node + left subtree path
* Type 3: Node + right subtree path
* Type 4: Bridge path (V-shaped)

<br>

3. The Critical Insight: **Two Different Values**

Here's the key realization that makes this problem tricky:

For each node, we need to track **TWO different values**:

1. **"Return Value"**: The maximum sum of a path that starts at this node and goes down to ONE of its subtrees (or just the node itself). This is what we can "offer" to the parent node as a possible extension.

2. **"Local Maximum"**: The maximum sum of ANY path that passes through this node (including the V-shaped path that uses both subtrees). This might be our answer but can't be extended further up.

**Why two values?** Because a V-shaped path (Type 4) creates a "complete" path that can't be extended by the parent. Once you've gone left-node-right, you can't add the parent without revisiting the node (which is illegal).

<br>

4. **The Recursive Challenge**

At each node, we need to:
- Get information from children (what's the best path they can offer upward?)
- Make a local decision (what's the best path through me?)
- Return useful information to parent (what can I offer upward?)
- Keep track of the global maximum (what's the best path seen so far?)

<br>

5. **Handling Negative Values**

Negative values add complexity because:
- We might choose NOT to extend a path if it would decrease our sum
- A child subtree might return a negative contribution, in which case we'd rather not include it
- Even the root might be negative, but we must include at least one node

<br>

6. **The Core Decision at Each Node**

At each node, we're essentially asking:
1. "What's the best path I can offer to my parent?" → max(node, node + left, node + right)
2. "What's the best path that peaks at me?" → max(node, node + left, node + right, node + left + right)
3. "Is this the best path in the entire tree?" → Compare with global max

<br>
<br>

### Visual Example of the Complexity

Consider this tree:
```
       5
      / \
     3   -2
    / \    \
   1   4    7
```

At node 3:
* Return value to parent: max(3, 3+1, 3+4) = 7 (node + right child)
* Local maximum through node: max(3, 3+1, 3+4, 3+1+4) = 8 (using both children)
* But we can only return 7 to parent, not 8!

__This separation between "what's best locally" and "what we can offer upward" is the heart of the problem.__

<br>
<br>

## Coding

```go
func maxPathSum(root *TreeNode) int {
	return 0
}
```
