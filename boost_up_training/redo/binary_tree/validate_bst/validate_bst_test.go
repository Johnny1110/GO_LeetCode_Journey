package validate_bst

import (
	"math"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			expected: false,
		},
		{
			name: "valid BST with larger tree",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Val:   15,
					Left:  &TreeNode{Val: 12},
					Right: &TreeNode{Val: 18},
				},
			},
			expected: true,
		},
		{
			name: "invalid BST - left subtree violation",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 15},
				},
				Right: &TreeNode{Val: 20},
			},
			expected: false,
		},
		{
			name: "invalid BST - right subtree violation",
			root: &TreeNode{
				Val:  10,
				Left: &TreeNode{Val: 5},
				Right: &TreeNode{
					Val:   15,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 20},
				},
			},
			expected: false,
		},
		{
			name: "left skewed valid BST",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			expected: true,
		},
		{
			name: "right skewed valid BST",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: true,
		},
		{
			name: "duplicate values - left child equal to parent",
			root: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name: "duplicate values - right child equal to parent",
			root: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name: "edge case with minimum integer",
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: math.MinInt32},
				Right: &TreeNode{Val: 1},
			},
			expected: true,
		},
		{
			name: "edge case with maximum integer",
			root: &TreeNode{
				Val:   math.MaxInt32,
				Left:  &TreeNode{Val: math.MaxInt32 - 1},
				Right: nil,
			},
			expected: true,
		},
		{
			name: "complex invalid BST",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 6},
					},
				},
				Right: &TreeNode{Val: 6},
			},
			expected: false,
		},
		{
			name: "negative values valid BST",
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -5},
				Right: &TreeNode{Val: 5},
			},
			expected: true,
		},
		{
			name: "negative values invalid BST",
			root: &TreeNode{
				Val: -10,
				Left: &TreeNode{
					Val:   -15,
					Left:  &TreeNode{Val: -20},
					Right: &TreeNode{Val: -5},
				},
				Right: &TreeNode{Val: -5},
			},
			expected: false,
		},
		{
			name: "single left child",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "single right child",
			root: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 7},
			},
			expected: true,
		},
		{
			name: "invalid single left child",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 7},
			},
			expected: false,
		},
		{
			name: "invalid single right child",
			root: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 3},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidBST(tt.root)
			if result != tt.expected {
				t.Errorf("isValidBST() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsValidBSTBoundaryValues(t *testing.T) {
	t.Run("tree with only minimum value", func(t *testing.T) {
		root := &TreeNode{Val: math.MinInt32}
		expected := true
		result := isValidBST(root)
		if result != expected {
			t.Errorf("isValidBST() = %v, expected %v", result, expected)
		}
	})

	t.Run("tree with only maximum value", func(t *testing.T) {
		root := &TreeNode{Val: math.MaxInt32}
		expected := true
		result := isValidBST(root)
		if result != expected {
			t.Errorf("isValidBST() = %v, expected %v", result, expected)
		}
	})

	t.Run("tree with extreme values valid BST", func(t *testing.T) {
		root := &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: math.MinInt32},
			Right: &TreeNode{Val: math.MaxInt32},
		}
		expected := true
		result := isValidBST(root)
		if result != expected {
			t.Errorf("isValidBST() = %v, expected %v", result, expected)
		}
	})
}

func TestIsValidBSTComplexCases(t *testing.T) {
	t.Run("deep valid BST", func(t *testing.T) {
		root := &TreeNode{
			Val: 50,
			Left: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 10},
					Right: &TreeNode{Val: 25},
				},
				Right: &TreeNode{
					Val:   40,
					Left:  &TreeNode{Val: 35},
					Right: &TreeNode{Val: 45},
				},
			},
			Right: &TreeNode{
				Val: 70,
				Left: &TreeNode{
					Val:   60,
					Left:  &TreeNode{Val: 55},
					Right: &TreeNode{Val: 65},
				},
				Right: &TreeNode{
					Val:   80,
					Left:  &TreeNode{Val: 75},
					Right: &TreeNode{Val: 85},
				},
			},
		}
		expected := true
		result := isValidBST(root)
		if result != expected {
			t.Errorf("isValidBST() = %v, expected %v", result, expected)
		}
	})

	t.Run("subtly invalid BST - violation deep in tree", func(t *testing.T) {
		root := &TreeNode{
			Val: 50,
			Left: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 10},
					Right: &TreeNode{Val: 25},
				},
				Right: &TreeNode{
					Val:   40,
					Left:  &TreeNode{Val: 35},
					Right: &TreeNode{Val: 60},
				},
			},
			Right: &TreeNode{
				Val:   70,
				Left:  &TreeNode{Val: 65},
				Right: &TreeNode{Val: 80},
			},
		}
		expected := false
		result := isValidBST(root)
		if result != expected {
			t.Errorf("isValidBST() = %v, expected %v", result, expected)
		}
	})
}

func BenchmarkIsValidBST(b *testing.B) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 18},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValidBST(root)
	}
}

func buildValidBST(depth int, minVal, maxVal int) *TreeNode {
	if depth <= 0 || minVal >= maxVal {
		return nil
	}
	midVal := minVal + (maxVal-minVal)/2
	return &TreeNode{
		Val:   midVal,
		Left:  buildValidBST(depth-1, minVal, midVal),
		Right: buildValidBST(depth-1, midVal+1, maxVal),
	}
}

func BenchmarkIsValidBSTLargeTree(b *testing.B) {
	root := buildValidBST(15, 1, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValidBST(root)
	}
}
