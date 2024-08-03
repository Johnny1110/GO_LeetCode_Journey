package remove_nth_node_from_end_of_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	// def dummyNode : dummy's next one is head.
	dummyNode := &ListNode{Next: head}
	// init 2 pointer, point @ dummyNode.
	slowPointer, fastPointer := dummyNode, dummyNode

	// move fastPointer go n ahead before slowPointer.
	for i := 0; i <= n; i++ {
		fastPointer = fastPointer.Next
	}

	// move both pointer to the right until fastPointer reaching to the end.
	for fastPointer != nil {
		fastPointer = fastPointer.Next
		slowPointer = slowPointer.Next
	}
	// after the loop from above, fastPointer already reached to the end.
	// and the slowPointer will point @ 1 step before the delete node.
	// link the node:
	slowPointer.Next = slowPointer.Next.Next

	return dummyNode.Next
}

// -------------------------------------------------------------------

func Go() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 2
	newHead := removeNthFromEnd(head, n)
	printList(newHead) // Output: 1 2 3 5

	head = &ListNode{1, nil}
	n = 1
	newHead = removeNthFromEnd(head, n)
	printList(newHead) // Output: []

	head = &ListNode{1, &ListNode{2, nil}}
	n = 2
	newHead = removeNthFromEnd(head, n)
	printList(newHead) // Output: []
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
