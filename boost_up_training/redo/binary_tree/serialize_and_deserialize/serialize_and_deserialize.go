package serialize_and_deserialize

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	Q := NodeQueue(make([]*TreeNode, 0))
	Q.Push(root)

	strVals := []string{}

	for Q.Len() > 0 {

		levelLen := Q.Len()
		for range levelLen {
			node := Q.Pop()
			if node != nil {
				strVals = append(strVals, strconv.Itoa(node.Val))
				Q.Push(node.Left)
				Q.Push(node.Right)
			} else {
				strVals = append(strVals, "null")
			}

		}
	}

	return strings.Join(strVals, ",")

}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	strVal := strings.Split(data, ",")
	Q := NodeQueue(make([]*TreeNode, 0))

	rootVal, _ := strconv.Atoi(strVal[0])
	root := &TreeNode{
		Val: rootVal,
	}

	Q.Push(root)

	idx := 1

	for idx < len(strVal) && Q.Len() > 0 {

		levelLen := Q.Len()
		for range levelLen {

			node := Q.Pop()

			leftValStr := strVal[idx]
			idx++
			if leftValStr != "null" {
				leftVal, _ := strconv.Atoi(leftValStr)
				leftNode := &TreeNode{
					Val: leftVal,
				}
				node.Left = leftNode
				Q.Push(leftNode)
			}

			rightValStr := strVal[idx]
			idx++
			if rightValStr != "null" {
				rightVal, _ := strconv.Atoi(rightValStr)
				rightNode := &TreeNode{
					Val: rightVal,
				}
				node.Right = rightNode
				Q.Push(rightNode)
			}
		}

	}

	return root

}

// ===================================================

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
