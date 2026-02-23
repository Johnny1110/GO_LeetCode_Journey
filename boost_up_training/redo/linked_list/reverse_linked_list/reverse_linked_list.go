package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	currentNode := head.Next
	previousNode := head
	head.Next = nil // head will be the tail

	for currentNode != nil {
		tmp := currentNode.Next
		currentNode.Next = previousNode

		previousNode = currentNode
		currentNode = tmp
	}

    return previousNode
}