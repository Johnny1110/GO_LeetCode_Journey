package serialize_and_deserialize

import (
	"fmt"
	"strings"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	tests := []struct {
		name               string
		buildTree          func() *TreeNode
		expectedSerialized string
	}{
		{
			name: "empty tree",
			buildTree: func() *TreeNode {
				return nil
			},
			expectedSerialized: "",
		},
		{
			name: "single node",
			buildTree: func() *TreeNode {
				return &TreeNode{Val: 1}
			},
			expectedSerialized: "1,null,null",
		},
		{
			//[1,2,3,null,null,4,5]
			name: "Example 1",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				root.Right.Left = &TreeNode{Val: 4}
				root.Right.Right = &TreeNode{Val: 5}
				return root
			},
			expectedSerialized: "1,2,null,null,3,4,null,null,5,null,null",
		},
		{
			name: "left skewed tree",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Left.Left = &TreeNode{Val: 3}
				root.Left.Left.Left = &TreeNode{Val: 4}
				return root
			},
			expectedSerialized: "1,2,3,4,null,null,null,null,null",
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
			expectedSerialized: "1,null,2,null,3,null,4,null,null",
		},
		{
			name: "complete binary tree",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				root.Left.Left = &TreeNode{Val: 4}
				root.Left.Right = &TreeNode{Val: 5}
				root.Right.Left = &TreeNode{Val: 6}
				root.Right.Right = &TreeNode{Val: 7}
				return root
			},
			expectedSerialized: "1,2,4,null,null,5,null,null,3,6,null,null,7,null,null",
		},
		{
			name: "tree with negative values",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: -1}
				root.Left = &TreeNode{Val: -2}
				root.Right = &TreeNode{Val: 3}
				return root
			},
			expectedSerialized: "-1,-2,null,null,3,null,null",
		},
		{
			name: "tree with large values",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1000}
				root.Left = &TreeNode{Val: 2000}
				root.Right = &TreeNode{Val: 3000}
				return root
			},
			expectedSerialized: "1000,2000,null,null,3000,null,null",
		},
		{
			name: "unbalanced tree",
			buildTree: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 2}
				root.Right = &TreeNode{Val: 3}
				root.Left.Right = &TreeNode{Val: 4}
				root.Right.Left = &TreeNode{Val: 5}
				return root
			},
			expectedSerialized: "1,2,null,4,null,null,3,5,null,null,null",
		},
	}

	codec := Constructor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.buildTree()

			// Test serialize
			serialized := codec.serialize(root)
			if tt.expectedSerialized != "" && serialized == "" {
				t.Errorf("serialize() returned empty string, expected: %s", tt.expectedSerialized)
				return
			}

			// Test deserialize
			deserialized := codec.deserialize(serialized)

			// Test round-trip: serialize -> deserialize -> serialize should be consistent
			serializedAgain := codec.serialize(deserialized)
			if serialized != serializedAgain {
				t.Errorf("Round-trip failed: original=%s, after round-trip=%s", serialized, serializedAgain)
			}

			// Test tree structure equivalence
			if !areTreesEqual(root, deserialized) {
				t.Errorf("Deserialized tree doesn't match original")
			}
		})
	}
}

func TestSerializeDeserializeEdgeCases(t *testing.T) {
	codec := Constructor()

	t.Run("tree with zero value", func(t *testing.T) {
		root := &TreeNode{Val: 0}
		root.Left = &TreeNode{Val: -1}
		root.Right = &TreeNode{Val: 1}

		serialized := codec.serialize(root)
		deserialized := codec.deserialize(serialized)

		if !areTreesEqual(root, deserialized) {
			t.Errorf("Failed to handle tree with zero value")
		}
	})

	t.Run("deep tree", func(t *testing.T) {
		// Create a deep tree
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 10; i++ {
			current.Left = &TreeNode{Val: i}
			current = current.Left
		}

		serialized := codec.serialize(root)
		deserialized := codec.deserialize(serialized)

		if !areTreesEqual(root, deserialized) {
			t.Errorf("Failed to handle deep tree")
		}
	})

	t.Run("tree with duplicate values", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 1}
		root.Left.Left = &TreeNode{Val: 1}

		serialized := codec.serialize(root)
		deserialized := codec.deserialize(serialized)

		if !areTreesEqual(root, deserialized) {
			t.Errorf("Failed to handle tree with duplicate values")
		}
	})
}

func TestDeserializeEdgeCases(t *testing.T) {
	codec := Constructor()

	t.Run("empty string", func(t *testing.T) {
		result := codec.deserialize("")
		if result != nil {
			t.Errorf("deserialize(\"\") should return nil, got %v", result)
		}
	})

	t.Run("malformed input handling", func(t *testing.T) {
		// Test with various potentially problematic inputs
		testCases := []string{
			"null",
			"1",
			"1,null",
			"1,2",
		}

		for _, tc := range testCases {
			// Should not panic
			result := codec.deserialize(tc)
			// Re-serialize to test consistency
			serialized := codec.serialize(result)
			// Should be able to deserialize again
			codec.deserialize(serialized)
		}
	})
}

func TestSerializeFormat(t *testing.T) {
	codec := Constructor()

	t.Run("check serialization format", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}

		serialized := codec.serialize(root)

		// Should contain commas for separation
		if !strings.Contains(serialized, ",") && serialized != "" {
			t.Errorf("Serialized format should use commas for separation, got: %s", serialized)
		}

		// Should handle null nodes
		if !strings.Contains(serialized, "null") && serialized != "" {
			t.Errorf("Serialized format should represent null nodes, got: %s", serialized)
		}
	})
}

func TestLargeTree(t *testing.T) {
	codec := Constructor()

	// Build a larger tree for stress testing
	root := buildLargeTree(100, 1)

	serialized := codec.serialize(root)
	deserialized := codec.deserialize(serialized)

	if !areTreesEqual(root, deserialized) {
		t.Errorf("Failed to handle large tree correctly")
	}
}

func BenchmarkSerialize(b *testing.B) {
	codec := Constructor()
	root := buildLargeTree(50, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.serialize(root)
	}
}

func BenchmarkDeserialize(b *testing.B) {
	codec := Constructor()
	root := buildLargeTree(50, 1)
	serialized := codec.serialize(root)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.deserialize(serialized)
	}
}

func BenchmarkSerializeDeserialize(b *testing.B) {
	codec := Constructor()
	root := buildLargeTree(30, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serialized := codec.serialize(root)
		codec.deserialize(serialized)
	}
}

// Helper functions

func areTreesEqual(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return areTreesEqual(root1.Left, root2.Left) && areTreesEqual(root1.Right, root2.Right)
}

func buildLargeTree(maxVal, val int) *TreeNode {
	if val > maxVal {
		return nil
	}
	root := &TreeNode{Val: val}
	root.Left = buildLargeTree(maxVal, val*2)
	root.Right = buildLargeTree(maxVal, val*2+1)
	return root
}

func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var result []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func TestTreeEquality(t *testing.T) {
	t.Run("helper function test", func(t *testing.T) {
		// Test the helper function itself
		root1 := &TreeNode{Val: 1}
		root1.Left = &TreeNode{Val: 2}
		root1.Right = &TreeNode{Val: 3}

		root2 := &TreeNode{Val: 1}
		root2.Left = &TreeNode{Val: 2}
		root2.Right = &TreeNode{Val: 3}

		if !areTreesEqual(root1, root2) {
			t.Errorf("areTreesEqual should return true for identical trees")
		}

		root3 := &TreeNode{Val: 1}
		root3.Left = &TreeNode{Val: 2}
		root3.Right = &TreeNode{Val: 4} // Different value

		if areTreesEqual(root1, root3) {
			t.Errorf("areTreesEqual should return false for different trees")
		}
	})
}

func Test_serialize(t *testing.T) {
	codec := Constructor()

	buildTree := func() *TreeNode {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Right.Left = &TreeNode{Val: 4}
		root.Right.Right = &TreeNode{Val: 5}
		return root
	}

	root := buildTree

	res := codec.serialize(root())
	fmt.Println("serialize : @.@")
	fmt.Println(res)
}
