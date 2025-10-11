# 100. Same Tree

<br>

---

<br>

link: https://leetcode.com/problems/same-tree/description/

<br>

## Thinking

ez..

<br>
<br>

## Coding

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

<br>

![1.png](imgs/1.png)