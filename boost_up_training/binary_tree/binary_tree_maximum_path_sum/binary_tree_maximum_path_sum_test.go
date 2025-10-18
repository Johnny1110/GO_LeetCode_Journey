package binary_tree_maximum_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	root := &TreeNode{Val: 1}

	nodeA := &TreeNode{Val: 2}
	nodeB := &TreeNode{Val: 3}

	root.Left = nodeA
	root.Right = nodeB

	assert.Equal(t, 6, maxPathSum(root))
}

func Test_2(t *testing.T) {
	root := &TreeNode{Val: -10}

	nodeA := &TreeNode{Val: 9}
	nodeB := &TreeNode{Val: 20}
	nodeC := &TreeNode{Val: 15}
	nodeD := &TreeNode{Val: 7}

	root.Left = nodeA
	root.Right = nodeB

	nodeB.Left = nodeC
	nodeC.Right = nodeD

	assert.Equal(t, 42, maxPathSum(root))
	// The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
}
