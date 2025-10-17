# 297. Serialize and Deserialize

<br>

---

<br>

link: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

<br>

> Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

> Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

> Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

<br>

Constraints:

* The number of nodes in the tree is in the range [0, 104].
* -1000 <= Node.val <= 1000

<br>


## Thinking

I saw many example by previous tree problems, there always show me the example by an Array like:

```go
[3, 5, 1, null, null, 7, 11]
```

I'm thinking about should I use this data structure to serialize tree?

<br>
<br>

## Claude AI Suggests

There are 2 solution to implement tree serialize and deserialize

* Preorder (DFS) with null markers
* Level-order (BFS)

<br>

### Preorder (DFS) with null markers

```
     1
   /   \
  2     3
   \   /
    4 5

→ "1,2,null,4,null,null,3,5,null,null,null"
```

**Pros:** Simpler to implement recursively, natural for DFS
**Cons:** More nulls stored, looks longer

* Recursive thinking with implicit state (call stack)
* How tree structure maps to linear sequences
* The elegant symmetry between serialize/deserialize

Core skill: Understanding how recursion "automatically" handles tree traversal

<br>
<br>

### Level-order (BFS)

```
     1
   /   \
  2     3
   \   /
    4 5

→ [1, 2, 3, null, 4, 5]
```

**Pros:** Compact, intuitive visualization
**Cons:** You need to track positions carefully, handle nulls for missing children

* Explicit state management with queues
* Parent-child index relationships
* Handling "layers" of a tree

Core skill: Thinking iteratively about tree structure, managing position tracking

<br>
<br>


## Thinking - Preorder (DFS) with null markers

<br>

Let me implement DFS Preorder first.

node's value range is -1000 ~ 1000, 1 bytes can only present -127 ~ 128, so we 2 bytes to present a node.

2 bytes can present  -32,768 ~ 32,767.

Nah, let me do it simple first. just convert to string and connect with `,`

<br>
<br>

## Coding

```go
import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ==========================================================

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}

	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	return fmt.Sprintf("%d,%s,%s", root.Val, left, right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// split element by ','
	strVals := strings.Split(data, ",")
	idx := 0

	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(strVals) || strVals[idx] == "nil" {
			idx++
			return nil
		}

		val, _ := strconv.Atoi(strVals[idx])
		idx++
		
		left := build()
		right := build()

		return &TreeNode{
			Val:   val,
			Left:  left,
			Right: right,
		}
	}

	return build()
}
```

<br>

![1.png](imgs/1.png)

<br>

let me explain what did I do:

**serialize**

```go
func (this *Codec) serialize(root *TreeNode) string {
	//	  1
	//	/   \
	// 2     3
	//  \   /
	//   4 5
	//
	//→ "1,2,nil,4,nil,nil,3,5,nil,nil,nil"

	if root == nil {
		return "nil"
	}

	val := root.Val

	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	return fmt.Sprintf("%d,%s,%s", val, left, right)
}
```

* `serialize(root *TreeNode)` input a node and return string
* When it comes to recursive func, all we need to do is figure out the responsibility of 1 call should take charge.
* What we can get from a node is it's value and it's left and right subtrees.
* we what we need to be done is append node's value and it's left & right subtree's structure.

<br>

**deserialize**

```go
func (this *Codec) deserialize(data string) *TreeNode {
	// split strVals from data
	strVals := strings.Split(data, ",")
	idx := 0

	// make a build func
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(strVals) || strVals[idx] == "nil" {
			// if current str_val is nil or reach the end, just return nil node.
			idx++
			return nil
		}
        
		// parse str_val to int val
		val, _ := strconv.Atoi(strVals[idx])
		// idx ++
		idx++
        
		// recursive call build to build left and right node
		left := build()
		right := build()

		return &TreeNode{
			Val:   val,
			Left:  left,
			Right: right,
		}
	}

	return build()
}
```

<br>
<br>

