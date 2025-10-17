package serialize_and_deserialize

import (
	"fmt"
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

	if root == nil {
		return "nil"
	}

	val := root.Val

	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	return fmt.Sprintf("%d,%s,%s", val, left, right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// split strVals from data
	strVals := strings.Split(data, ",")
	idx := 0

	// make a build func
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(strVals) || strVals[idx] == "nil" {
			idx++
			return nil
		}

		val, _ := strconv.Atoi(strVals[idx])
		idx++

		left := build()
		right := build()

		return &TreeNode{
			Val:   val,
			Left:  left,
			Right: right,
		}
	}

	return build()
}

// --------------------------------------------------------------------------
