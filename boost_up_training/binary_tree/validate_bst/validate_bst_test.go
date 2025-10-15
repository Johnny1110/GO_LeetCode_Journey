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

func Test_5(t *testing.T) {
	// [32,26,47,19,null,null,56,null,27]
	// false

	node32 := &TreeNode{Val: 32}
	node26 := &TreeNode{Val: 26}
	node47 := &TreeNode{Val: 47}
	node19 := &TreeNode{Val: 19}
	node56 := &TreeNode{Val: 56}
	node27 := &TreeNode{Val: 27}

	node32.Left = node26
	node32.Right = node47

	node26.Left = node19

	node47.Right = node56

	node19.Right = node27

	// Explanation: The root node's value is 5 but its right child's value is 4.

	assert.False(t, isValidBST(node32))
}

func Test_6(t *testing.T) {
	// [3,1,5,0,2,4,6]
	// true

	node3 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 5}
	node0 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}

	node3.Left = node1
	node3.Right = node5

	node1.Left = node0
	node1.Right = node2

	node5.Left = node4
	node5.Right = node6

	assert.True(t, isValidBST(node3))
}
