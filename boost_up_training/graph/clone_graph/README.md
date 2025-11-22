# 133. Clone Graph

<br>

---

<br>

link: https://leetcode.com/problems/clone-graph

<br>
<br>

## Thinking

Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

<br>

Graph Properties:

* Undirected (edges work both ways)
* Connected (all nodes are reachable from the given node)
* No duplicate edges or self-loops

<br>

**This problem combines two fundamental CS concepts:**

1. Graph Traversal:
   We need to visit every node in the original graph exactly once. Think about it: How do you explore an unknown graph systematically?

2. The Mapping Problem:
   Here's the crucial insight: When we clone node A and it has a neighbor B, we need to connect the cloned A to the cloned B. But what if we haven't created the cloned B yet? Or what if we've already cloned B when we encounter it again through a different path?

<br>
<br>

### Deconstructing the Problem:

**Step 1: The Fundamental Challenge**

Consider this simple graph: `1 -- 2 -- 3` (where 1 connects to 2, and 2 connects to 3)

When we clone node 1:

* We create a new node 1'
* Node 1 has neighbor 2, so we need to add 2' as a neighbor of 1'
* But 2' doesn't exist yet! What do we do?

Question: What data structure would help us keep track of "original node → cloned node" mappings?

<br>

**Step 2: The Traversal Strategy**

* DFS (Depth-First Search): Go as deep as possible before backtracking
* BFS (Breadth-First Search): Explore all neighbors before going deeper

Question: For this cloning problem, does it matter which one we choose? What are the trade-offs?


<br>

**Step 3: Handling Cycles**

Graphs can have cycles. For example: `1 -- 2 -- 3 -- 1` (a triangle)

Question: What happens if we don't track which nodes we've already visited/cloned? Walk through what would happen starting from node 1.

<br>
<br>

### Algorithm Framework

1. Create a `map[*Node][*Node]` which is `Origin Node -> Clone Node`, if `map[*Node]` exist means we already created clone node for this origin Node.
2. Travel method using DFS
   * func dfs(originNode *Node) cloneNode *Node {...}: 
     * If input originNode already cloned(exists in the map), just return it.
     * If input originNode not clone yet, create cloneNode and put it into `map[*Node][*Node]`.
       * Iterate through all neighbors:
         * call dfs(neighbor) to generate cloned neighbor node.
         * add cloned neighbor node into clone node's neighbors
       * return cloneNode as result. 

Let's try this ~

<br>
<br>

## Coding - DFS

```go
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	// 1. create a map for store existing cloneNode
	// origin node -> clone node
	cloneMap := map[*Node]*Node{}

	var dfs func(*Node) *Node
	dfs = func(originNode *Node) *Node {
		cloneNode, ok := cloneMap[originNode]
		if ok { // already cloned.
			return cloneNode
		} else { // not clone yet.
			// do clone, Neighbors is empty at first.
			cloneNode = &Node{
				Val:       originNode.Val,
				Neighbors: make([]*Node, len(originNode.Neighbors)),
			}

			// put clone node into cloneMap.
			cloneMap[originNode] = cloneNode

			// process Neighbors
			for i, neighbor := range originNode.Neighbors {
				cloneNeighbor := dfs(neighbor)
				cloneNode.Neighbors[i] = cloneNeighbor
			}

			return cloneNode
		}
	}

	return dfs(node)
}
```

<br>

Result:

![1.png](imgs/1.png)

<br>
<br>

## Coding - BFS

BFS approach just like what we do in previous coding, just using BFS(queue) approach to replace DFS func.

想：

* 每一個 Node 一定都會輪到。
* 為每一個 Node 建立 Clone 後，放入 map
* 輪詢 Node 的 Neighbors 分別也建立 clone 並放入 map
* 最後把建立的 Clone Neighbors 依次加入到 Clone 的 Neighbors 中


```go
type Queue []*Node

func (q *Queue) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (*Node, bool) {
	if len(*q) == 0 {
		return nil, false
	}
	n := (*q)[0]
	*q = (*q)[1:]
	return n, true
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	// 1. create a map for store existing cloneNode
	// origin node -> clone node
	cloneMap := map[*Node]*Node{}
	queue := &Queue{}
	queue.Push(node)

	// 2. create head node's clone by manual.
	cloneMap[node] = &Node{Val: node.Val}

	for !queue.Empty() {
		currentNode, _ := queue.Pop()

		for _, neighbor := range currentNode.Neighbors {
			if _, exists := cloneMap[neighbor]; !exists {
				// first-time to see neighbor
				cloneMap[neighbor] = &Node{Val: neighbor.Val}
				queue.Push(neighbor)
			}

			cloneMap[currentNode].Neighbors = append(cloneMap[currentNode].Neighbors, cloneMap[neighbor])
		}

	}

	return cloneMap[node]
}
```

<br>

Result:

![2.png](imgs/2.png)

<br>
<br>
