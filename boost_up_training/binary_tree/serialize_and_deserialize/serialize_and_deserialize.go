package serialize_and_deserialize

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ==========================================================

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	queue := NewNodeQueue()
	queue.Push(root)

	result := []string{}

	for queue.Size() != 0 {
		node, _ := queue.Pop()

		if node == nil {
			result = append(result, "nil")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue.Push(node.Left)
			queue.Push(node.Right)
		}
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strVals := strings.Split(data, ",")

	if len(strVals) == 0 || strVals[0] == "nil" {
		return nil
	}

	rootVal, _ := strconv.Atoi(strVals[0])
	idx := 1

	root := &TreeNode{
		Val: rootVal,
	}

	queue := NewNodeQueue()
	queue.Push(root)

	for queue.Size() != 0 {
		node, _ := queue.Pop()
		if node == nil {
			continue
		}

		// handle left
		if idx < len(strVals) {
			strLeftVal := strVals[idx]
			idx++

			if strLeftVal != "nil" {
				leftVal, _ := strconv.Atoi(strLeftVal)
				node.Left = &TreeNode{
					Val: leftVal,
				}
				queue.Push(node.Left)
			}
		}

		// handle right
		if idx < len(strVals) {
			strRightVal := strVals[idx]
			idx++

			if strRightVal != "nil" {
				rightVal, _ := strconv.Atoi(strRightVal)
				node.Right = &TreeNode{
					Val: rightVal,
				}
				queue.Push(node.Right)
			}
		}
	}

	return root
}

// --------------------------------------------------------------------------

type NodeQueue struct {
	container []*TreeNode
}

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{
		container: make([]*TreeNode, 0),
	}
}

func (this *NodeQueue) Push(node *TreeNode) {
	this.container = append(this.container, node)
}

func (this *NodeQueue) Pop() (*TreeNode, bool) {
	if len(this.container) == 0 {
		return nil, false
	}
	node := this.container[0]
	this.container = this.container[1:]
	return node, true
}

func (this *NodeQueue) Size() int {
	return len(this.container)
}
