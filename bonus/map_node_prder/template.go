package map_node_prder

import (
	"fmt"
	"go_leetcode_journey/common"
)

type TreeNode struct {
	id       string
	subNodes *[]TreeNode
}

func preOrder(root *TreeNode) string {
	ans := root.id
	if root.subNodes != nil {
		parseSubNodePreOrder(&ans, root.subNodes)
	}
	return ans
}

func parseSubNodePreOrder(ans *string, subNodes *[]TreeNode) {
	for _, node := range *subNodes {
		*ans += node.id
		if node.subNodes != nil {
			parseSubNodePreOrder(ans, node.subNodes)
		}
	}
}

//------------------------------------------------------------------------

func postOrder(root *TreeNode) string {
	ans := ""
	if root.subNodes != nil {
		parseSubNodePostOrder(&ans, root.subNodes)
	}
	return ans + root.id
}

func parseSubNodePostOrder(ans *string, subNodes *[]TreeNode) {
	for _, node := range *subNodes {
		if node.subNodes == nil {
			*ans += node.id
		} else {
			parseSubNodePostOrder(ans, node.subNodes)
			*ans += node.id
		}
	}
}

//------------------------------------------------------------------------

type Queue struct {
	nodes []*TreeNode
}

// Enqueue adds a TreeNode to the queue
func (q *Queue) Enqueue(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

// Dequeue removes and returns the front TreeNode from the queue
func (q *Queue) Dequeue() *TreeNode {
	if len(q.nodes) == 0 {
		return nil
	}
	frontNode := q.nodes[0]
	q.nodes = q.nodes[1:] // Remove the front element
	return frontNode
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

func levelOrder(root *TreeNode) string {
	ans := ""

	queue := &Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		ans += node.id
		if node.subNodes != nil {
			for _, subNode := range *node.subNodes {
				queue.Enqueue(&subNode)
			}
		}
	}

	return ans
}

//------------------------------------------------------------------------

// Go call this func in main.go
func Go() {
	node9 := TreeNode{id: "9"}
	node8 := TreeNode{id: "8"}
	node7 := TreeNode{id: "7"}
	node6 := TreeNode{id: "6"}
	node5 := TreeNode{id: "5"}
	node4 := TreeNode{id: "4", subNodes: &[]TreeNode{node9}}
	node3 := TreeNode{id: "3", subNodes: &[]TreeNode{node7, node8}}
	node2 := TreeNode{id: "2", subNodes: &[]TreeNode{node5, node6}}
	node1 := TreeNode{id: "1", subNodes: &[]TreeNode{node3, node4}}
	node0 := TreeNode{id: "0", subNodes: &[]TreeNode{node1, node2}}

	pre := preOrder(&node0)
	post := postOrder(&node0)
	level := levelOrder(&node0)

	fmt.Println("pre: ", pre)
	common.Assert_answer(pre, "0137849256")
	fmt.Println("post: ", post)
	common.Assert_answer(post, "7839415620")
	fmt.Println("level: ", level)
	common.Assert_answer(level, "0123456789")
}
