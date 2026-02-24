package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ==============================

type NodeQueue []*TreeNode

func (q NodeQueue) Len() int {
	return len(q)
}

func (q *NodeQueue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *NodeQueue) Pop() *TreeNode {
	ret := (*q)[0]
	(*q)[0] = nil
	*q = (*q)[1:]
	return ret
}

// ==============================

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	Q := NodeQueue(make([]*TreeNode, 0))
	Q.Push(root)

	for Q.Len() > 0 {
		currentLayerLen := Q.Len()
		currentLayerResult := make([]int, currentLayerLen)
		result = append(result, currentLayerResult)

		for i := range currentLayerLen {
			node := Q.Pop()
			currentLayerResult[i] = node.Val
			if node.Left != nil {
				Q.Push(node.Left)
			}
			if node.Right != nil {
				Q.Push(node.Right)
			}
		}
	}

	return result
}
