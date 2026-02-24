package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

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
