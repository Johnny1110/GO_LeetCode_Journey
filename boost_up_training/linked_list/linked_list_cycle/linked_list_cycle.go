package linked_list_cycle

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	seenNodes := make(map[*ListNode]interface{})

	for head.Next != nil {
		if _, seen := seenNodes[head]; seen {
			return true
		} else {
			seenNodes[head] = nil
			head = head.Next
		}
	}

	return false
}
