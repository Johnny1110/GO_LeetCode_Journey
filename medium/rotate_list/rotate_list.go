package rotate_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newTailNode := head

	length := calculateLenAndClosedLoop(head)
	newTailIndex := locateTailNodeIdx(length, k)

	// loop to new tail and cut the next pointer
	for i := 1; i <= newTailIndex; i++ {
		newTailNode = newTailNode.Next
	}

	newHeadNode := newTailNode.Next
	newTailNode.Next = nil

	return newHeadNode
}

func calculateLenAndClosedLoop(head *ListNode) int {
	if head == nil {
		return 0
	}

	tmp := head
	count := 1
	for tmp.Next != nil {
		count++
		tmp = tmp.Next
	}

	// connect tail to head
	tmp.Next = head

	return count
}

func locateTailNodeIdx(nodeLen, k int) int {
	return (nodeLen - 1) - k%nodeLen
}
