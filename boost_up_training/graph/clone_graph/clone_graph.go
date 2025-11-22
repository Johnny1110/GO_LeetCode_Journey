package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

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
