package reorder_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. using 2 pointer to find the middle.
	slow, fast := head, head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow will be pointing at the middle.
	// cut second list:
	secHead := slow.Next
	slow.Next = nil
	// now we got head, and sec head. ex: [1, 2, 3] and [4, 5]

	// 2. reverse the sec half linked-list.(using secHead)
	prev := secHead
	curr := secHead.Next
	prev.Next = nil // secHead will be the tail of reversed linked-list
	for curr != nil {
		next := curr.Next // save next
		curr.Next = prev  // reverse link
		prev = curr       // move prev forward
		curr = next       // move curr forward
	}
	// prev is the revered-linked-list new head
	revSecHead := prev

	// 3. merge both linked-list and return. (first list's length will equal or 1 bigger than sec list)
	mergeTwoList(head, revSecHead)
}

func mergeTwoList(head1 *ListNode, head2 *ListNode) {
	for head1 != nil && head2 != nil {
		// 1. store current 2 pointer.
		currA, currB := head1, head2
		// 2. move both pointer forward.
		head1 = head1.Next
		head2 = head2.Next
		// 3. rebuild link
		currA.Next = currB
		currB.Next = head1
	}
}

func debug(node *ListNode) {
	// print list
	res := []int{}

	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	fmt.Println("DEBUG: node-list:", res)
}
