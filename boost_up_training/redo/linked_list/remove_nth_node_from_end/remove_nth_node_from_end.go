package remove_nth_node_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

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
