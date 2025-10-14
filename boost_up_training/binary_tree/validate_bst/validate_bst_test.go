package validate_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: root = [2,1,3]
	// Output: true

	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 1}
	nodeA3 := &TreeNode{Val: 3}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	assert.True(t, isValidBST(nodeA1))
}

func Test_2(t *testing.T) {
	// Input: root = [5,1,4,null,null,3,6]
	// Output: false

	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}

	node5.Left = node1
	node5.Right = node4
	node4.Left = node3
	node4.Right = node6

	assert.False(t, isValidBST(node5))
}
