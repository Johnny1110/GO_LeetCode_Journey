## 1

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	pointer := dummy

	var dfs func(l1 *ListNode, l2 *ListNode, num int)
	dfs = func(l1 *ListNode, l2 *ListNode, num int) {
		if l1 == nil && l2 == nil && num == 0 {
			return
		}

		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + num
		val := sum % 10
		next := sum / 10

		node := &ListNode{
			Val: val,
		}

		pointer.Next = node
		pointer = node

		dfs(l1, l2, next)
	}

	dfs(l1, l2, 0)

	return dummy.Next
}
```