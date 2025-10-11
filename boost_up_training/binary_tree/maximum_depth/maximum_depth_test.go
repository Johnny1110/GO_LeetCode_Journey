package maximum_depth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: 3

	node3 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}

	// connect node
	node3.Left = node9
	node3.Right = node20
	node20.Left = node15
	node20.Right = node7

	assert.Equal(t, 3, maxDepth(node3))
}

func Test_2(t *testing.T) {
	// Input: root = [1, null, 2]
	// Output: 2

	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}

	// connect node
	node1.Right = node2

	assert.Equal(t, 2, maxDepth(node1))
}

func Test_3(t *testing.T) {
	// Input: root = [1, null, 2]
	// Output: 2

	node1 := &TreeNode{Val: 1}

	assert.Equal(t, 1, maxDepth(node1))
}
