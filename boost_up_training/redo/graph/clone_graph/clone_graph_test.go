package clone_graph

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name     string
		graph    *Node
		expected [][]int
	}{
		{
			name:     "empty graph",
			graph:    nil,
			expected: nil,
		},
		{
			name: "single node",
			graph: &Node{
				Val:       1,
				Neighbors: []*Node{},
			},
			expected: [][]int{{1}},
		},
		{
			name:     "two nodes connected",
			graph:    createTwoNodesGraph(),
			expected: [][]int{{1, 2}, {2, 1}},
		},
		{
			name:     "four nodes cycle (LeetCode example 1)",
			graph:    createFourNodesCycle(),
			expected: [][]int{{1, 2, 4}, {2, 1, 3}, {3, 2, 4}, {4, 1, 3}},
		},
		{
			name:     "complete graph with 3 nodes",
			graph:    createCompleteThreeNodesGraph(),
			expected: [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}},
		},
		{
			name:     "linear chain of 4 nodes",
			graph:    createLinearChain(),
			expected: [][]int{{1, 2}, {2, 1, 3}, {3, 2, 4}, {4, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloned := cloneGraph(tt.graph)
			
			if tt.graph == nil {
				if cloned != nil {
					t.Errorf("Expected nil for empty graph, got non-nil")
				}
				return
			}

			if cloned == nil {
				t.Errorf("Expected non-nil cloned graph, got nil")
				return
			}

			if !validateClonedGraph(tt.graph, cloned, tt.expected) {
				t.Errorf("Cloned graph structure is incorrect")
			}

			if sharesNodes(tt.graph, cloned) {
				t.Errorf("Cloned graph shares nodes with original (deep copy required)")
			}
		})
	}
}

func createTwoNodesGraph() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	
	node1.Neighbors = []*Node{node2}
	node2.Neighbors = []*Node{node1}
	
	return node1
}

func createFourNodesCycle() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}
	
	return node1
}

func createCompleteThreeNodesGraph() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	
	node1.Neighbors = []*Node{node2, node3}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node1, node2}
	
	return node1
}

func createLinearChain() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	
	node1.Neighbors = []*Node{node2}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node3}
	
	return node1
}

func validateClonedGraph(original, cloned *Node, expected [][]int) bool {
	visited := make(map[*Node]bool)
	clonedNodes := collectNodes(cloned, visited)
	
	if len(clonedNodes) != len(expected) {
		return false
	}

	for _, node := range clonedNodes {
		nodeIdx := node.Val - 1
		if nodeIdx < 0 || nodeIdx >= len(expected) {
			return false
		}
		
		expectedNeighbors := expected[nodeIdx]
		if len(node.Neighbors) != len(expectedNeighbors)-1 {
			return false
		}
		
		for _, neighbor := range node.Neighbors {
			found := false
			for _, expectedNeighborVal := range expectedNeighbors[1:] {
				if neighbor.Val == expectedNeighborVal {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	
	return true
}

func collectNodes(node *Node, visited map[*Node]bool) []*Node {
	if node == nil || visited[node] {
		return []*Node{}
	}
	
	visited[node] = true
	nodes := []*Node{node}
	
	for _, neighbor := range node.Neighbors {
		nodes = append(nodes, collectNodes(neighbor, visited)...)
	}
	
	return nodes
}

func sharesNodes(original, cloned *Node) bool {
	originalNodes := make(map[*Node]bool)
	collectOriginalNodes(original, originalNodes)
	
	return checkForSharedNodes(cloned, originalNodes, make(map[*Node]bool))
}

func collectOriginalNodes(node *Node, nodes map[*Node]bool) {
	if node == nil || nodes[node] {
		return
	}
	
	nodes[node] = true
	
	for _, neighbor := range node.Neighbors {
		collectOriginalNodes(neighbor, nodes)
	}
}

func checkForSharedNodes(cloned *Node, originalNodes map[*Node]bool, visited map[*Node]bool) bool {
	if cloned == nil || visited[cloned] {
		return false
	}
	
	if originalNodes[cloned] {
		return true
	}
	
	visited[cloned] = true
	
	for _, neighbor := range cloned.Neighbors {
		if checkForSharedNodes(neighbor, originalNodes, visited) {
			return true
		}
	}
	
	return false
}
