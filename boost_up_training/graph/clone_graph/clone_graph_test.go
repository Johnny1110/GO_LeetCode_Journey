package clone_graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
	//Output: [[2,4],[1,3],[2,4],[1,3]]

	// Build original graph
	node_1 := &Node{Val: 1}
	node_2 := &Node{Val: 2}
	node_3 := &Node{Val: 3}
	node_4 := &Node{Val: 4}

	node_1.Neighbors = []*Node{node_2, node_4}
	node_2.Neighbors = []*Node{node_1, node_3}
	node_3.Neighbors = []*Node{node_2, node_4}
	node_4.Neighbors = []*Node{node_1, node_3}

	// Clone the graph
	cloned_1 := cloneGraph(node_1)

	// Test 1: Verify it's a deep copy (different memory addresses)
	assert.NotSame(t, node_1, cloned_1, "Cloned node should be a different object")

	// Test 2: Verify node values
	assert.Equal(t, node_1.Val, cloned_1.Val)

	// Test 3: Verify neighbor count
	cloned_1_neighbors := cloned_1.Neighbors
	assert.Equal(t, len(node_1.Neighbors), len(cloned_1_neighbors))

	// Get cloned nodes through neighbors
	cloned_2 := cloned_1_neighbors[0]
	cloned_4 := cloned_1_neighbors[1]

	// Test 4: Verify neighbor values
	assert.Equal(t, node_2.Val, cloned_2.Val)
	assert.Equal(t, node_4.Val, cloned_4.Val)

	// Test 5: Verify these are also different objects (deep copy)
	assert.NotSame(t, node_2, cloned_2, "Cloned node 2 should be different object")
	assert.NotSame(t, node_4, cloned_4, "Cloned node 4 should be different object")

	// Test 6: Verify complete structure - check node 2's neighbors
	assert.Equal(t, len(node_2.Neighbors), len(cloned_2.Neighbors))
	assert.Equal(t, 2, len(cloned_2.Neighbors)) // Should have nodes 1 and 3

	// Find cloned_3 through cloned_2's neighbors
	var cloned_3 *Node
	for _, neighbor := range cloned_2.Neighbors {
		if neighbor.Val == 3 {
			cloned_3 = neighbor
			break
		}
	}
	assert.NotNil(t, cloned_3, "Should find node 3 in cloned graph")
	assert.Equal(t, node_3.Val, cloned_3.Val)
	assert.NotSame(t, node_3, cloned_3, "Cloned node 3 should be different object")

	// Test 7: Verify node 3's neighbors
	assert.Equal(t, len(node_3.Neighbors), len(cloned_3.Neighbors))
	assert.Equal(t, 2, len(cloned_3.Neighbors)) // Should have nodes 2 and 4

	// Test 8: Verify node 4's complete structure
	assert.Equal(t, len(node_4.Neighbors), len(cloned_4.Neighbors))
	assert.Equal(t, 2, len(cloned_4.Neighbors)) // Should have nodes 1 and 3

	// Test 9: Verify circular references are maintained correctly
	// cloned_1's neighbor (cloned_2) should have cloned_1 as a neighbor
	found_1_in_2 := false
	for _, neighbor := range cloned_2.Neighbors {
		if neighbor == cloned_1 { // Same reference, not just same value
			found_1_in_2 = true
			break
		}
	}
	assert.True(t, found_1_in_2, "Cloned graph should maintain circular references")

	// Test 10: Use BFS to verify all nodes are reachable and correctly connected
	visited := make(map[*Node]bool)
	queue := []*Node{cloned_1}
	visited[cloned_1] = true
	nodeCount := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		nodeCount++

		for _, neighbor := range current.Neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	assert.Equal(t, 4, nodeCount, "All 4 nodes should be reachable in cloned graph")
}

// Additional test for edge cases
func Test_CloneGraph_EdgeCases(t *testing.T) {
	// Test 1: Empty graph (nil input)
	var nilNode *Node
	cloned := cloneGraph(nilNode)
	assert.Nil(t, cloned, "Cloning nil should return nil")

	// Test 2: Single node with no neighbors
	single := &Node{Val: 1}
	clonedSingle := cloneGraph(single)
	assert.NotNil(t, clonedSingle)
	assert.Equal(t, 1, clonedSingle.Val)
	assert.Equal(t, 0, len(clonedSingle.Neighbors))
	assert.NotSame(t, single, clonedSingle)

	// Test 3: Two nodes with bidirectional connection
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Neighbors = []*Node{node2}
	node2.Neighbors = []*Node{node1}

	cloned1 := cloneGraph(node1)
	assert.Equal(t, 1, cloned1.Val)
	assert.Equal(t, 1, len(cloned1.Neighbors))
	cloned2 := cloned1.Neighbors[0]
	assert.Equal(t, 2, cloned2.Val)
	assert.Equal(t, 1, len(cloned2.Neighbors))
	assert.Equal(t, cloned1, cloned2.Neighbors[0], "Should maintain bidirectional reference")
}
