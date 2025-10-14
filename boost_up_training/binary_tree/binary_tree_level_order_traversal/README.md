# 102. Binary Tree Level Order Traversal

<br>

---

<br>

link: https://leetcode.com/problems/binary-tree-level-order-traversal/description/

> Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).


<br>
<br>

## Thinking

<br>

The Topic say we should using BFS (Breadth-First Search). than let go take some time to study it.

<br>

## BFS

BFS 是「廣度優先搜尋」（Breadth-First Search）的縮寫，是一種用於圖形或樹狀結構的搜尋演算法。
其特點是從起始節點開始，一層一層地向外擴散進行搜尋。 它會先遍歷所有離起始點為1 的節點，再遍歷離起始點為2 的節點，依此類推，直到所有可達的節點都被遍歷。 
BFS 通常使用 __佇列（先進先出)__ 的資料結構來儲存待探索的節點，並用於尋找圖中的最短路徑或最少步驟的問題。


**與 DFS（深度優先搜尋）的比較**:

* BFS 採用「一層一層」的方式遍歷，善於尋找最短路徑，使用佇列。
* DFS 則採取「一條路走到黑」的方式遍歷，善於尋找特定路徑或遍歷所有節點，通常使用堆疊（或遞迴）。

<br>

Now I know about the most important thing about how to using DFS to solve this problem is 
using queue to store the next level nodes.


I think I can try like make a queue to do like.

* put head node into that queue.
* pop a node form that queue, put that node's left and right son node into queue
* pop a node from that queue again, then this time and next time popped node is the first layer's nodes.
* then we can do iterate by each node order by layer-level then left to right.

<br>

**The problem is How to know the layer?**

Now we know how to iterate through the tree structure by each layer. but we still have no clue how to identify the different layer.

<br>

1. In a standard BFS, all nodes mix together
2. A queue only knows "first in, first out" (it has no concept of levels)
3. When I pop node 2's children (4, 5) and node 3's children (6, 7), they're all just in the queue together

```
      1          <- Level 0
     / \
    2   3        <- Level 1
   / \ / \
  4  5 6  7      <- Level 2
```

<br>

## Claude AI Solutions

> The key insight: At any moment, the queue contains nodes from at most 2 consecutive levels.

Before processing nodes, count how many nodes are currently in the queue - that's exactly one complete level.

<br>

## Thinking - 2

I think we can do more pop node action from each recursive call.

like if there are 3 layers, we should call levelOrder function 3 times only.

collect current layer result into that `[][]int`.

Then how do we know about how many nodes we should pop from that queue.

* Get current queue size at first, that the amount we should pop by current recursive call.

<br>

let try it out.


<br>
<br>

## Coding

```go
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	container []*TreeNode
}

func NewQueue() *Queue {
	return &Queue{
		container: make([]*TreeNode, 0),
	}
}

func (q *Queue) Push(node *TreeNode) {
	q.container = append(q.container, node)
}

func (q *Queue) Pop() (*TreeNode, bool) {
	if len(q.container) == 0 {
		return nil, false
	}
	node := q.container[0]
	q.container = q.container[1:]
	return node, true
}

func (q *Queue) Size() int {
	return len(q.container)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := NewQueue()
	// init, put head into queue.
	queue.Push(root)
	level := [][]int{}
	perform(queue, &level)

	return level
}

func perform(queue *Queue, result *[][]int) {
	currentSize := queue.Size()
	if currentSize == 0 {
		return
	}

	currentLayer := []int{}

	for range currentSize {
		node, ok := queue.Pop()
		if !ok {
			panic("invalid queue size")
		}

		currentLayer = append(currentLayer, node.Val)
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
	}

	*result = append(*result, currentLayer)
	perform(queue, result)
}
```

<br>

result:

![1.png](imgs/1.png)

<br>

the performance is poor...

<br>

## Revamp

<br>

Same idea, just remove unnecessary recursive call. instead, using 2 nested loop in one function.

<br>

```go
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	container []*TreeNode
}

func NewQueue() *Queue {
	return &Queue{
		container: make([]*TreeNode, 0),
	}
}

func (q *Queue) Push(node *TreeNode) {
	q.container = append(q.container, node)
}

func (q *Queue) Pop() (*TreeNode, bool) {
	if len(q.container) == 0 {
		return nil, false
	}
	node := q.container[0]
	q.container = q.container[1:]
	return node, true
}

func (q *Queue) Size() int {
	return len(q.container)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := NewQueue()
	// init, put head into queue.
	queue.Push(root)
	level := [][]int{}
    
	// using 2 nested loop.
	for queue.Size() > 0 {
		currentSize := queue.Size()
		currentLayer := []int{}
		for range currentSize {
			if node, ok := queue.Pop(); ok {
				currentLayer = append(currentLayer, node.Val)
				// put next layer's node into queue.
				if node.Left != nil {
					queue.Push(node.Left)
				}
				if node.Right != nil {
					queue.Push(node.Right)
				}
			} else {
				panic("queue is empty")
			}
		}
		level = append(level, currentLayer)
	}

	return level
}
```

![2.png](imgs/2.png)