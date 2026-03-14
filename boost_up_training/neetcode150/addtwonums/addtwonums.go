package addtwonums

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy

	num := 0
	for l1 != nil || l2 != nil || num != 0 {

		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		node := &ListNode{
			Val: num % 10,
		}
		tmp.Next = node
		tmp = node

		num = num / 10

	}

	return dummy.Next
}
