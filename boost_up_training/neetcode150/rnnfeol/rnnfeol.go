package rnnfeol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	// main process
	pointerA, pointerB := dummy
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
