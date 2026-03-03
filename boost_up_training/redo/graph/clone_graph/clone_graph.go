package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		val := node.Val
		if cpNode, ok := visited[val]; ok {
			return cpNode
		}

		cpNode := &Node{
			Val:       val,
			Neighbors: make([]*Node, len(node.Neighbors)), // empty neighbors
		}

		visited[val] = cpNode

		for idx, neighbor := range node.Neighbors {
			cpNode.Neighbors[idx] = dfs(neighbor)
		}

		return cpNode
	}

	return dfs(node)
}
