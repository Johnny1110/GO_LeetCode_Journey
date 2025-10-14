package binary_tree_level_order_traversal

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

func (q *Queue) Size() int {
	return len(q.container)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := NewQueue()
	// init, put head into queue.
	queue.Push(root)
	level := [][]int{}

	for queue.Size() > 0 {
		currentSize := queue.Size()
		currentLayer := []int{}
		for range currentSize {
			if node, ok := queue.Pop(); ok {
				currentLayer = append(currentLayer, node.Val)
				// put next layer's node into queue.
				if node.Left != nil {
					queue.Push(node.Left)
				}
				if node.Right != nil {
					queue.Push(node.Right)
				}
			} else {
				panic("queue is empty")
			}
		}
		level = append(level, currentLayer)
	}

	return level
}
