package binary_tree_level_order_traversal

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	container []*TreeNode
}

func NewQueue() *Queue {
	return &Queue{
		container: make([]*TreeNode, 0),
	}
}

func (q *Queue) Push(node *TreeNode) {
	q.container = append(q.container, node)
}

func (q *Queue) Pop() (*TreeNode, bool) {
	if len(q.container) == 0 {
		return nil, false
	}
	node := q.container[0]
	q.container = q.container[1:]
	return node, true
}

func levelOrder(root *TreeNode) [][]int {
	queue := NewQueue()
	// init, put head into queue.
	queue.Push(root)
	return perform(queue)
}

func perform(queue *Queue) [][]int {
	// pop node
	node, ok := queue.Pop()
	if !ok {
		// reach the end.
		return nil
	}

	// push next 2 node into queue.
	if node.Left != nil {
		queue.Push(node.Left)
	}
	if node.Right != nil {
		queue.Push(node.Right)
	}

	fmt.Println("found node:", node.Val)

	// continue.
	return perform(queue)
}
