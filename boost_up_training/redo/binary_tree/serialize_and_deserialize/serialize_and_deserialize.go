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

	strVals := []string{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			strVals = append(strVals, "null")
			return
		}

		sv := strconv.Itoa(node.Val)
		strVals = append(strVals, sv)

		// left & right
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return strings.Join(strVals, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	strVals := strings.Split(data, ",")

	idx := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {

		if idx >= len(strVals) {
			return nil
		}

		strV := strVals[idx]

		idx++

		if strV == "null" {
			return nil
		} else {
			val, _ := strconv.Atoi(strV)
			node := &TreeNode{
				Val:   val,
				Left:  dfs(),
				Right: dfs(),
			}
			return node
		}

	}

	return dfs()
}

// ===================================================
