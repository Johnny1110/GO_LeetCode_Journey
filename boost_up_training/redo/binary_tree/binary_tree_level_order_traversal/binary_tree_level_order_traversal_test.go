package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "two levels",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20},
			},
			expected: [][]int{
				{3},
				{9, 20},
			},
		},
		{
			name: "three levels - example from leetcode",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{
				{1},
				{2, 3},
				{4, 5, 6, 7},
			},
		},
		{
			name: "incomplete tree with missing nodes",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{Val: 5},
				},
			},
			expected: [][]int{
				{1},
				{2, 3},
				{4, 5},
			},
		},
		{
			name: "tree with negative values",
			root: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: -15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{
				{-10},
				{9, 20},
				{-15, 7},
			},
		},
		{
			name: "large values",
			root: &TreeNode{
				Val:  1000,
				Left: &TreeNode{Val: 2000},
				Right: &TreeNode{Val: 3000},
			},
			expected: [][]int{
				{1000},
				{2000, 3000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrderEdgeCases(t *testing.T) {
	t.Run("single node with zero value", func(t *testing.T) {
		root := &TreeNode{Val: 0}
		expected := [][]int{{0}}
		result := levelOrder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("levelOrder() = %v, expected %v", result, expected)
		}
	})

	t.Run("tree with duplicate values", func(t *testing.T) {
		root := &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 1},
			},
		}
		expected := [][]int{
			{1},
			{1, 1},
			{1, 1},
		}
		result := levelOrder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("levelOrder() = %v, expected %v", result, expected)
		}
	})
}

func BenchmarkLevelOrder(b *testing.B) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrder(root)
	}
}

func buildLargeTree(depth int, val int) *TreeNode {
	if depth <= 0 {
		return nil
	}
	return &TreeNode{
		Val:   val,
		Left:  buildLargeTree(depth-1, val*2),
		Right: buildLargeTree(depth-1, val*2+1),
	}
}

func BenchmarkLevelOrderLargeTree(b *testing.B) {
	root := buildLargeTree(10, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrder(root)
	}
}