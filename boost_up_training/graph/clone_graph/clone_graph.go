package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

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
