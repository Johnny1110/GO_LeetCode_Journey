# 543. Diameter of Binary Tree

<br>

---

<br>

Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.

<br>

```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
```

## Coding

```go
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
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = number of nodes

Time: O(n) -> DFS iterate through all nodes

Space: O(log n) -> DFS recursive call, max call stack is tree height. 
```

<br>
<br>

## AI Feedback

### Space Complexity: $O(h)$

* Refinement: You noted $O(\log n)$, which is true for a balanced binary tree.

* However, in the worst-case scenario (a skewed tree, essentially a linked list), the height $h$ can be $n$.

* Therefore, it is more accurate to say the space complexity is $O(h)$, where $h$ is the height of the tree, or $O(n)$ in the absolute worst case.

<br>
<br>

## Refine

```go
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    
    var height func(*TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        left := height(node.Left)
        right := height(node.Right)
        
        // The diameter at this node is the sum of left and right heights
        if left + right > maxDiameter {
            maxDiameter = left + right
        }
        
        // Return height of this node to parent
        return max(left, right) + 1
    }
    
    height(root)
    return maxDiameter
}
```