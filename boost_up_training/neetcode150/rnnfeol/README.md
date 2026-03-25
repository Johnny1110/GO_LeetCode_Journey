19. Remove Nth Node From End of List

<br>

---

<br>

## Coding

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	// main process
	pointerA, pointerB := dummy, dummy
	for range n {
		pointerB = pointerB.Next
	}

	for pointerB.Next != nil {
		pointerA = pointerA.Next
		pointerB = pointerB.Next
	}

	deleteNode := pointerA.Next
	pointerA.Next = deleteNode.Next

	return dummy.Next
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = size of node list 

Time: O(n)

Space: O(1)
```