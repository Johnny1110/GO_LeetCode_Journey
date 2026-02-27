package binary_tree_max_imumpathsum

import (
	"math"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name      string
		buildTree func() *TreeNode
		expected  int
	}{
		{
			name: "single node positive",
			buildTree: func() *TreeNode {
				return &TreeNode{Val: 1}
			},
			expected: 1,
		},
		{
			name: "single node negative",
			buildTree: func() *TreeNode {
				return &TreeNode{Val: -3}
			},
			expected: -3,
		},
		{
			name: "Example 1",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				return root
			},
			expected: 6, // 2 + 1 + 3
		},
		{
			name: "Example 2 - [-10,9,20,null,null,15,7]",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: -10}
				root.Left = &TreeNode{Val: 9}
				root.Right = &TreeNode{Val: 20}
				root.Right.Left = &TreeNode{Val: 15}
				root.Right.Right = &TreeNode{Val: 7}
				return root
			},
			expected: 42, // 15 + 20 + 7
		},
		{
			name: "path through root",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 2}
				root.Left = &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 3}
				return root
			},
			expected: 6, // 1 + 2 + 3
		},
		{
			name: "path not through root",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: -3}
				root.Left = &TreeNode{Val: 3}
				root.Right = &TreeNode{Val: -2}
				root.Left.Left = &TreeNode{Val: -1}
				root.Left.Right = &TreeNode{Val: 2}
				return root
			},
			expected: 5, // 3 + 2 (left subtree path)
		},
		{
			name: "all negative values",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: -3}
				root.Left = &TreeNode{Val: -5}
				root.Right = &TreeNode{Val: -1}
				return root
			},
			expected: -1, // single node with highest value
		},
		{
			name: "left skewed tree",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 5}
				root.Left = &TreeNode{Val: 4}
				root.Left.Left = &TreeNode{Val: 11}
				root.Left.Left.Left = &TreeNode{Val: 7}
				root.Left.Left.Right = &TreeNode{Val: 2}
				return root
			},
			expected: 20, // 7 + 11 + 2
		},
		{
			name: "right skewed tree",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 2}
				root.Right.Right = &TreeNode{Val: 3}
				root.Right.Right.Right = &TreeNode{Val: 4}
				return root
			},
			expected: 10, // 1 + 2 + 3 + 4
		},
		{
			name: "complex tree with multiple possible paths",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 5}
				root.Left = &TreeNode{Val: 4}
				root.Right = &TreeNode{Val: 8}
				root.Left.Left = &TreeNode{Val: 11}
				root.Left.Left.Left = &TreeNode{Val: 7}
				root.Left.Left.Right = &TreeNode{Val: 2}
				root.Right.Left = &TreeNode{Val: 13}
				root.Right.Right = &TreeNode{Val: 4}
				root.Right.Right.Right = &TreeNode{Val: 1}
				return root
			},
			expected: 48, // 7 + 11 + 4 + 5 + 8 + 13
		},
		{
			name: "tree with zero values",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 0}
				root.Left = &TreeNode{Val: 0}
				root.Right = &TreeNode{Val: 0}
				return root
			},
			expected: 0, // 0 + 0 + 0
		},
		{
			name: "mixed positive and negative",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: -2}
				root.Right = &TreeNode{Val: 3}
				root.Left.Left = &TreeNode{Val: 1}
				root.Left.Right = &TreeNode{Val: 3}
				root.Right.Left = &TreeNode{Val: -2}
				root.Right.Right = &TreeNode{Val: -1}
				return root
			},
			expected: 4, // 1 + 3 (right side)
		},
		{
			name: "large positive values",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1000}
				root.Left = &TreeNode{Val: 2000}
				root.Right = &TreeNode{Val: 3000}
				return root
			},
			expected: 6000, // 2000 + 1000 + 3000
		},
		{
			name: "deep tree with best path at leaves",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Left.Left = &TreeNode{Val: 4}
				root.Left.Right = &TreeNode{Val: 5}
				root.Left.Left.Left = &TreeNode{Val: 100}
				root.Left.Left.Right = &TreeNode{Val: 200}
				return root
			},
			expected: 306, // 100 + 4 + 200
		},
		{
			name: "single path down",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 2}
				root.Left = &TreeNode{Val: 1}
				root.Left.Left = &TreeNode{Val: 1}
				return root
			},
			expected: 4, // 1 + 1 + 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.buildTree()
			result := maxPathSum(root)
			if result != tt.expected {
				t.Errorf("maxPathSum() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMaxPathSumEdgeCases(t *testing.T) {
	t.Run("minimum integer value", func(t *testing.T) {
		root := &TreeNode{Val: math.MinInt32}
		expected := math.MinInt32
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})

	t.Run("maximum integer value", func(t *testing.T) {
		root := &TreeNode{Val: math.MaxInt32}
		expected := math.MaxInt32
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})

	t.Run("tree with extreme negative and positive", func(t *testing.T) {
		root := &TreeNode{Val: -1000}
		root.Left = &TreeNode{Val: 2000}
		root.Right = &TreeNode{Val: 3000}
		expected := 4000 // 2000 + (-1000) + 3000
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})
}

func TestMaxPathSumComplexScenarios(t *testing.T) {
	t.Run("unbalanced tree favoring left", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}
		root.Left.Left.Left.Left = &TreeNode{Val: 5}
		expected := 15 // 5 + 4 + 3 + 2 + 1
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})

	t.Run("unbalanced tree favoring right", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 3}
		root.Right.Right.Right = &TreeNode{Val: 4}
		root.Right.Right.Right.Right = &TreeNode{Val: 5}
		expected := 15 // 1 + 2 + 3 + 4 + 5
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})

	t.Run("diamond shaped tree", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 5}
		root.Right.Left = &TreeNode{Val: 6}
		root.Right.Right = &TreeNode{Val: 7}
		// Best path could be 4 + 2 + 1 + 3 + 7 = 17
		// or other combinations
		result := maxPathSum(root)
		if result < 17 { // At least this much should be achievable
			t.Errorf("maxPathSum() = %d, expected at least 17", result)
		}
	})

	t.Run("alternating positive negative", func(t *testing.T) {
		root := &TreeNode{Val: 10}
		root.Left = &TreeNode{Val: -5}
		root.Right = &TreeNode{Val: -5}
		root.Left.Left = &TreeNode{Val: 15}
		root.Left.Right = &TreeNode{Val: 15}
		root.Right.Left = &TreeNode{Val: 15}
		root.Right.Right = &TreeNode{Val: 15}
		// Best path should be one of the leaf nodes: 15
		result := maxPathSum(root)
		if result < 15 { // At least single leaf value
			t.Errorf("maxPathSum() = %d, expected at least 15", result)
		}
	})
}

func TestMaxPathSumSpecialTrees(t *testing.T) {
	t.Run("perfect binary tree depth 3", func(t *testing.T) {
		// Build perfect binary tree with all 1s
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 1}
		root.Left.Left = &TreeNode{Val: 1}
		root.Left.Right = &TreeNode{Val: 1}
		root.Right.Left = &TreeNode{Val: 1}
		root.Right.Right = &TreeNode{Val: 1}
		// Best path would go through multiple nodes
		result := maxPathSum(root)
		if result < 4 { // At least 4 nodes in some path
			t.Errorf("maxPathSum() = %d, expected at least 4", result)
		}
	})

	t.Run("tree with only left children", func(t *testing.T) {
		root := &TreeNode{Val: 10}
		current := root
		sum := 10
		for i := 9; i >= 1; i-- {
			current.Left = &TreeNode{Val: i}
			current = current.Left
			sum += i
		}
		expected := sum // Sum of all nodes in the chain
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})

	t.Run("tree with only right children", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		current := root
		sum := 1
		for i := 2; i <= 10; i++ {
			current.Right = &TreeNode{Val: i}
			current = current.Right
			sum += i
		}
		expected := sum // Sum of all nodes in the chain
		result := maxPathSum(root)
		if result != expected {
			t.Errorf("maxPathSum() = %d, expected %d", result, expected)
		}
	})
}

func BenchmarkMaxPathSum(b *testing.B) {
	// Build a moderately complex tree for benchmarking
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxPathSum(root)
	}
}

func BenchmarkMaxPathSumDeepTree(b *testing.B) {
	// Build a deep tree
	root := buildDeepTree(15, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxPathSum(root)
	}
}

func BenchmarkMaxPathSumWideTree(b *testing.B) {
	// Build a wide tree
	root := buildWideTree(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxPathSum(root)
	}
}

// Helper functions for benchmarks

func buildDeepTree(depth, val int) *TreeNode {
	if depth <= 0 {
		return nil
	}
	root := &TreeNode{Val: val}
	root.Left = buildDeepTree(depth-1, val+1)
	root.Right = buildDeepTree(depth-1, val+2)
	return root
}

func buildWideTree(numNodes int) *TreeNode {
	if numNodes <= 0 {
		return nil
	}
	
	root := &TreeNode{Val: 1}
	queue := []*TreeNode{root}
	val := 2
	
	for len(queue) > 0 && val <= numNodes {
		node := queue[0]
		queue = queue[1:]
		
		if val <= numNodes {
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
			val++
		}
		
		if val <= numNodes {
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
			val++
		}
	}
	
	return root
}