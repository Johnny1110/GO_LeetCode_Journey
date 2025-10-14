package binary_tree_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: [[3],[9,20],[15,7]]

	nodeA1 := &TreeNode{Val: 3}
	nodeA2 := &TreeNode{Val: 9}
	nodeA3 := &TreeNode{Val: 20}
	nodeA4 := &TreeNode{Val: 15}
	nodeA5 := &TreeNode{Val: 7}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeA3.Left = nodeA4
	nodeA3.Right = nodeA5

	result := levelOrder(nodeA1)
	expect := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}

	assert.Equal(t, expect, result)
}

func Test_2(t *testing.T) {
	// Input: root = [1]
	// Output: [[1]]

	nodeA1 := &TreeNode{Val: 1}

	result := levelOrder(nodeA1)
	expect := [][]int{
		{1},
	}

	assert.Equal(t, expect, result)
}

func Test_3(t *testing.T) {
	// Input: root = []
	// Output: []

	result := levelOrder(nil)
	expect := [][]int{}

	assert.Equal(t, expect, result)
}
