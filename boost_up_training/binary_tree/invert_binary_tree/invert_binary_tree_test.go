package invert_binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	root := &TreeNode{Val: 4}

	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 1}
	nodeA3 := &TreeNode{Val: 3}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 7}
	nodeB2 := &TreeNode{Val: 6}
	nodeB3 := &TreeNode{Val: 9}

	nodeB1.Left = nodeB2
	nodeB1.Right = nodeB3

	root.Left = nodeA1
	root.Right = nodeB1

	// expect
	expectTest1 := mockTestExpect_1()
	resultTest1 := invertTree(root)

	debug(resultTest1)

	isSame := isSameTree(expectTest1, resultTest1)

	assert.True(t, isSame)
}

func mockTestExpect_1() *TreeNode {
	root := &TreeNode{Val: 4}

	nodeA1 := &TreeNode{Val: 7}
	nodeA2 := &TreeNode{Val: 9}
	nodeA3 := &TreeNode{Val: 6}

	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 2}
	nodeB2 := &TreeNode{Val: 3}
	nodeB3 := &TreeNode{Val: 1}

	nodeB1.Left = nodeB2
	nodeB1.Right = nodeB3

	root.Left = nodeA1
	root.Right = nodeB1

	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func debug(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println("node:", root.Val)
	fmt.Println("left:", root.Left)
	fmt.Println("right:", root.Right)
	fmt.Println("---------------------------------------------------")

	debug(root.Left)
	debug(root.Right)
}
