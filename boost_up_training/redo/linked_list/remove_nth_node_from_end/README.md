# 19. Remove Nth Node From End

<br>

---

<br>

link: https://leetcode.com/problems/remove-nth-node-from-end-of-list

Given the head of a linked list, remove the nth node from the end of the list and return its head.

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Input: head = [1,2], n = 1
Output: [1]
```

<br>
<br>

## Coding-1

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	slow, fast := dummy, dummy

	for range n {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	deleteNode := slow.Next
	if deleteNode.Next != nil {
		slow.Next = deleteNode.Next // concat
		deleteNode.Next = nil
	} else {
		slow.Next = nil
	}

	return dummy.Next
}
```