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

// [1, 2, 3, nil, nil, 4, 5]
func (this *Codec) serialize(root *TreeNode) string {
	Q := Queue(make([]*TreeNode, 0))
	ret := []string{}

	Q.Push(root)

	for Q.Len() > 0 {

		thisLevelLen := Q.Len()
		for range thisLevelLen {
			node := Q.Pop()
			if node == nil {
				ret = append(ret, "null")
			} else {
				ret = append(ret, strconv.Itoa(node.Val))
				Q.Push(node.Left)
				Q.Push(node.Right)
			}
		}

	}

	res := strings.Join(ret, ",")
	return res
}

func (this *Codec) deserialize(data string) *TreeNode {
	strVals := strings.Split(data, ",")
	idx := 0

	var buildNode func(strVals []string, idx int) *TreeNode
	buildNode = func(strVals []string, idx int) *TreeNode {
		if idx >= len(strVals) {
			return nil
		}

		valstr := strVals[idx]
		if valstr == "null" {
			return nil
		}

		ival, _ := strconv.Atoi(valstr)
		this := &TreeNode{
			Val:   ival,
			Left:  buildNode(strVals, leftIdx(idx)),
			Right: buildNode(strVals, rightIdx(idx)),
		}

		return this
	}

	return buildNode(strVals, idx)
}

// ===================================================
// func left(idx): idx*2+1
func leftIdx(idx int) int {
	return idx*2 + 1
}

// func right(idx): idx*2+2
func rightIdx(idx int) int {
	return idx*2 + 2
}

type Queue []*TreeNode

func (q Queue) Len() int {
	return len(q)
}

func (q *Queue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Pop() *TreeNode {
	ret := (*q)[0]
	(*q)[0] = nil
	*q = (*q)[1:]
	return ret
}
