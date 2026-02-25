package lowest_common_ancestor

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name      string
		buildTree func() (*TreeNode, *TreeNode, *TreeNode)
		expected  int
	}{
		{
			name: "Example 1",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				root.Left.Left = &TreeNode{Val: 6}
				root.Left.Right = &TreeNode{Val: 2}
				root.Right.Left = &TreeNode{Val: 0}
				root.Right.Right = &TreeNode{Val: 8}
				root.Left.Right.Left = &TreeNode{Val: 7}
				root.Left.Right.Right = &TreeNode{Val: 4}

				p := root.Left  // 5
				q := root.Right // 1
				return root, p, q
			},
			expected: 3,
		},
		{
			name: "Example 2",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				root.Left.Left = &TreeNode{Val: 6}
				root.Left.Right = &TreeNode{Val: 2}
				root.Right.Left = &TreeNode{Val: 0}
				root.Right.Right = &TreeNode{Val: 8}
				root.Left.Right.Left = &TreeNode{Val: 7}
				root.Left.Right.Right = &TreeNode{Val: 4}

				p := root.Left             // 5
				q := root.Left.Right.Right // 4
				return root, p, q
			},
			expected: 5,
		},
		{
			name: "LCA of two adjacent nodes",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}

				p := root.Left  // 2
				q := root.Right // 3
				return root, p, q
			},
			expected: 1,
		},
		{
			name: "LCA where one node is ancestor of another",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Left.Left = &TreeNode{Val: 4}
				root.Left.Right = &TreeNode{Val: 5}

				p := root.Left      // 2
				q := root.Left.Left // 4
				return root, p, q
			},
			expected: 2,
		},
		{
			name: "LCA in left subtree",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				root.Left.Left = &TreeNode{Val: 4}
				root.Left.Right = &TreeNode{Val: 5}
				root.Left.Left.Left = &TreeNode{Val: 8}
				root.Left.Right.Right = &TreeNode{Val: 9}

				p := root.Left.Left.Left   // 8
				q := root.Left.Right.Right // 9
				return root, p, q
			},
			expected: 2,
		},
		{
			name: "LCA in right subtree",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				root.Right.Left = &TreeNode{Val: 6}
				root.Right.Right = &TreeNode{Val: 7}
				root.Right.Left.Left = &TreeNode{Val: 12}
				root.Right.Right.Right = &TreeNode{Val: 13}

				p := root.Right.Left.Left   // 12
				q := root.Right.Right.Right // 13
				return root, p, q
			},
			expected: 3,
		},
		{
			name: "LCA of leaf nodes in different subtrees",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 10}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 15}
				root.Left.Left = &TreeNode{Val: 3}
				root.Left.Right = &TreeNode{Val: 7}
				root.Right.Left = &TreeNode{Val: 12}
				root.Right.Right = &TreeNode{Val: 18}

				p := root.Left.Left   // 3
				q := root.Right.Right // 18
				return root, p, q
			},
			expected: 10,
		},
		{
			name: "LCA where p is root",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				root.Left.Left = &TreeNode{Val: 4}

				p := root           // 1
				q := root.Left.Left // 4
				return root, p, q
			},
			expected: 1,
		},
		{
			name: "LCA in deep tree",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Left.Left = &TreeNode{Val: 4}
				root.Left.Left.Left = &TreeNode{Val: 8}
				root.Left.Left.Left.Left = &TreeNode{Val: 16}
				root.Left.Left.Right = &TreeNode{Val: 9}
				root.Left.Left.Right.Right = &TreeNode{Val: 17}

				p := root.Left.Left.Left.Left   // 16
				q := root.Left.Left.Right.Right // 17
				return root, p, q
			},
			expected: 8,
		},
		{
			name: "LCA with negative values",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: -1}
				root.Left = &TreeNode{Val: -2}
				root.Right = &TreeNode{Val: -3}
				root.Left.Left = &TreeNode{Val: -4}
				root.Right.Right = &TreeNode{Val: -5}

				p := root.Left.Left   // -4
				q := root.Right.Right // -5
				return root, p, q
			},
			expected: -1,
		},
		{
			name: "LCA in unbalanced tree",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				root := &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 2}
				root.Right.Right = &TreeNode{Val: 3}
				root.Right.Right.Right = &TreeNode{Val: 4}
				root.Right.Right.Right.Left = &TreeNode{Val: 5}
				root.Right.Right.Right.Right = &TreeNode{Val: 6}

				p := root.Right.Right.Right.Left  // 5
				q := root.Right.Right.Right.Right // 6
				return root, p, q
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, p, q := tt.buildTree()
			result := lowestCommonAncestor(root, p, q)
			if result == nil {
				t.Errorf("lowestCommonAncestor() returned nil, expected node with value %d", tt.expected)
			} else if result.Val != tt.expected {
				t.Errorf("lowestCommonAncestor() = %d, expected %d", result.Val, tt.expected)
			}
		})
	}
}

func TestLowestCommonAncestorEdgeCases(t *testing.T) {
	t.Run("simple two-node tree", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}

		p := root
		q := root.Left
		result := lowestCommonAncestor(root, p, q)
		if result == nil || result.Val != 1 {
			t.Errorf("Expected LCA to be 1, got %v", result)
		}
	})

	t.Run("three-node linear tree", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}

		p := root.Left
		q := root.Left.Left
		result := lowestCommonAncestor(root, p, q)
		if result == nil || result.Val != 2 {
			t.Errorf("Expected LCA to be 2, got %v", result)
		}
	})

	t.Run("large values", func(t *testing.T) {
		root := &TreeNode{Val: 1000000}
		root.Left = &TreeNode{Val: 999999}
		root.Right = &TreeNode{Val: 1000001}

		p := root.Left
		q := root.Right
		result := lowestCommonAncestor(root, p, q)
		if result == nil || result.Val != 1000000 {
			t.Errorf("Expected LCA to be 1000000, got %v", result)
		}
	})
}

func TestLowestCommonAncestorComplex(t *testing.T) {
	t.Run("complex tree with multiple levels", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 5}
		root.Right.Left = &TreeNode{Val: 6}
		root.Right.Right = &TreeNode{Val: 7}
		root.Left.Left.Left = &TreeNode{Val: 8}
		root.Left.Left.Right = &TreeNode{Val: 9}
		root.Left.Right.Left = &TreeNode{Val: 10}
		root.Left.Right.Right = &TreeNode{Val: 11}
		root.Right.Left.Left = &TreeNode{Val: 12}
		root.Right.Left.Right = &TreeNode{Val: 13}
		root.Right.Right.Left = &TreeNode{Val: 14}
		root.Right.Right.Right = &TreeNode{Val: 15}

		testCases := []struct {
			pVal, qVal, expectedLCA int
		}{
			{8, 9, 4},   // Both in left subtree of 4
			{10, 11, 5}, // Both in right subtree of 5
			{8, 10, 2},  // Different subtrees of 2
			{12, 15, 3}, // Different subtrees of 3
			{8, 15, 1},  // Across left and right subtrees
			{4, 13, 1},  // Across left and right subtrees
		}

		nodeMap := make(map[int]*TreeNode)
		var collectNodes func(*TreeNode)
		collectNodes = func(node *TreeNode) {
			if node != nil {
				nodeMap[node.Val] = node
				collectNodes(node.Left)
				collectNodes(node.Right)
			}
		}
		collectNodes(root)

		for _, tc := range testCases {
			p := nodeMap[tc.pVal]
			q := nodeMap[tc.qVal]
			result := lowestCommonAncestor(root, p, q)
			if result == nil || result.Val != tc.expectedLCA {
				t.Errorf("LCA of %d and %d: expected %d, got %v",
					tc.pVal, tc.qVal, tc.expectedLCA, result)
			}
		}
	})
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	p := root.Left
	q := root.Right

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lowestCommonAncestor(root, p, q)
	}
}

func buildDeepTree(depth int, val int) *TreeNode {
	if depth <= 0 {
		return nil
	}
	node := &TreeNode{Val: val}
	node.Left = buildDeepTree(depth-1, val*2)
	node.Right = buildDeepTree(depth-1, val*2+1)
	return node
}

func BenchmarkLowestCommonAncestorDeep(b *testing.B) {
	root := buildDeepTree(12, 1)

	var findNode func(*TreeNode, int) *TreeNode
	findNode = func(node *TreeNode, val int) *TreeNode {
		if node == nil || node.Val == val {
			return node
		}
		if left := findNode(node.Left, val); left != nil {
			return left
		}
		return findNode(node.Right, val)
	}

	p := findNode(root, 1024)
	q := findNode(root, 1025)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lowestCommonAncestor(root, p, q)
	}
}
