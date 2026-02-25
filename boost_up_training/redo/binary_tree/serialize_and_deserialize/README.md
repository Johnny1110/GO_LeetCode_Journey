# 297. Serialize and Deserialize

<br>

---

<br>

link: https://leetcode.com/problems/serialize-and-deserialize-binary-tree

<br>

## Coding-1 (BFS)

```go
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
	if root == nil {
		return ""
	}

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

	if data == "" || len(strVals) == 0 {
		return nil
	}
	
	Q := Queue(make([]*TreeNode, 0))
	val, _ := strconv.Atoi(strVals[0])
	head := &TreeNode{
		Val: val, 
	}

	Q.Push(head)

	idx := 1

	// BFS ================================================
	for Q.Len() != 0 {

		thisLevelLen := Q.Len()
		for range thisLevelLen {
			node := Q.Pop()

			if idx < len(strVals) && strVals[idx] != "null"  {
				leftVal, _ := strconv.Atoi(strVals[idx])
				node.Left = &TreeNode {
					Val: leftVal,
				}
				Q.Push(node.Left)
			} 
			idx++

			if  idx < len(strVals) && strVals[idx] != "null"  {
				rightVal, _ := strconv.Atoi(strVals[idx])
				node.Right = &TreeNode {
					Val: rightVal,
				}
				Q.Push(node.Right)
			}
			idx++
		}
	}
	// BFS ================================================

	return head

}

// ===================================================

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
```

<br>
<br>

## Coding-2 (DFS)

```go
// TODO
```