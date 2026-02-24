# 143. Reorder List

<br>

---

<br>

link: https://leetcode.com/problems/reorder-list

<br>

You are given the head of a singly linked-list. The list can be represented as:

```
L0 → L1 → … → Ln - 1 → Ln
```

Reorder the list to be on the following form:

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
```

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

```
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

<br>
<br>

## Coding-1

```go
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		// we need at least 3 nodes.
		return
	}

	// 1. find middle node.
	pointerA, pointerB := head, head

	for pointerB.Next != nil {

		pointerB = pointerB.Next
		if pointerB.Next != nil {
			pointerA = pointerA.Next
			// one more step for pointerB
			pointerB = pointerB.Next
		}
	}

	// 2. reverse
	reverseFrom := pointerA.Next
	pointerA.Next = nil

	previousNode := reverseFrom
	currentNode := reverseFrom.Next
	reverseFrom.Next = nil // this will be the new tail, after reverse

	for currentNode != nil {
		tmp := currentNode.Next
		currentNode.Next = previousNode

		previousNode = currentNode
		currentNode = tmp
	}

	pointerA, pointerB = head, previousNode
	// merge 2 linked list (A, B)

	for pointerB != nil {
		tmpA, tmpB := pointerA.Next, pointerB.Next
		pointerA.Next = pointerB
		pointerB.Next = tmpA

		pointerA = tmpA
		pointerB = tmpB
	}
}
```