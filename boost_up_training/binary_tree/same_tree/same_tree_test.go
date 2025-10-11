package same_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 2}
	nodeA3 := &TreeNode{Val: 3}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 2}
	nodeB3 := &TreeNode{Val: 3}

	nodeB1.Left = nodeB2
	nodeB1.Right = nodeB3

	assert.True(t, isSameTree(nodeA1, nodeB1))
}

func Test_2(t *testing.T) {
	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 2}
	nodeA3 := &TreeNode{Val: 3}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 2}
	nodeB3 := &TreeNode{Val: 1}

	nodeB1.Left = nodeB2
	nodeB1.Right = nodeB3

	assert.False(t, isSameTree(nodeA1, nodeB1))
}

func Test_3(t *testing.T) {
	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 2}
	nodeA3 := &TreeNode{Val: 3}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 2}

	nodeB1.Left = nodeB2

	assert.False(t, isSameTree(nodeA1, nodeB1))
}

func Test_4(t *testing.T) {
	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 2}
	nodeA3 := &TreeNode{Val: 3}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 3}
	nodeB3 := &TreeNode{Val: 2}

	nodeB1.Left = nodeB2
	nodeB1.Right = nodeB3

	assert.False(t, isSameTree(nodeA1, nodeB1))
}
