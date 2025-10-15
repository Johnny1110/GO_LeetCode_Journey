package lowest_common_ancestor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: root = [2,1,3]
	// Output: true

	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}

	node3.Left = node5
	node3.Right = node1

	node5.Left = node6
	node5.Right = node2

	node1.Left = node0
	node1.Right = node8

	node2.Left = node7
	node2.Right = node4

	assert.Equal(t, node3, lowestCommonAncestor(node3, node5, node1))
	assert.Equal(t, node2, lowestCommonAncestor(node3, node7, node4))
	assert.Equal(t, node5, lowestCommonAncestor(node3, node6, node4))
	assert.Equal(t, node5, lowestCommonAncestor(node3, node7, node6))
	assert.Equal(t, node3, lowestCommonAncestor(node3, node7, node8))
}
