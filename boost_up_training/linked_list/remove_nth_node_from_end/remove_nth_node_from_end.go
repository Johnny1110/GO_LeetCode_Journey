package remove_nth_node_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		// because: 1 <= n <= sz, so we can just return nil
		return nil
	}

	cusHeader := &ListNode{
		Val:  -1,
		Next: head,
	}

	// make 2 pointer.
	pointerA := cusHeader
	pointerB := cusHeader
	for range n {
		pointerB = pointerB.Next
	}

	// let pointer-B reach the end
	for pointerB.Next != nil {
		pointerA = pointerA.Next
		pointerB = pointerB.Next
	}

	pointerA.Next = pointerA.Next.Next

	return cusHeader.Next
}
