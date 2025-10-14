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

	// Explanation: The root node's value is 5 but its right child's value is 4.

	assert.False(t, isValidBST(node5))
}

func Test_3(t *testing.T) {
	// Input: root = [2,2,2]
	// Output: false

	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 2}
	nodeA3 := &TreeNode{Val: 2}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	assert.False(t, isValidBST(nodeA1))
}

func Test_4(t *testing.T) {
	// [5,4,6,null,null,3,7]
	// false

	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 3}
	node7 := &TreeNode{Val: 7}

	node5.Left = node4
	node5.Right = node6
	node6.Left = node3
	node6.Right = node7

	// Explanation: The root node's value is 5 but its right child's value is 4.

	assert.False(t, isValidBST(node5))

}
